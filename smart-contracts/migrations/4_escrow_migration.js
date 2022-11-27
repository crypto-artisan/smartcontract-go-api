const EscrowContract = artifacts.require("Escrow");

module.exports = async function (deployer, network, [owner]) {
  deployer.deploy(EscrowContract, owner);
};
