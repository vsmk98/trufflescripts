pragma solidity ^0.4.18;

contract TestContract {

    uint counter = 0;

    event TestEvent(uint count);

    function emitEvent() public {
        emit TestEvent(counter++);
    }

}
