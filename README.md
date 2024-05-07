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