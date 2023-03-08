KEY="alice"
CHAINID="spidex_9222-1"
MONIKER="local"
KEYRING="os"
LOGLEVEL="info"
HOMEDIR="$HOME/.spidex"
# to trace evm
TRACE="--trace"

CODE="608060405234801561001057600080fd5b506004361061007d5760003560e01c80635a2724031161005b5780635a272403146100bd5780637936f8b7146100d0578063e0be4aa314610100578063ff1b636d1461013957600080fd5b806314735e2c1461008257806323c2a31e14610097578063485cc955146100aa575b600080fd5b6100956100903660046103fe565b61014c565b005b6100956100a536600461045f565b6101eb565b6100956100b8366004610481565b61024c565b6100956100cb36600461045f565b6102fa565b6002546100e3906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b61012b61010e3660046104b4565b805160208183018101805160008252928201919093012091525481565b6040519081526020016100f7565b6001546100e3906001600160a01b031681565b6001546001600160a01b031633148061016f57506002546001600160a01b031633145b6101945760405162461bcd60e51b815260040161018b906104f1565b60405180910390fd5b806000836040516101a5919061051f565b90815260405160209181900382018120929092558282527fcbfdd485c676f9bb5c081fdf614d5a06394642373d34d5098c8ef7120c7ba4e5910160405180910390a15050565b6001546001600160a01b031633148061020e57506002546001600160a01b031633145b61022a5760405162461bcd60e51b815260040161018b906104f1565b600280546001600160a01b0319166001600160a01b0392909216919091179055565b600254600160a01b900460ff16156102bd5760405162461bcd60e51b815260206004820152602e60248201527f436f6e747261637420696e7374616e63652068617320616c726561647920626560448201526d195b881a5b9a5d1a585b1a5e995960921b606482015260840161018b565b60028054600180546001600160a01b039586166001600160a01b0319909116179055919092166001600160a81b031990911617600160a01b179055565b6001546001600160a01b031633148061031d57506002546001600160a01b031633145b6103395760405162461bcd60e51b815260040161018b906104f1565b600180546001600160a01b0319166001600160a01b0392909216919091179055565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261038257600080fd5b813567ffffffffffffffff8082111561039d5761039d61035b565b604051601f8301601f19908116603f011681019082821181831017156103c5576103c561035b565b816040528381528660208588010111156103de57600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000806040838503121561041157600080fd5b823567ffffffffffffffff81111561042857600080fd5b61043485828601610371565b95602094909401359450505050565b80356001600160a01b038116811461045a57600080fd5b919050565b60006020828403121561047157600080fd5b61047a82610443565b9392505050565b6000806040838503121561049457600080fd5b61049d83610443565b91506104ab60208401610443565b90509250929050565b6000602082840312156104c657600080fd5b813567ffffffffffffffff8111156104dd57600080fd5b6104e984828501610371565b949350505050565b60208082526014908201527326b9b39039b2b73232b9103737ba1030b236b4b760611b604082015260600190565b6000825160005b818110156105405760208186018101518583015201610526565b8181111561054f576000828501525b50919091019291505056fea264697066735822122027581984ba657ba06c8368d5e9e443a5a33d8fe99d3f4d03a0a492938e9d39c764736f6c634300080d0033"
HASH="0xd61ca71d2847f20070a9660868be1f5c51f9b9d96e45930b44bee1fa57ab41a9"

# validate dependencies are installed
#command -v jq > /dev/null 2>&1 || { echo >&2 "jq not installed. More info: https://stedolan.github.io/jq/download/"; exit 1; }

# Reinstall daemon
rm -rf $HOMEDIR/data/
rm -rf $HOMEDIR/config/
#rm -rf $HOMEDIR/
rm -rf ~/.ignite/relayer

# Set client config
spidexd config keyring-backend $KEYRING --home $HOMEDIR
spidexd config chain-id $CHAINID --home $HOMEDIR

# if $KEY exists it should be deleted
spidexd keys add $KEY --keyring-backend $KEYRING --home $HOMEDIR

# Set moniker and chain-id for cosmos (Moniker can be anything, chain-id must be an integer)
spidexd init $MONIKER --chain-id $CHAINID --home $HOMEDIR

cp $HOMEDIR/config/genesis.json $HOMEDIR/config/tmp_genesis.json

# Change parameter aspx denominations to aspx
cat $HOMEDIR/config/genesis.json | jq '.app_state["staking"]["params"]["bond_denom"]="aspx"' > $HOMEDIR/config/tmp_genesis.json && mv $HOMEDIR/config/tmp_genesis.json $HOMEDIR/config/genesis.json
cat $HOMEDIR/config/genesis.json | jq '.app_state["slashing"]["params"]["signed_blocks_window"]="5000"' > $HOMEDIR/config/tmp_genesis.json && mv $HOMEDIR/config/tmp_genesis.json $HOMEDIR/config/genesis.json
cat $HOMEDIR/config/genesis.json | jq '.app_state["crisis"]["constant_fee"]["denom"]="aspx"' > $HOMEDIR/config/tmp_genesis.json && mv $HOMEDIR/config/tmp_genesis.json $HOMEDIR/config/genesis.json
cat $HOMEDIR/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="aspx"' > $HOMEDIR/config/tmp_genesis.json && mv $HOMEDIR/config/tmp_genesis.json $HOMEDIR/config/genesis.json
cat $HOMEDIR/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["amount"]="64"' > $HOMEDIR/config/tmp_genesis.json && mv $HOMEDIR/config/tmp_genesis.json $HOMEDIR/config/genesis.json
cat $HOMEDIR/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["max_deposit_period"]="180s"' > $HOMEDIR/config/tmp_genesis.json && mv $HOMEDIR/config/tmp_genesis.json $HOMEDIR/config/genesis.json
cat $HOMEDIR/config/genesis.json | jq '.app_state["gov"]["voting_params"]["voting_period"]="180s"' > $HOMEDIR/config/tmp_genesis.json && mv $HOMEDIR/config/tmp_genesis.json $HOMEDIR/config/genesis.json

cat $HOMEDIR/config/genesis.json | jq '.app_state["evm"]["params"]["evm_denom"]="aspx"' > $HOMEDIR/config/tmp_genesis.json && mv $HOMEDIR/config/tmp_genesis.json $HOMEDIR/config/genesis.json
cat $HOMEDIR/config/genesis.json | jq '.app_state["inflation"]["params"]["mint_denom"]="aspx"' > $HOMEDIR/config/tmp_genesis.json && mv $HOMEDIR/config/tmp_genesis.json $HOMEDIR/config/genesis.json
cat $HOMEDIR/config/genesis.json | jq '.app_state["mint"]["params"]["mint_denom"]="aspx"' > $HOMEDIR/config/tmp_genesis.json && mv $HOMEDIR/config/tmp_genesis.json $HOMEDIR/config/genesis.json

cat $HOMEDIR/config/genesis.json | jq '.app_state["evm"]["accounts"][0]["address"]="0x4200000000000000000000000000000000000042"' > $HOMEDIR/config/tmp_genesis.json && mv $HOMEDIR/config/tmp_genesis.json $HOMEDIR/config/genesis.json
cat $HOMEDIR/config/genesis.json | jq -r --arg CODE "$CODE" '.app_state["evm"]["accounts"][0]["code"]=$CODE' > $HOMEDIR/config/tmp_genesis.json && mv $HOMEDIR/config/tmp_genesis.json $HOMEDIR/config/genesis.json


# Set gas limit in genesis
cat $HOMEDIR/config/genesis.json | jq '.consensus_params["block"]["max_gas"]="40000000"' > $HOMEDIR/config/tmp_genesis.json && mv $HOMEDIR/config/tmp_genesis.json $HOMEDIR/config/genesis.json

# Set claims start time
node_address=$(spidexd keys list --home $HOMEDIR | grep  "address: " | cut -c12-)
current_date=$(date -u +"%Y-%m-%dT%TZ")
cat $HOMEDIR/config/genesis.json | jq -r --arg current_date "$current_date" '.app_state["claims"]["params"]["airdrop_start_time"]=$current_date' > $HOMEDIR/config/tmp_genesis.json && mv $HOMEDIR/config/tmp_genesis.json $HOMEDIR/config/genesis.json

# Set claims records for validator account
amount_to_claim=10000
cat $HOMEDIR/config/genesis.json | jq -r --arg node_address "$node_address" --arg amount_to_claim "$amount_to_claim" '.app_state["claims"]["claims_records"]=[{"initial_claimable_amount":$amount_to_claim, "actions_completed":[false, false, false, false],"address":$node_address}]' > $HOMEDIR/config/tmp_genesis.json && mv $HOMEDIR/config/tmp_genesis.json $HOMEDIR/config/genesis.json

# Set claims decay
cat $HOMEDIR/config/genesis.json | jq -r --arg current_date "$current_date" '.app_state["claims"]["params"]["duration_of_decay"]="1000000s"' > $HOMEDIR/config/tmp_genesis.json && mv $HOMEDIR/config/tmp_genesis.json $HOMEDIR/config/genesis.json
cat $HOMEDIR/config/genesis.json | jq -r --arg current_date "$current_date" '.app_state["claims"]["params"]["duration_until_decay"]="100000s"' > $HOMEDIR/config/tmp_genesis.json && mv $HOMEDIR/config/tmp_genesis.json $HOMEDIR/config/genesis.json

# Set oracle params
cat $HOMEDIR/config/genesis.json | jq -r --arg node_address "$node_address" '.app_state["oracle"]["params"]["admin"]=$node_address' > $HOMEDIR/config/tmp_genesis.json && mv $HOMEDIR/config/tmp_genesis.json $HOMEDIR/config/genesis.json
cat $HOMEDIR/config/genesis.json | jq  '.app_state["oracle"]["params"]["scriptId"]=111' > $HOMEDIR/config/tmp_genesis.json && mv $HOMEDIR/config/tmp_genesis.json $HOMEDIR/config/genesis.json
cat $HOMEDIR/config/genesis.json | jq  '.app_state["oracle"]["params"]["multiplier"]=1000000' > $HOMEDIR/config/tmp_genesis.json && mv $HOMEDIR/config/tmp_genesis.json $HOMEDIR/config/genesis.json
cat $HOMEDIR/config/genesis.json | jq  '.app_state["oracle"]["coinSymbolList"][0]["id"]=0' > $HOMEDIR/config/tmp_genesis.json && mv $HOMEDIR/config/tmp_genesis.json $HOMEDIR/config/genesis.json
cat $HOMEDIR/config/genesis.json | jq  '.app_state["oracle"]["coinSymbolList"][0]["symbol"]="BTC"' > $HOMEDIR/config/tmp_genesis.json && mv $HOMEDIR/config/tmp_genesis.json $HOMEDIR/config/genesis.json
cat $HOMEDIR/config/genesis.json | jq  '.app_state["oracle"]["coinSymbolList"][0]["isDeleted"]=false' > $HOMEDIR/config/tmp_genesis.json && mv $HOMEDIR/config/tmp_genesis.json $HOMEDIR/config/genesis.json

# disable produce empty block
if [[ "$OSTYPE" == "darwin"* ]]; then
    sed -i '' 's/127.0.0.1:26657/0.0.0.0:26657/g' $HOMEDIR/config/config.toml
    sed -i '' 's/create_empty_blocks = true/create_empty_blocks = false/g' $HOMEDIR/config/config.toml
    sed -i '' '$H;x;1,/enable = false/s/enable = false/enable = true/;1d' $HOMEDIR/config/app.toml
    sed -i '' ' s/swagger = false/swagger = true/g' $HOMEDIR/config/app.toml
    sed -i '' 's/enabled-unsafe-cors = false/enabled-unsafe-cors = true/g' $HOMEDIR/config/app.toml
  else
    sed -i 's/127.0.0.1:26657/0.0.0.0:26657/g' $HOMEDIR/config/config.toml
    sed -i 's/create_empty_blocks = true/create_empty_blocks = false/g' $HOMEDIR/config/config.toml
    sed -i '$H;x;1,/enable = false/s/enable = false/enable = true/;1d' $HOMEDIR/config/app.toml
    sed -i ' s/swagger = false/swagger = true/g' $HOMEDIR/config/app.toml
    sed -i 's/enabled-unsafe-cors = false/enabled-unsafe-cors = true/g' $HOMEDIR/config/app.toml
fi


if [[ "$OSTYPE" == "darwin"* ]]; then
      sed -i '' 's/create_empty_blocks_interval = "0s"/create_empty_blocks_interval = "30s"/g' $HOMEDIR/config/config.toml
      sed -i '' 's/timeout_propose = "3s"/timeout_propose = "2s"/g' $HOMEDIR/config/config.toml
      sed -i '' 's/timeout_propose_delta = "500ms"/timeout_propose_delta = "500ms"/g' $HOMEDIR/config/config.toml
      sed -i '' 's/timeout_prevote = "1s"/timeout_prevote = "1s"/g' $HOMEDIR/config/config.toml
      sed -i '' 's/timeout_prevote_delta = "500ms"/timeout_prevote_delta = "500ms"/g' $HOMEDIR/config/config.toml
      sed -i '' 's/timeout_precommit = "1s"/timeout_precommit = "1s"/g' $HOMEDIR/config/config.toml
      sed -i '' 's/timeout_precommit_delta = "500ms"/timeout_precommit_delta = "500ms"/g' $HOMEDIR/config/config.toml
      sed -i '' 's/timeout_commit = "5s"/timeout_commit = "2s"/g' $HOMEDIR/config/config.toml
      sed -i '' 's/timeout_broadcast_tx_commit = "2m30s"/timeout_broadcast_tx_commit = "10s"/g' $HOMEDIR/config/config.toml
  else
      sed -i 's/create_empty_blocks_interval = "0s"/create_empty_blocks_interval = "30s"/g' $HOMEDIR/config/config.toml
      sed -i 's/timeout_propose = "3s"/timeout_propose = "2s"/g' $HOMEDIR/config/config.toml
      sed -i 's/timeout_propose_delta = "500ms"/timeout_propose_delta = "500ms"/g' $HOMEDIR/config/config.toml
      sed -i 's/timeout_prevote = "1s"/timeout_prevote = "1s"/g' $HOMEDIR/config/config.toml
      sed -i 's/timeout_prevote_delta = "500ms"/timeout_prevote_delta = "500ms"/g' $HOMEDIR/config/config.toml
      sed -i 's/timeout_precommit = "1s"/timeout_precommit = "1s"/g' $HOMEDIR/config/config.toml
      sed -i 's/timeout_precommit_delta = "500ms"/timeout_precommit_delta = "500ms"/g' $HOMEDIR/config/config.toml
      sed -i 's/timeout_commit = "5s"/timeout_commit = "2s"/g' $HOMEDIR/config/config.toml
      sed -i 's/timeout_broadcast_tx_commit = "2m30s"/timeout_broadcast_tx_commit = "10s"/g' $HOMEDIR/config/config.toml
fi

# Allocate genesis accounts (cosmos formatted addresses)
spidexd add-genesis-account $KEY 1000000000000000000000000000aspx --keyring-backend $KEYRING --home $HOMEDIR

cat $HOMEDIR/config/genesis.json | jq '.app_state["auth"]["accounts"][1]["@type"]="/ethermint.types.v1.EthAccount"' > $HOMEDIR/config/tmp_genesis.json && mv $HOMEDIR/config/tmp_genesis.json $HOMEDIR/config/genesis.json
cat $HOMEDIR/config/genesis.json | jq '.app_state["auth"]["accounts"][1]["base_account"]["address"]="spx1ggqqqqqqqqqqqqqqqqqqqqqqqqqqqqzzssvl7k"' > $HOMEDIR/config/tmp_genesis.json && mv $HOMEDIR/config/tmp_genesis.json $HOMEDIR/config/genesis.json
cat $HOMEDIR/config/genesis.json | jq '.app_state["auth"]["accounts"][1]["base_account"]["pub_key"]=null' > $HOMEDIR/config/tmp_genesis.json && mv $HOMEDIR/config/tmp_genesis.json $HOMEDIR/config/genesis.json
cat $HOMEDIR/config/genesis.json | jq '.app_state["auth"]["accounts"][1]["base_account"]["account_number"]="1"' > $HOMEDIR/config/tmp_genesis.json && mv $HOMEDIR/config/tmp_genesis.json $HOMEDIR/config/genesis.json
cat $HOMEDIR/config/genesis.json | jq '.app_state["auth"]["accounts"][1]["base_account"]["sequence"]="0"' > $HOMEDIR/config/tmp_genesis.json && mv $HOMEDIR/config/tmp_genesis.json $HOMEDIR/config/genesis.json
cat $HOMEDIR/config/genesis.json | jq  -r --arg HASH "$HASH" '.app_state["auth"]["accounts"][1]["code_hash"]=$HASH' > $HOMEDIR/config/tmp_genesis.json && mv $HOMEDIR/config/tmp_genesis.json $HOMEDIR/config/genesis.json

# Update total supply with claim values
validators_supply=$(cat $HOMEDIR/config/genesis.json | jq -r '.app_state["bank"]["supply"][0]["amount"]')
## Bc is required to add this big numbers
## total_supply=$(bc <<< "$amount_to_claim+$validators_supply")
total_supply=1000000000000000000000000000
cat $HOMEDIR/config/genesis.json | jq -r '.app_state["bank"]["supply"][0]["denom"]="aspx"' > $HOMEDIR/config/tmp_genesis.json && mv $HOMEDIR/config/tmp_genesis.json $HOMEDIR/config/genesis.json
cat $HOMEDIR/config/genesis.json | jq -r --arg total_supply "$total_supply" '.app_state["bank"]["supply"][0]["amount"]=$total_supply' > $HOMEDIR/config/tmp_genesis.json && mv $HOMEDIR/config/tmp_genesis.json $HOMEDIR/config/genesis.json

# Sign genesis transaction
spidexd gentx $KEY 1000spx --keyring-backend $KEYRING --chain-id $CHAINID --home $HOMEDIR

# Collect genesis tx
spidexd collect-gentxs --home $HOMEDIR

# Run this to ensure everything worked and that the genesis file is setup correctly
spidexd validate-genesis --home $HOMEDIR

#Start the node
echo spidexd start --pruning=nothing $TRACE --log_level $LOGLEVEL --json-rpc.api eth,txpool,net,web3 --home $HOMEDIR
