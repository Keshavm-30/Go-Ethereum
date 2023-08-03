// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract basic{
    uint public num;

    function inc() public returns(uint)
    {
       num = num+1;
       return num;
    }

    function dec() public returns(uint)
    {
        num = num-1;
        return num;
    }
}