// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

/**
 * Our own implementation of USDC token for private testnet use.
 * On mainnet we use the official USDC token address.
 */
contract USDC is ERC20 {
  constructor() ERC20("USDC Token", "USDC") {
    /// @notice mint 1,000,000 tokens to the owner
    _mint(msg.sender, 1000000e18);
  }
}
