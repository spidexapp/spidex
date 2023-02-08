package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	"github.com/spidexapp/spidex/x/oracle/types"
)

var _ = strconv.Itoa(0)

func CmdCreateCoinSymbol() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-coin-symbol [symbol]",
		Short: "Broadcast message create-coin-symbol",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSymbol := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateCoinSymbol(
				clientCtx.GetFromAddress().String(),
				argSymbol,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
