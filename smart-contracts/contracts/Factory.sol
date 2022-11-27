// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./CustodianWalletLogic.sol";
import "./CustodianWalletProxy.sol";
import "./Escrow.sol";

/**
 * @dev Factotory responsible for deploying Escrow, Custodian Wallet Logic and
 * managed CustodianWalletProxy for both Vendor and Customer.
 *
 * Note Controlled by deployer account.
 */
contract Factory {
  ////////////////////////////////////////
  //                                    //
  //         STATE VARIABLES            //
  //                                    //
  ////////////////////////////////////////

  /// @notice EOA of deployer wallet
  address public ochestrator;

  /// @notice the address of the Custodian Wallet Logic
  address public custodianWalletLogic;

  address public escrowContractAddress;

  /// @notice mapping of account unique id to custodian wallet
  /// Note avoid passing predictable number such as incremental number. Use UUID string instead
  mapping(string => address) public accounts;

  event NewCustodian(string uniqueId, address indexed account);

  ////////////////////////////////////////
  //                                    //
  //              FUNCTIONS             //
  //                                    //
  ////////////////////////////////////////
  modifier onlyOchesrator() {
    require(msg.sender == ochestrator, "F: only ochestrator");
    _;
  }

  constructor() {
    ochestrator = msg.sender;
    escrowContractAddress = address(new Escrow(msg.sender));
    custodianWalletLogic = address(new CustodianWalletLogic());
  }

  /**
   * @dev create a new custodian wallet
   * @param uuid is the unique id of the custodian wallet
   */
  function newCustodian(string memory uuid)
    public
    onlyOchesrator
    returns (string memory, address)
  {
    require(accounts[uuid] == address(0x0), "F: account exist");

    address wallet = address(
      new CustodianWalletProxy(custodianWalletLogic, ochestrator, address(this))
    );
    accounts[uuid] = wallet;

    emit NewCustodian(uuid, wallet);

    return (uuid, wallet);
  }
}
