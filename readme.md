# Taroly

The Multi Ethereum Fork Simulation Manager

## Usage

You need to set the following environment variables:
```
TAROLY_ROOT_USER Admin telegram user id
TAROLY_RPC_URL Mapping of url and chain names. 
TAROLY_TELEGRAM_TOKEN Telegram bot token
TAROLY_WEB_URL Web url for trace command
```

TAROLY_RPC_URL Example:
```json
{
  "eth": "https://rpc.ankr.com/eth",
  "bsc": "https://bsc-dataseed.binance.org/"
}
```

Optional environment variables:
```
FOUNDRY_PATH_OVERRIDE Path to the foundry binary
FORK_PORT_START Port to start the forked chains. default 11400
FORK_PORT_END Port to end the forked chains. default FORK_PORT_START + 100
```

### Deployment

#### Docker

Use `nekogawa/taroly:latest` image.

#### Docker Compose

See `docker-compose.yml` for example.  
Please notice that docker-compose.override.yml is for running with traefik.

## Commands

### /auth

Authenticates the user to use the bot.

### /onlinemode [true|false]

Sets the online mode. If online mode is true, the bot will use the real blockchain. If online mode is false, the bot will use the forked blockchain.
Notice: transact will not work in online mode.

### /getchain 

See the chain valid to use for online mode.

### /setchain [chain name]

Sets the chain to use for online mode.

### /newfork chain_name [fork_block_number]

Creates a new fork for the given chain name. If fork_block_number is not provided, the latest block will be used.

### /getrpc 

Get endpoint for your rpc

### /stopfork

Stops the current fork.

### /setaccount address

Sets the account to use for the current fork.

### /getaccount

Gets the account to use for the current fork.

### /call contract_address [function_name] [arguments]

Calls the given function of the given contract with the given arguments. if no function name is provided, the fallback function will be called.

### /calldata contract_address [calldata]

Calls the given contract with the given calldata.

### /getbalance [address]

Gets the balance of the given address. If no address is provided, the balance of the current account will be returned.

### /setbalance amount

Sets the balance of the current account.

### /transact contract_address [function_name] [arguments]

Transacts the given function of the given contract with the given arguments. if no function name is provided, the fallback function will be called.

### /transactdata contract_address [calldata]

todo

Transacts the given contract with the given calldata.

### /transactwithvalue contract_address amount [function_name] [arguments]

todo

Same with transact, but with the given amount.

### /transactdatawithvalue contract_address amount [calldata]

todo

Same with transactwithvalue, but with the given amount.

### /trace [tx_hash]

Get transaction trace result.

Traces the given transaction.

### /4byte function_signature

Gets the function signature of the given function signature.

### /4decode call_data

Decodes the given call data.

### /4encode function_signature [arguments]

Encodes the given function signature with the given arguments.