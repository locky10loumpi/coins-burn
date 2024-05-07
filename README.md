#Coins-Burn

Based on [an initial idea](https://github.com/atmoner/ignite-burn), it proposes an implementation with ignite v28.3.

## Get started

```bash
cd protocol
ignite c serve
```

## Build process

```bash
ignite s chain mychain --minimal --no-module --skip-git --path ./protocol
```

In the `./prorocol` directory.

construct the coin-actions module
```bash
ignite s module coin-actions --dep bank
```

create the burn-action message
```bash
ignite s message BurnCoinsAction coins:coins --module coin-actions
```

edit: `x/coinactions/types/expected_keepers.go`, add in `BankKeeper`
```go
// BankKeeper defines the expected interface for the Bank module.
type BankKeeper interface {
	SpendableCoins(context.Context, sdk.AccAddress) sdk.Coins
	// Methods imported from bank should be defined here
	BurnCoins(ctx context.Context, burn string, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx context.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
}
```

then edit: `x/coinactions/keeper/msg_server_burn_coins_action.go` with the burn logic

## local tests

In the `./prorocol` directory.
```bash
ignite c serve
```

then in an other terminal. Run the helper
```bash
export alice=`mychaind keys show -a alice`
```
it defines some env data from ignite keys.


Then show the alice balances
```bash
mychaind q bank balances $alice
```

Execute a burn
```bash
mychaind tx coinactions burn-coins-action 2000stake --from $alice --chain-id="mychain"
```

Check the updates
```bash
mychaind q bank balances $alice
```

Wouhou !!!