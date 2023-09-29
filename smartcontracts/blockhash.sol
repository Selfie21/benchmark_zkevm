pragma solidity >=0.7.0 <0.9.0;

contract Analysis {
    
    function Benchmark(uint limit) public {
        for(uint i = 0; i < limit; i++){
            blockhash(0);
        }
    }

}
