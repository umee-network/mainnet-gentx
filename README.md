# Umee Mainnet Gentxs

This repository contains validated gentx submissions from Umee mainnet genesis
validators.

To create a gentx:
```shell
# Install Umee
$ cd $HOME
$ git clone https://github.com/umee-network/umee.git
$ cd umee
$ git pull
$ git checkout tags/v1.0.1
$ make build
$ sudo cp $HOME/umee/build/umeed /usr/local/bin
$ umeed version
# Should be v1.0.1
```
```shell
# Create gentx
$ umeed init moniker --chain-id umee-1
$ umeed add-genesis-account wallet 1000000uumee
$ umeed gentx-gravity [key_name] [amount] [eth-address] [orchestrator-address]
```

- `key_name` and `amount` are the same as in the default gentx command.
- `eth-address` is the Ethereum address that is going to be used to sign batches
and validator set updates going to Ethereum. It is recommended that this account
is only used for this purpose and nothing else.
- `orchestrator-address`is the Umee account that is going to be used to sign
claims going from Ethereum to Umee. It is recommended that this account
is only used for this purpose and nothing else.

> Please note that the command is `umeed gentx-gravity` and not `umeed gentx`.

Submit your gentx to this repostiory via a Pull Request with the gentx file named
after the validator's moniker, e.g. `myval.json`.

## FAQ

### What is the chain-ID?

`umee-1`

### What version of `umeed` do I use to generate a gentx?

You can use the latest release, v0.7.x at the time of this writing. We will be
publishing an official v1.0.0 release that will be used for mainnet shortly. You
can use that version as well.

### What are the validator commission rules?

There is no minimum or maximum validator commission enforcement by the Umee
protocol. However, we will not accept gentxs with a commission rate of less than
2%. In addition, the Foundation will not delegate to validators whose commission
rate exceeds 10%.

### What do I specify for `amount`, i.e. how much do I self-delegate?

If you have tokens at genesis, you are free to self-delegate how ever many tokens
you wish. The team will make delegation decisions based on previous testnet
performance and strategic considerations. If you do not have any tokens at
genesis but you are highly involved in the Cosmos or Umee ecosystem, you are
suggested to reach out to the team to be considered as a strategic validator and
get a delegation; the team will provide you 1 umee token to bootstrap your validator
and further delegations after TGE.

### Can I use a different account for the orchestrator and the validator?

Yes, you can, and that's the recommended way to do it.

### Can I change the orchestrator address or the orchestrator eth address?

No, once they are set they can't be changed. They are linked to your validator,
so in case you need to change any of these you'll have to start a new validator.

### Can I use these accounts in another system/bot/computer?

It's not recommended, as Peggo will be permanently sending transactions using
these addresses.

If a transaction is made with any of these accounts (orchestrator address or
Ethereum address), it will result in a "nonce unsync" and Peggo will start
throwing some errors. Peggo will re-sync and re-try, but it's better to avoid
that.
