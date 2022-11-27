// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Types {
  /// @notice structure of an order
  struct Order {
    address seller;
    address buyer;
    address receiver;
    uint256 amount;
    uint256 rate;
    uint256 fee;
    uint8 orderType;
    uint8 orderStatus;
    uint256 startTime;
    uint256 fulfiledTime;
  }

  Order[] public orders;

  mapping(address => uint256[]) public openOrders;

  /// @notice 0 = BUY, 1 = SELL
  enum OrderType {
    BUY,
    SELL
  }

  enum OrderStatus {
    OPEN, // order is open
    FULFILLED, // order is fulfilled
    CANCELLED, // order is cancelled
    EXPIRED, // order is expired (not fulfilled)
    REJECTED, // order is rejected by seller
    REFUNDED // order is refunded to the seller
  }

  ////////////////////////////////////////
  //                                    //
  //              EVENTS                //
  //                                    //
  ////////////////////////////////////////

  event OpenOrder(
    uint256 orderId,
    address indexed seller,
    address indexed buyer,
    address indexed receiver,
    uint256 amount,
    uint256 rate,
    uint256 fee,
    uint8 orderType,
    uint8 orderStatus
  );

  event OrderFulfilled(uint256 orderId);

  event ClosedOrder(
    uint256 orderId,
    address indexed seller,
    address indexed buyer,
    address indexed receiver,
    uint256 amount,
    uint256 rate,
    uint256 fee,
    uint8 orderType,
    uint256 fulfiledTime,
    uint8 orderStatus
  );

  event RejectedOrder(uint256 orderId);

  event ApproveRejectedOrder(uint256 orderId);
}
