package cli

import (
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/schrsi/d-email/x/email/types"
)

func CmdCreateEmail() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-email [from] [to] [senderSignature] [senderAddressVersion] [subject] [body] [replyTo] [trackIds] [sendedAt] [decryptionKeys] [previousDecryptionKey] [id]",
		Short: "Creates a new email",
		Args:  cobra.ExactArgs(12),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsFrom := args[0]
			argsTo := args[1]
			argsSenderSignature := args[2]

			argsSenderAddressVersion, err := strconv.ParseUint(args[3], 10, 64)
			if err != nil {
				return err
			}

			argsSubject := args[4]
			argsBody := args[5]
			argsReplyTo := args[6]
			argsTrackIds := strings.Split(args[7], ";")
			argsSendedAt := args[8]
			argsDecryptionKeys := strings.Split(args[9], ";")
			argsPreviousDecryptionKey := args[10]
			argsID := args[11]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateEmail(clientCtx.GetFromAddress().String(), argsFrom, argsTo, argsSenderSignature, argsSenderAddressVersion, argsSubject, argsBody, argsReplyTo, argsTrackIds, argsSendedAt, argsDecryptionKeys, argsPreviousDecryptionKey, argsID)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
