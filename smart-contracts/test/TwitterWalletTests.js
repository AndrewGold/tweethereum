const BN = web3.utils.BN
const MockERC20 = artifacts.require("MockERC20")
const TwitterWallet = artifacts.require("TwitterWallet")
const mockHandle = "mockHandle"
const mockTokenAddress = "0xe41d2489571d322189246dafa5ebde1f4699f498"

contract("TwitterWallet", async accounts => {
    it("wallet ether balance should be empty", async () => {
        const wallet = await TwitterWallet.new()
        const balance = await wallet.etherBalance(mockHandle)

        assert.equal(balance, 0)
    })

    it("wallet token balance should be empty", async () => {
        const wallet = await TwitterWallet.new()
        const balance = await wallet.tokenBalance(mockHandle, mockTokenAddress)

        assert.equal(balance, 0)
    })

    it("shoud fund twitter handle with ether", async () => {
        const newBalance = 10
        const wallet = await TwitterWallet.new()

        await wallet.fundHandleEther(mockHandle, {from: accounts[0], value: newBalance})

        const balance = await wallet.etherBalance(mockHandle)
        assert.equal(balance, newBalance)
    })

    it("should fund twitter handle with token", async () => {
        const tokenBalance = 7
        const walletBalance = 5
        const wallet = await TwitterWallet.new()
        const token = await MockERC20.new()

        await token.mint(accounts[0], tokenBalance)
        await token.approve(wallet.address, 1000)

        let initialTokenBalance = await token.balanceOf(accounts[0])
        assert.equal(initialTokenBalance, tokenBalance)

        await wallet.fundHandleToken(mockHandle, token.address, walletBalance)

        let newWalletBalance = await wallet.tokenBalance(mockHandle, token.address)
        assert.equal(newWalletBalance, walletBalance)

        let remainingTokenBalance = await token.balanceOf(accounts[0])
        assert.equal(remainingTokenBalance, tokenBalance - walletBalance)
    })

    it("should transfer ether balance to handle", async () => {
        const mockDestinationHandle = "destinationHandle"
        const walletBalance = 10
        const transferAmount = 4
        const wallet = await TwitterWallet.new()

        await wallet.fundHandleEther(mockHandle, {from: accounts[0], value: walletBalance})
        await wallet.transferEtherToHandle(mockHandle, mockDestinationHandle, transferAmount)

        const newSenderBalance = await wallet.etherBalance(mockHandle)
        assert.equal(newSenderBalance, walletBalance - transferAmount)

        const destinationBalance = await wallet.etherBalance(mockDestinationHandle)
        assert.equal(destinationBalance, transferAmount)
    })

    it("should transfer ether balance to address", async () => {
        const walletBalance = 10
        const transferAmount = 4
        const wallet = await TwitterWallet.new()
        const originalDestinationBalance = new BN(await web3.eth.getBalance(accounts[1]))

        await wallet.fundHandleEther(mockHandle, {from: accounts[0], value: walletBalance})        
        await wallet.transferEtherToAddress(mockHandle, accounts[1], transferAmount)

        const newSenderBalance = await wallet.etherBalance(mockHandle)
        assert.equal(newSenderBalance, walletBalance - transferAmount)

        const newDestinationBalance = new BN(await web3.eth.getBalance(accounts[1]))
        assert.equal(newDestinationBalance.sub(originalDestinationBalance), transferAmount)
    })

    it("should transfer token balance to handle", async () => {
        const mockDestinationHandle = "destinationHandle"
        const walletBalance = 10
        const transferAmount = 4
        const wallet = await TwitterWallet.new()
        const token = await MockERC20.new()

        await token.mint(accounts[0], 1000)
        await token.approve(wallet.address, 1000)

        await wallet.fundHandleToken(mockHandle, token.address, walletBalance)
        await wallet.transferTokenToHandle(mockHandle, mockDestinationHandle, token.address, transferAmount)

        const newSenderBalance = await wallet.tokenBalance(mockHandle, token.address)
        assert.equal(newSenderBalance, walletBalance - transferAmount)

        const newDestinationBalance = await wallet.tokenBalance(mockDestinationHandle, token.address)
        assert.equal(newDestinationBalance, transferAmount)
    })

    it("should transfer token balance to address", async () => {
        const walletBalance = 10
        const transferAmount = 4
        const wallet = await TwitterWallet.new()
        const token = await MockERC20.new()

        await token.mint(accounts[0], 1000)
        await token.approve(wallet.address, 1000)

        await wallet.fundHandleToken(mockHandle, token.address, walletBalance)
        await wallet.transferTokenToAddress(mockHandle, accounts[1], token.address, transferAmount)

        const newSenderBalance = await wallet.tokenBalance(mockHandle, token.address)
        assert.equal(newSenderBalance, walletBalance - transferAmount)

        const newDestinationBalance = await token.balanceOf(accounts[1])
        assert.equal(newDestinationBalance, transferAmount)
    })
})

