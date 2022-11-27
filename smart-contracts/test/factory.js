const Factory = artifacts.require("Factory");
const Escrow = artifacts.require("Escrow");
const USDC = artifacts.require("USDC");
const { expectRevert } = require("@openzeppelin/test-helpers");
const { utils } = require("ethers");

contract("Factory", function ([deployer, account2]) {
  before(async function () {
    this.factory = await Factory.deployed();
    this.usdc = await USDC.deployed();
    this.escrow = await Escrow.deployed();
  });

  it("should set deployer as ochestrator address", async function () {
    const ochestrator = await this.factory.ochestrator();

    return assert.equal(deployer, ochestrator);
  });

  it("should be able to update USDC token address", async function () {
    const newTokenAddress = await this.usdc.address;

    await this.escrow.setUsdcTokenAddress(newTokenAddress);

    const usdcToken = await this.escrow.usdcToken();

    return assert.equal(usdcToken, newTokenAddress);
  });

  it("should revert if called by another account rather than deployer", async function () {
    const newTokenAddress = await this.usdc.address;

    await expectRevert(
      this.escrow.setUsdcTokenAddress(newTokenAddress, { from: account2 }),
      "F: only ochestrator"
    );
  });

  it("should revert if called with an invalid address", async function () {
    const invalidAddress = "0x0000000000000000000000000000000000000000";

    await expectRevert(
      this.escrow.setUsdcTokenAddress(invalidAddress),
      "F: invalid address"
    );
  });

  it("should create custodian wallet", async function () {
    // Pass a UUID to idenify a user account onchain
    // Ideal UUID must be a string and aboid non incremental numbers
    // Example package [https://www.npmjs.com/package/uuid]
    const response = await this.factory.newCustodian(
      "1b9d6bcd-bbfd-4b2d-9b5d-ab8dfbbd4bed"
    );

    const { 0: uuid, 1: custodianWallet } = response.logs[0].args;

    assert.equal(uuid, "1b9d6bcd-bbfd-4b2d-9b5d-ab8dfbbd4bed");
    assert.isTrue(utils.isAddress(custodianWallet));
  });

  it("should revert when uuid with account exist", async function () {
    await expectRevert(
      this.factory.newCustodian("1b9d6bcd-bbfd-4b2d-9b5d-ab8dfbbd4bed"),
      "F: account exist"
    );
  });
});
