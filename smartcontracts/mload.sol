pragma solidity >=0.7.0 <0.9.0;

contract Analysis {
    
    function Benchmark(uint limit) public {
        uint256 sumOfArray;
        uint256[] memory _array = new uint256[](1);
        _array[0] = 13984;
        for(uint i = 0; i < limit; i++){
            sumOfArray += _array[0];
        }
    } 

}
