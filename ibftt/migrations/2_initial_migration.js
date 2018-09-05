var Product  = artifacts.require("./Product.sol");

module.exports = function(deployer) {
  deployer.deploy(Product, "My-Prod1", "ABC1");
};
