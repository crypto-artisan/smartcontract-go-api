const FactoryContract = artifacts.require("Factory");

module.exports = async function (deployer) {
  deployer.deploy(FactoryContract);
};
