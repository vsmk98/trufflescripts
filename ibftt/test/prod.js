var ProdCatalog = artifacts.require("./ProductCatalog.sol")


var perm;
var foundEvent = "NO";
var eventName = "";
var nodeId = 0;

var prodCatalog = ProdCatalog.at("0x9d13c6d3afe1721beef56b55d303b09e021e27ab");

function printEvent (result){
	for (var i = 0; i < result.logs.length; i++) {
    var log = result.logs[i];
    console.log(log);
    eventName = log.event;
    foundEvent = "YES";
    product = log.args.productAddress;
    productName = log.args.name;
    productId = log.args.productId;
  }
  if (foundEvent == "YES"){
    console.log("Yes. detected the event!!!!!" + eventName + "---" + product +"/"+productName + "/" + productId);
  } else {
    console.log("No.Did not find the event !!!!!!");

  }
}
prodCatalog.createProduct("MyProd2", "ABC2") .then (function(result) {
	printEvent(result);
});
//	perm.getNodeStatus("bnhb1096ca56b9f6d004b779ae3728bf83f8e22453404cc3cef16a3d9b96608bc67c4b30db88e0a5a6c6390213f7acbe1153ff6d23ce57380104288ae19373ef") .then (function(result) {
//	printEvent(result);
//	console.log(result);
//});
