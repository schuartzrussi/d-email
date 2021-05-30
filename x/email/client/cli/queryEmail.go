package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/schrsi/d-email/x/email/types"
	"github.com/spf13/cobra"
)

func CmdListEmail() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-email",
		Short: "list all email",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllEmailRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.EmailAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowEmail() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-email [id]",
		Short: "shows a email",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			id := args[0]

			params := &types.QueryGetEmailByIdRequest{
				Id: id,
			}

			res, err := queryClient.EmailById(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
