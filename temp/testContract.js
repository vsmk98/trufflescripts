const Web3 = require('web3');

const TestContract = artifacts.require('TestContract');

// List of nodes we're checking for transactions.
const nodeNames = [
  'master-quorum',
  'partner-a-quorum',
  'partner-b-quorum',
];

// Checks nodes for transactions.
async function checkNode(nodeName, transactionHashes) {
  const providerURL = "http://127.0.0.1";
  const providerPort1 = "22000";
  const providerPort = "22001";
  const providerPort = "22002";
  if (nodeName == 'master-quorum'){
    const web3 = new Web3(new Web3.providers.HttpProvider(providerURL + ":" + providerPort1));
  } else if (nodeName == ''){
    const web3 = new Web3(new Web3.providers.HttpProvider(providerURL + ":" + providerPort2));
  } else {
    const web3 = new Web3(new Web3.providers.HttpProvider(providerURL + ":" + providerPort3));
  }
  const outputObject = {
    found: 0,
    missing: 0,
    logs: [],
  };
  const promises = [];
  transactionHashes.forEach((hash) => {
    promises.push(web3.eth.getTransactionReceipt(hash));
  });
  const receipts = await Promise.all(promises);
  receipts.forEach((receipt) => {
    if (Array.isArray(receipt.logs) && receipt.logs.length > 0) {
      outputObject.found += 1;
      let removedZeroes = receipt.logs[0].data.replace(/0x0*/, '');
      if (removedZeroes === '') removedZeroes = 0;
      outputObject.logs.push(removedZeroes);
    } else {
      outputObject.missing += 1;
    }
  });
  return outputObject;
}


module.exports = async (deployer) => {
  console.log('Creating a contract private for partner A');
  await deployer.deploy(TestContract, { privateFor: ['QfeDAys9MPDs2XHExtc84jKGHxZg/aj52DTh0vtA3Xc='] });
  const instance = await TestContract.deployed();
  let counter = 0;
  const transactionHashes = [];
  console.log('Creating 10 functions private for partner A');
  const interval = setInterval(async () => {
    const functionReturn = await instance.emitEvent({ privateFor: ['QfeDAys9MPDs2XHExtc84jKGHxZg/aj52DTh0vtA3Xc='] });
    transactionHashes.push(functionReturn.tx);
    counter += 1;
    if (counter === 10) {
      clearInterval(interval);
      nodeNames.forEach(async (node) => {
        const results = await checkNode(node, transactionHashes);
        // console.log(results);
        console.log(`In node ${node}, ${results.found} transactions were found and ${results.missing} were not`);
        console.log('This node can see the following logs:', results.logs);
      });
    }
  }, 1000);
};
