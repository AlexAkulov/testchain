# testchain
**testchain** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

# Usage CLI

## set-state

```
testchaind tx cache set-state 1 one --from alice
testchaind tx cache set-state 2 two --from alice
```

## get-state
```
testchaind q cache get-state 1
value: one
testchaind q cache get-state 2
value: two
```

# Usage API

## get-state

```
curl "http://localhost:1317/testchain/cache/get_state/1"
{"value":"one"}%
curl "http://localhost:1317/testchain/cache/get_state/1"
{"value":"two"}%
```
