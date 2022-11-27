// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./Types.sol";

contract CustodianWalletProxy is Types {
    /// @notice address of factory
    address public factory;

    /// @notice address of wallet logic to
    /// copy code from and call using delegatecall
    address public immutable logic;

    /// @notice address of escrow contract
    address public immutable ochestrator;

    /**
     * @param _logic address of already deployed Custodian Wallet that can receive upgrade
   * @param _ochestrator address of Escrow that can has sole control over all custodian wallets
   */
    constructor(
        address _logic,
        address _ochestrator,
        address _factory
    ) {
        ochestrator = _ochestrator;
        logic = _logic;
        factory = _factory;
    }

    // prettier-ignore
    /**
     * @dev Forward any call to any function with any set of parameters to the logic contract
   * without it needing to know anything in particular of the logic contract’s interface.
   *
   * Note By using this proxy approach we are able to upgrade the logic contract at any time and allow
   * already deployed proxy wallets benefit from new updates.
   *
   * Since msg.sender when calling Wallet Logic chnages to address(this) due to usage of delegate call
   * we are unable to determinable set a modifier on functions in Wallet Logic to limit calls to Escrow.
   * So we require that caller to all wallet proxies to call our wallet logic can only be escrow.
   *
   *
   * Credit to https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/proxy/Proxy.sol
   */
    fallback() external payable {
        // solhint-disable-line no-complex-fallback

        require(msg.sender == ochestrator, "WP: deployer only");

        address _impl = logic;

        // This is never suppose to happen by any chance
        if (_impl == address(0)) {
            revert("Logic contract not set");
        }

        assembly {
        // solhint-disable-line no-inline-assembly
        // Copy msg.data. We take full control of memory in this inline assembly
        // block because it will not return to Solidity code. We overwrite the
        // Solidity scratch pad at memory position 0.
            calldatacopy(0, 0, calldatasize())

        // Call the implementation.
        // out and outsize are 0 because we don't know the size yet.
            let result := delegatecall(gas(), _impl, 0, calldatasize(), 0, 0)

        // Copy the returned data.
            returndatacopy(0, 0, returndatasize())

            switch result
            // delegatecall returns 0 on error.
            case 0 {
                revert(0, returndatasize())
            }
            default {
                return (0, returndatasize())
            }
        }
    }

    // ======== Receive =========

    receive() external payable {} // solhint-disable-line no-empty-blocks
}
