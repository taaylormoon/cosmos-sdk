<!--
order: 3
-->

# End Block

At each `EndBlock`, the fees received are transferred to the distribution `ModuleAccount`, as it's the account the one who keeps track of the flow of coins in (as in this case) and out the module. The fees are also allocated to the proposer, community fund and global pool. When the validator is the proposer of the round, that validator (and their delegators) receives between 1% and 5% of fee rewards, the reserve community tax is then charged, then the remainder is distributed proportionally by voting power to all bonded validators independent of whether they voted (social distribution). Note the social distribution is applied to proposer validator in addition to the proposer reward.

The amount of proposer reward is calculated from pre-commits Tendermint messages in order to incentivize validators to wait and include additional pre-commits in the block. All provision rewards are added to a provision reward pool which validator holds individually (`ValidatorDistribution.ProvisionsRewardPool`).

```go
func AllocateTokens(feesCollected sdk.Coins, feePool FeePool, proposer ValidatorDistribution, 
              sumPowerPrecommitValidators, totalBondedTokens, communityTax, 
              proposerCommissionRate sdk.Dec)

     SendCoins(FeeCollectorAddr, DistributionModuleAccAddr, feesCollected)
     feesCollectedDec = MakeDecCoins(feesCollected)
     proposerReward = feesCollectedDec * (0.01 + 0.04 
                       * sumPowerPrecommitValidators / totalBondedTokens)

     commission = proposerReward * proposerCommissionRate
     proposer.PoolCommission += commission
     proposer.Pool += proposerReward - commission

     communityFunding = feesCollectedDec * communityTax
     feePool.CommunityFund += communityFunding

     poolReceived = feesCollectedDec - proposerReward - communityFunding
     feePool.Pool += poolReceived

     SetValidatorDistribution(proposer)
     SetFeePool(feePool)
```
