var Simple  = artifacts.require("./simplestorage.sol");

module.exports = function(deployer) {
  deployer.deploy(Simple, 42);
};
