# trufflescripts
Smart contract based permissions model for managing account and node permissions, transaction manager key management. 
Account permissions
Account can have any one of the following 4 permission types – “Read Only”, “Transact”, “Contract Deploy” and “Full access”
“Read Only” account can read the chain data.
An account with “Transact” permission can perform value transfer or contract level transactions
An account with “Contract Deploy” permissions can perform value transfer, contract level transactions and deploy new contracts 
“Full Access” accounts can perform network level activities and all the above
Node Permissions
Node permissions feature will allow the existing nodes to propose a new node to join the network and once majority voting is done, the new node can join the network
At network level voters need to be added for voting on any node permission activities. These accounts should have “Full Access”
Transaction manager key management
This feature allows multiple transaction manager keys to be grouped under a single organization name. The organization name can then be used in “privateFor” attribute instead of individual transaction manager keys.
It allows to define a hierarchy of master organization and multiple sub organizations under the master org
