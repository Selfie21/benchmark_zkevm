pragma solidity >=0.7.0 <0.9.0;

contract Analysis {
    
    function Benchmark(uint limit) public{
        bytes32 messageHash;
        for(uint i = 0; i < limit; i++){
            messageHash = keccak256(bytes("Hello!"));
        }
    }
}
