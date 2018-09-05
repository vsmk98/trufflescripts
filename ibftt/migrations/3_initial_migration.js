var ProductCatalog  = artifacts.require("./ProductCatalog.sol");

module.exports = function(deployer) {
  deployer.deploy(ProductCatalog, 42);
};
