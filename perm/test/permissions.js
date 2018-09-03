var Permissions = artifacts.require("./Permissions.sol")


var perm;
var foundEvent = "NO";
var eventName = "";
var nodeId = 0;

var perm = Permissions.at("0x0000000000000000000000000000000000000020");
// var perm = Permissions.at("0x10ae69385c79ef3eb815ac008a7013d6878f1d38");

function printEvent (result){
	for (var i = 0; i < result.logs.length; i++) {
      var log = result.logs[i];
      console.log(log);
      eventName = log.event;
      foundEvent = "YES";
      nodeId = log.args._nodeId;
    }
    if (foundEvent == "YES"){
            console.log("Yes. detected the event!!!!!" + eventName + "---" + nodeId);
    } else {
            console.log("No.Did not find the event !!!!!!");
    }
}

 // perm.updateAcctAccess("0xed9d02e382b34818e88b88a309c7fe71e65f419d", 1) .then (function(result) {
 // printEvent(result);
 // });
// perm.ProposeNode("3701f007bfa4cb26512d7df18e6bbd202e8484a6e11d387af6e482b525fa25542d46ff9c99db87bd419b980c24a086117a397f6d8f88e74351b41693880ea0cb","true", "true", "127.0.0.1:21004", "0", "50445").then (function(result) {
// 	printEvent(result);
//   });
// perm.ApproveNode("3701f007bfa4cb26512d7df18e6bbd202e8484a6e11d387af6e482b525fa25542d46ff9c99db87bd419b980c24a086117a397f6d8f88e74351b41693880ea0cb").then (function(result) {
//   printEvent(result);
// });
// perm.getNodeStatus("xxxxx74c4b0e7a9e12d2fe5fee6595eda841d6d992c35dbbcc50fcee4aa86dfbbdeff7dc7e72c2305d5a62257f82737a8cffc80474c15c611c037f52db1a3a7b") .then (function(result) {
//   console.log(result);
// });

// perm.ProposeDeactivation("axxxx74c4b0e7a9e12d2fe5fee6595eda841d6d992c35dbbcc50fcee4aa86dfbbdeff7dc7e72c2305d5a62257f82737a8cffc80474c15c611c037f52db1a3a7b") .then (function(result) {
//   printEvent(result);
// });
// perm.DeactivateNode("xxxxx74c4b0e7a9e12d2fe5fee6595eda841d6d992c35dbbcc50fcee4aa86dfbbdeff7dc7e72c2305d5a62257f82737a8cffc80474c15c611c037f52db1a3a7b") .then (function(result) {
// 	printEvent(result);
// });
// perm.ProposeNodeBlacklisting("eacaa74c4b0e7a9e12d2fe5fee6595eda841d6d992c35dbbcc50fcee4aa86dfbbdeff7dc7e72c2305d5a62257f82737a8cffc80474c15c611c037f52db1a3a7b", "127.0.0.1:21005", "0", "50446").then (function(result) {
// 	printEvent(result);
// });
perm.BlacklistNode("eacaa74c4b0e7a9e12d2fe5fee6595eda841d6d992c35dbbcc50fcee4aa86dfbbdeff7dc7e72c2305d5a62257f82737a8cffc80474c15c611c037f52db1a3a7b").then (function(result) {
	printEvent(result);
});
