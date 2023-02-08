KEY="alice"
CHAINID="spidex_9222-1"
MONIKER="local"
KEYRING="test"
LOGLEVEL="info"
HOMEDIR="$HOME/.spidex"
# to trace evm
TRACE="--trace"

CODE="608060405234801561001057600080fd5b50600436106100575760003560e01c806314735e2c1461005c5780635a27240314610078578063c4d66de814610094578063e0be4aa3146100b0578063ff1b636d146100e0575b600080fd5b6100766004803603810190610071919061051b565b6100fe565b005b610092600480360381019061008d91906105d5565b6101b5565b005b6100ae60048036038101906100a991906105d5565b610289565b005b6100ca60048036038101906100c59190610602565b610337565b6040516100d7919061065a565b60405180910390f35b6100e8610365565b6040516100f59190610684565b60405180910390f35b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461018e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610185906106fc565b60405180910390fd5b8060008360405161019f9190610796565b9081526020016040518091039020819055505050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610245576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161023c906106fc565b60405180910390fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600160149054906101000a900460ff16156102d9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102d09061081f565b60405180910390fd5b60018060146101000a81548160ff02191690831515021790555080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000818051602081018201805184825260208301602085012081835280955050505050506000915090505481565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6103f2826103a9565b810181811067ffffffffffffffff82111715610411576104106103ba565b5b80604052505050565b600061042461038b565b905061043082826103e9565b919050565b600067ffffffffffffffff8211156104505761044f6103ba565b5b610459826103a9565b9050602081019050919050565b82818337600083830152505050565b600061048861048384610435565b61041a565b9050828152602081018484840111156104a4576104a36103a4565b5b6104af848285610466565b509392505050565b600082601f8301126104cc576104cb61039f565b5b81356104dc848260208601610475565b91505092915050565b6000819050919050565b6104f8816104e5565b811461050357600080fd5b50565b600081359050610515816104ef565b92915050565b6000806040838503121561053257610531610395565b5b600083013567ffffffffffffffff8111156105505761054f61039a565b5b61055c858286016104b7565b925050602061056d85828601610506565b9150509250929050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006105a282610577565b9050919050565b6105b281610597565b81146105bd57600080fd5b50565b6000813590506105cf816105a9565b92915050565b6000602082840312156105eb576105ea610395565b5b60006105f9848285016105c0565b91505092915050565b60006020828403121561061857610617610395565b5b600082013567ffffffffffffffff8111156106365761063561039a565b5b610642848285016104b7565b91505092915050565b610654816104e5565b82525050565b600060208201905061066f600083018461064b565b92915050565b61067e81610597565b82525050565b60006020820190506106996000830184610675565b92915050565b600082825260208201905092915050565b7f4d73672073656e646572206e6f742061646d696e000000000000000000000000600082015250565b60006106e660148361069f565b91506106f1826106b0565b602082019050919050565b60006020820190508181036000830152610715816106d9565b9050919050565b600081519050919050565b600081905092915050565b60005b83811015610750578082015181840152602081019050610735565b8381111561075f576000848401525b50505050565b60006107708261071c565b61077a8185610727565b935061078a818560208601610732565b80840191505092915050565b60006107a28284610765565b915081905092915050565b7f436f6e747261637420696e7374616e63652068617320616c726561647920626560008201527f656e20696e697469616c697a6564000000000000000000000000000000000000602082015250565b6000610809602e8361069f565b9150610814826107ad565b604082019050919050565b60006020820190508181036000830152610838816107fc565b905091905056fea26469706673582212200f493bafc50b32721d3e3bc80275421aa63281ee92bf6e115d1d81df140a0d1664736f6c634300080f0033"
HASH="0x86be6f2ee43edcb01b0882faad8a04cbe2da1a74382e2c8df5cd633eaa199fcd"

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
    sed -i '' 's/create_empty_blocks = true/create_empty_blocks = false/g' $HOMEDIR/config/config.toml
    sed -i '' '$H;x;1,/enable = false/s/enable = false/enable = true/;1d' $HOMEDIR/config/app.toml
    sed -i '' ' s/swagger = false/swagger = true/g' $HOMEDIR/config/app.toml
    sed -i '' 's/enabled-unsafe-cors = false/enabled-unsafe-cors = true/g' $HOMEDIR/config/app.toml
    sed -i '' 's/0.0.0.0:9090/0.0.0.0:9092/g' $HOMEDIR/config/app.toml
  else
    sed -i 's/create_empty_blocks = true/create_empty_blocks = false/g' $HOMEDIR/config/config.toml
    sed -i '$H;x;1,/enable = false/s/enable = false/enable = true/;1d' $HOMEDIR/config/app.toml
    sed -i ' s/swagger = false/swagger = true/g' $HOMEDIR/config/app.toml
    sed -i 's/enabled-unsafe-cors = false/enabled-unsafe-cors = true/g' $HOMEDIR/config/app.toml
    sed -i 's/0.0.0.0:9090/0.0.0.0:9092/g' $HOMEDIR/config/app.toml
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
