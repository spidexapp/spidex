package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/spidexapp/spidex/x/oracle/types"
)

// CmdRequestCoinRatesData creates and broadcast a CoinRates request transaction
func CmdRequestCoinRatesData() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "coin-rates-data [oracle-script-id] [requested-validator-count] [sufficient-validator-count]",
		Short: "Make a new CoinRates query request via an existing BandChain oracle script",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {

			// retrieve the requested validator count.
			askCount, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			// retrieve the sufficient(minimum) validator count.
			minCount, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			// retrieve the amount of coins allowed to be paid for oracle request fee from the pool account.
			coinStr, err := cmd.Flags().GetString(flagFeeLimit)
			if err != nil {
				return err
			}
			feeLimit, err := sdk.ParseCoinsNormalized(coinStr)
			if err != nil {
				return err
			}

			// retrieve the amount of gas allowed for the prepare step of the oracle script.
			prepareGas, err := cmd.Flags().GetUint64(flagPrepareGas)
			if err != nil {
				return err
			}

			// retrieve the amount of gas allowed for the execute step of the oracle script.
			executeGas, err := cmd.Flags().GetUint64(flagExecuteGas)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCoinRatesData(
				clientCtx.GetFromAddress().String(),
				askCount,
				minCount,
				feeLimit,
				prepareGas,
				executeGas,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(flagFeeLimit, "", "the maximum tokens that will be paid to all data source providers")
	cmd.Flags().Uint64(flagPrepareGas, 200000, "Prepare gas used in fee counting for prepare request")
	cmd.Flags().Uint64(flagExecuteGas, 200000, "Execute gas used in fee counting for execute request")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
