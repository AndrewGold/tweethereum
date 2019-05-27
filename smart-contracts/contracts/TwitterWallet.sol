pragma solidity ^0.5.2;

import '../node_modules/openzeppelin-solidity/contracts/math/SafeMath.sol';
import '../node_modules/openzeppelin-solidity/contracts/ownership/Ownable.sol';
import '../node_modules/openzeppelin-solidity/contracts/token/ERC20/ERC20.sol';
import '../node_modules/openzeppelin-solidity/contracts/token/ERC20/SafeERC20.sol';

/**
  * @title The Tweethereum Twitter Wallet Contract
  * @author Andrew Gold
 */
contract TwitterWallet is Ownable {
    using SafeMath for uint;
    using SafeERC20 for ERC20;

    /**
      * @dev map: Twitter handle -> wei balance
     */
    mapping(string => uint) private _weiBalances;

    /**
      * @dev map: Twitter handle -> contract address -> token balance
     */
    mapping(string => mapping(address => uint)) private _tokenBalances;

    /**
      * @dev fund a Twitter handle with wei in the transaction
      * @param handle the Twitter handle to fund
      * @return true if successful
     */
    function fundHandleEther(string memory handle) public payable returns (bool) {
        _weiBalances[handle] = _weiBalances[handle].add(msg.value);
        return true;
    }

    /**
      * @notice the token must be pre-approved for transfer
      * @dev fund a Twitter handle with tokens
      * @param handle the Twitter handle to fund
      * @param token the address of the token to fund
      * @param value the value of the token to fund
      * @return true if successful
     */
    function fundHandleToken(string memory handle, address token, uint value) public returns (bool) {
        ERC20 tokenContract = ERC20(token);

        tokenContract.transferFrom(msg.sender, address(this), value);
        _tokenBalances[handle][token] = _tokenBalances[handle][token].add(value);
        return true;
    }

    /**
      * @dev get the wei balance of Twitter handle
      * @param handle the Twitter handle
      * @return the wei balance of the handle
     */
    function etherBalance(string memory handle) public view returns (uint) {
        return _weiBalances[handle];
    }

    /**
      * @dev get the token balance of a Twitter handle
      * @param handle the Twitter handle
      * @param token the address of the token
      * @return the token balance of the handle
     */
    function tokenBalance(string memory handle, address token) public view returns (uint) {
        return _tokenBalances[handle][token];
    }

    /**
      * @dev transfer wei to a Twitter handle
      * @param from the Twitter handle to transfer from
      * @param to the Twitter handle to transfer to
      * @param value the value in wei to transfer
      * @return true if successful
     */
    function transferEtherToHandle(string memory from, string memory to, uint value) public onlyOwner() returns (bool) {
        _weiBalances[from] = _weiBalances[from].sub(value);
        _weiBalances[to] = _weiBalances[to].add(value);
        return true;
    }

    /**
      * @dev transfer wei to an address
      * @param from the Twitter handle to transfer from
      * @param to the address to transfer to
      * @param value the value in wei to transfer
      * @return true if successful
     */
    function transferEtherToAddress(string memory from, address payable to, uint value) public onlyOwner() returns (bool) {
        _weiBalances[from] = _weiBalances[from].sub(value);
        to.transfer(value);
        return true;
    }

    /**
      * @dev transfer tokens to a Twitter handle
      * @param from the Twitter handle to transfer from
      * @param to the Twitter handle to transfer to
      * @param token the address of the token
      * @param value the value of tokens to transfer
      * @return true if successful
     */
    function transferTokenToHandle(string memory from, string memory to, address token, uint value) public onlyOwner() returns (bool) {
        _tokenBalances[from][token] = _tokenBalances[from][token].sub(value);
        _tokenBalances[to][token] = _tokenBalances[to][token].add(value);
        return true;
    }

    /**
      * @dev transfer tokens to an address
      * @param from the Twitter handle to transfer from
      * @param to the address to transfer to
      * @param token the address of the token
      * @param value the value of tokens to transfer
      * @return true if successful
     */
    function transferTokenToAddress(string memory from, address to, address token, uint value) public onlyOwner() returns (bool) {
        ERC20 tokenContract = ERC20(token);

        _tokenBalances[from][token] = _tokenBalances[from][token].sub(value);
        return tokenContract.transfer(to, value);
    }
}