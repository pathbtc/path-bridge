pragma solidity 0.6.4;
pragma experimental ABIEncoderV2;

import "./interfaces/IBridgeCounter.sol";

contract BridgeCounter is IBridgeCounter{

    mapping(address => bool) public approvedBridge;

    address public owner;

    // destinationChainID => number of deposits
    mapping(uint8 => uint64) public depositCounts;

    constructor() public {
        owner = msg.sender;
    }


    modifier onlyOwner() {
        require(msg.sender == owner,
            "sender is not  admin");
        _;
    }

    modifier onlyBridge() {
        require(approvedBridge[msg.sender],
            "sender is not bridge");
        _;
    }

    function incr(uint8 destinationChainID) public override onlyBridge returns(uint64) {
        return ++depositCounts[destinationChainID];
    }

    function addBridge(address bridge_) public onlyOwner {
        approvedBridge[bridge_] = true;
    }
    function removeBridge(address bridge_) public onlyOwner{
        delete approvedBridge[bridge_];
    }

    function setDestStartNonce(uint8 destinationChainID,uint64 start) public onlyOwner {
        depositCounts[destinationChainID] = start;
    }
}
