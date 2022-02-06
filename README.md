# Umee Mainnet Gentxs

This repository contains validated gentx submissions from Umee mainnet genesis
validators.

To create a gentx:

```
$ umeed gentx-gravity [key_name] [amount] [eth-address] [orchestrator-address]
```

> Please note that the command is `umeed gentx-gravity` and not `umeed gentx`.

Submit your gentx to this repostiory via a Pull Request with the gentx file named
after the validator's moniker, e.g. `myval.json`.


## FAQ

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
