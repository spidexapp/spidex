check_results=`hermes keys add --chain spidex_9111-1 --mnemonic-file .mem`
chainaAddress=$(echo $check_results | cut -d '(' -f2|cut -d ')' -f1)

check_results=`hermes keys add --chain band-laozi-testnet5 --mnemonic-file .mem`
chainbAddress=$(echo $check_results | cut -d '(' -f2|cut -d ')' -f1)

echo $chainaAddress
echo $chainbAddress

#get some token
curl --location --request POST 'https://laozi-testnet5.bandchain.org/faucet' \
--header 'Content-Type: application/json' \
--data-raw '{"address":"'$chainbAddress'"}'

spidexd tx bank send alice $chainaAddress 100000000000000000aspx --from alice -y

#conn_results=`hermes create connection spidex_9111-1 band-laozi-testnet5`
#
#connection=$(echo $conn_results | tail -n 22 | head -n 1 |awk -F \' '{print $2}')

channel_results=`hermes create channel --a-chain spidex_9111-1 --b-chain band-laozi-testnet5 --a-port oracle --b-port oracle --new-client-connection --channel-version bandchain-1 --yes`

$channel_results > channelout.txt
echo $(tail -n 11 $channel_results | head -n 1 |awk -F \" '{print $2}')
echo $(tail -n 39 $channel_results | head -n 1 |awk -F \" '{print $2}')



#echo tail -n 11 channelout.txt | head -n 1 |awk -F \' '{print $2}'
#echo channelout.txt | tail -n 22 | head -n 1 |awk -F \' '{print $2}'
