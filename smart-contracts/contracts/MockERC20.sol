pragma solidity ^0.5.2;

import '../node_modules/openzeppelin-solidity/contracts/token/ERC20/ERC20.sol';

contract MockERC20 is ERC20 {
    function mint(address account, uint256 value) public {
        _mint(account, value);
    }
}