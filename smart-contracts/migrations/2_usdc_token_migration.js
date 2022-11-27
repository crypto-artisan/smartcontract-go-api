const USDCContract = artifacts.require("USDC");

module.exports = async function (deployer) {
  deployer.deploy(USDCContract);
};
