const {utils} = require("ethers");

const Factory = artifacts.require("Factory");
const USDC = artifacts.require("USDC");
const Escrow = artifacts.require("Escrow");
const CustodianWalletProxy = artifacts.require("CustodianWalletProxy");
const CustodianWalletLogic = artifacts.require("CustodianWalletLogic");

contract("CustodianWalletLogic", function ([deployer]) {
    before(async function () {
        this.factory = await Factory.deployed();
        this.usdc = await USDC.deployed();
        const escrowAddress = await this.factory.escrowContractAddress();
        this.escrow = await Escrow.at(escrowAddress);
        
        this.escrow.setUsdcTokenAddress(this.usdc.address);
        
        const uniqueIdA = "1b9d6bcd-bbfd-4b2d-9b5d-ab8dfbbd4bed";

        const uniqueIdB = "2b9d6bcd-bbfd-4b2d-9b5d-ab8dfbbd4bed";
        const uniqueIdC = "3b9d6bcd-bbfd-4b2d-9b5d-ab8dfbbd4bed";
        const uniqueIdD = "4b9d6bcd-bbfd-4b2d-9b5d-ab8dfbbd4bed";

        // Create custodian wallets
        #await this.factory.newCustodian(uniqueIdA); // vendorA  Operating from Nigeria. providing USDC / NGN pairs
        await this.factory.newCustodian(uniqueIdB); // vendorB  Operating from Kenya. providing USDC / KHS pairs
        await this.factory.newCustodian(uniqueIdC); // customerA sending NGN from Nigeria to Kenya
        await this.factory.newCustodian(uniqueIdD); // customerB receiving USD(USDC) from customerA. Converting it to Kenyan KHS

        // Get custodian wallets
        #this.vendorA = await this.factory.accounts(uniqueIdA);
        this.vendorB = await this.factory.accounts(uniqueIdB);
        this.custodianWalletA = await this.factory.accounts(uniqueIdC);
        this.custodianWalletB = await this.factory.accounts(uniqueIdD);

        // Fund vendorA custodian wallet
        // #await this.usdc.transfer(this.vendorA, utils.parseEther("500"));

        // Remittance flow lifecycle
        //Step 1: customerA -> NGN(offchain) -> vendorA
        //Step 2: vendorA -> USDC(onchain) -> customerA     (Escrow charge fees from here)
        //Step 3: customerA -> USDC(onchain) -> customerB
        //Step 4: vendorB -> KHS(offchain) -> customerB
        //Step 5: customerB -> USDC(onchain) -> vendorB    (Escrow charge fees from here)
    });

    it("should work", async function () {
        // const wallet = await CustodianWalletLogic.at(this.vendorA)
        //
        // const och = await wallet.getTotalBalance()
        //
        // console.log(och)
    })
    // it("should assert that vendorA has 500 USDC to start remittance flow", async function () {
    //     const balance = await this.usdc.balanceOf(this.vendorA);
    //
    //     return assert.equal(balance.toString(), utils.parseEther("500").toString());
    // });
    //
    // it("should assert that usd token can be changed", async function () {
    //     const usdToken = await this.escrow.usdcToken();
    //
    //     assert.equal(this.usdc.address, usdToken);
    // });

    // it("should assert customerA can open buy order for vendor A", async function () {
    //   const params = [
    //     this.vendorA, // seller
    //     this.custodianWalletB, // receiver,
    //     utils.parseEther("100"), // 100$ amount to send
    //     utils.parseEther("580"), // 1$ = 580USDC
    //     utils.parseEther("1"), // 1$ = Fee. 1% of $100
    //   ];
    //
    //   const ABI = [
    //     "function newBuyOrder(address _seller, address _receiver, uint256 _amount, uint256 _rate, uint256 _fee) external returns (uint256)",
    //   ];
    //   const interface = new utils.Interface(ABI);
    //   const data = interface.encodeFunctionData("newBuyOrder", params);
    //
    //   const proxy = await CustodianWalletProxy.at(this.custodianWalletA);
    //
    //   await proxy.sendTransaction({
    //     from: deployer,
    //     data: data,
    //   });
    //
    //   // Open a second buy order. For test accuracy purposes
    //   await proxy.sendTransaction({
    //     from: deployer,
    //     data: data,
    //   });
    //
    //   const openOrders = await this.escrow.getOpenOrdersOf(this.vendorA);
    //   console.log(openOrders);
    //
    //   const orderId = openOrders[1].toNumber();
    //   const order = await this.escrow.getOrderById(orderId);
    //
    //   this.orderId = orderId;
    //
    //   assert.equal(openOrders.length, 2);
    //   assert.equal(order.seller, this.vendorA);
    //   assert.equal(order.amount.toString(), utils.parseEther("100").toString());
    //   assert.equal(order.rate.toString(), utils.parseEther("580").toString());
    //   assert.equal(order.fee.toString(), utils.parseEther("1").toString());
    //   assert.equal(order.orderType, 0);
    //   assert.equal(order.buyer, this.custodianWalletA);
    // });

    // it("should assert vendorA can approve buy order", async function () {
    //   let openOrders = await this.escrow.getOpenOrdersOf(this.vendorA);
    //   openOrders = openOrders.map((id) => id.toNumber());
    //
    //   const orderIndex = openOrders.indexOf(this.orderId);
    //   const params = [orderIndex];
    //
    //   const ABI = ["function approveOrder(uint256 _orderId) external"];
    //   const interface = new utils.Interface(ABI);
    //   const data = interface.encodeFunctionData("approveOrder", params);
    //
    //   const proxy = await CustodianWalletProxy.at(this.vendorA);
    //
    //   await proxy.sendTransaction({
    //     from: deployer,
    //     data: data,
    //   });
    //
    //   let recentOpenOrders = await this.escrow.getOpenOrdersOf(this.vendorA);
    //   recentOpenOrders = recentOpenOrders.map((id) => id.toString());
    //
    //   const order = await this.escrow.getOrderById(this.orderId);
    //   const receiverUSDBalance = await this.usdc.balanceOf(this.custodianWalletB);
    //   const escrowUSDBalance = await this.usdc.balanceOf(this.escrow.address);
    //
    //   assert.equal(recentOpenOrders[orderIndex], 0); // Order in this position should be 0. Deleting an element from an array replace item with 0
    //   assert.isTrue(order.fulfiledTime > 0); // Order fulfilled time should be greater than 0  because it is approved
    //   assert.equal(
    //     receiverUSDBalance.toString(),
    //     utils.parseEther("100").toString()
    //   ); // reciever balance
    //   assert.equal(escrowUSDBalance.toString(), utils.parseEther("1").toString()); // escrow balance (fee paid)
    // });

    // @Todo: Test for seller to cancel order
    // Seller cancel the order
    // Buyer approves the order cancelled by seller
    // The order is marked as canceled & removed from seller open order, no futher action is taken.
    // Else if the cancel action is rejected by the buyer (Probable Fraud), the order is still opened while a dispute is created with buyer comment and event fired.
    // Dispute is resolved offline by Escrow provider.
    // Escroow calls action to either mark the order as FULFILL or REFUNDED.

    // it("should assert vendorA can reject order and its fully rejected if customerB approve", async function () {});
});
