pragma solidity ^0.4.0;

contract Permissions {
    struct nodeDetails {
        string enodeId;
        bool canWrite;
        bool canLead;
        bool active;
    }
    uint nodeNum;
    mapping (uint => nodeDetails) nodelist;
    
    event NewNodeAdded (uint _nodeId, string _enodeId, bool _canWrite, bool _canLead);
    
    function Permissions(){
        
    }
    
    function AddNode(string _enodeId, bool _canWrite, bool _canLead)  returns (uint _nodeId){
        nodeNum ++;
        _nodeId = nodeNum;
        nodelist[_nodeId] = nodeDetails(_enodeId, _canWrite, _canLead, true);
        
        NewNodeAdded (_nodeId, _enodeId, _canWrite, _canLead);
        return _nodeId;
    }
    
}
