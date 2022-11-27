// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./Types.sol";
import "./Factory.sol";

contract CustodianWalletLogic is Types {
  /// @notice address of factory
  address public factory;

  function getTotalBalance() external view returns (uint256) {
    return _getUsdcBalance();
  }

  function getOpenOrders() external view returns (uint256[] memory) {
    return _getOpenOrders();
  }

  function _getUsdcBalance() internal view returns (uint256) {
    return IERC20(_getEscrow().usdcToken()).balanceOf(msg.sender);
  }

  function _getEscrow() internal view returns (Escrow) {
    return Escrow(Factory(factory).escrowContractAddress());
  }

  /// @notice when a customer buy USD with local fiat
  function newBuyOrder(
    address _seller,
    address _receiver,
    uint256 _amount,
    uint256 _rate,
    uint256 _fee
  ) external returns (uint256) {
    address usdcAddress = _getEscrow().usdcToken();

    require(_seller != address(0x0), "CWL: seller not set");
    require(usdcAddress != address(0x0), "CWL: usdc token not set");

    require(
      IERC20(usdcAddress).balanceOf(_seller) >= _amount + _fee,
      "C: not enough USD"
    );

    return
      _getEscrow().newOrder(
        _seller, // vendor
        address(this),
        _receiver,
        _amount,
        _rate,
        _fee,
        0 // buy
      );
  }

  /// @notice when a customer sell USD for local fiat to vendor
  function newSellOrder(
    address _buyer,
    address _receiver,
    uint256 _amount,
    uint256 _rate,
    uint256 _fee
  ) external returns (uint256) {
    return
      _getEscrow().newOrder(
        address(this),
        _buyer, //vendor
        _receiver,
        _amount,
        _rate,
        _fee,
        1 // sell
      );
  }

  /// @notice returns operating balance of the seller custodian wallet (total USD balance - open orders against wallet)
  function availBalance() external view returns (uint256) {
    return _availBalance();
  }

  function _availBalance() internal view returns (uint256) {
    uint256[] memory openOrders = _getOpenOrders();
    uint256 balance = _getUsdcBalance();

    for (uint256 queue = 0; queue < (openOrders.length - 1); queue++) {
      Order memory order = _getEscrow().getOrderById(openOrders[queue]);
      balance -= order.amount; // subtract amount of open order
      balance -= order.fee; // subtract fee of open order
    }

    return balance;
  }

  function _getOpenOrders() internal view returns (uint256[] memory) {
    return _getEscrow().getOpenOrdersOf(address(this));
  }

  function approveOrder(uint256 _openOrderIndex) external {
    uint256[] memory openOrders = _getOpenOrders();

    require(openOrders.length > 0, "CWL: no open orders");
    uint256 _orderId = openOrders[_openOrderIndex];

    Order memory order = _getEscrow().getOrderById(_orderId);

    require(order.seller == address(this), "CWL: invalid order");

    _getEscrow().closeOpenOrder(address(this), _openOrderIndex);

    _sendFunds(order.receiver, order.amount, order.fee);

    emit OrderFulfilled(_orderId);
  }

  function rejectOrder(uint256 _orderId) external {
    Order memory order = _getEscrow().getOrderById(_orderId);

    require(order.seller == address(this), "CWL: only seller");

    _getEscrow().rejectOrder(address(this), _orderId);
  }

  function consentOrderRejected(uint256 _openOrderIndex) external {
    uint256[] memory openOrders = _getOpenOrders();

    uint256 _orderId = openOrders[_openOrderIndex];

    Order memory order = _getEscrow().getOrderById(_orderId);

    require(order.buyer == address(this), "CWL: only buyer");

    _getEscrow().consentOrderRejected(address(this), _openOrderIndex);
  }

  function _sendFunds(
    address _to,
    uint256 _amount,
    uint256 _fee
  ) internal {
    require(_to != address(this), "CWL: self forbidden");
    require(_to != address(0x0), "CWL: invalid to address");
    require(_amount > 0, "CWL: amount cannot equal 0");
    require(_availBalance() >= _amount, "CWL: insufficient funds");

    IERC20(_getEscrow().usdcToken()).transfer(_to, _amount);
    IERC20(_getEscrow().usdcToken()).transfer(address(_getEscrow()), _fee);
  }
}
