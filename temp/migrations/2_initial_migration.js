const TestContract = artifacts.require('./TestContract.sol');

module.exports = (deployer) => {
  deployer.deploy(TestContract);
};
