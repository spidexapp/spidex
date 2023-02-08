//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.0;

contract Oracle {
    mapping(string => uint256) public Price;
    address public Admin;
    bool private initialized;

    function initialize(address newAdmin) public {
        require(!initialized, "Contract instance has already been initialized");
        initialized = true;
        Admin = newAdmin;
    }

    modifier onlyAdmin() {
        require(msg.sender == Admin, "Msg sender not admin");
        _;
    }

    function SetPrice(string memory symbol, uint256 price) public onlyAdmin {
        Price[symbol] = price;
    }

    function SetAdmin(address newAdmin) public onlyAdmin {
        Admin = newAdmin;
    }
}
