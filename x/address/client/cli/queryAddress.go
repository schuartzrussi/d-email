package cli

import (
	"context"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/schrsi/d-email/x/address/types"
	"github.com/spf13/cobra"
)

func CmdListAddress() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-address",
		Short: "list all address",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllAddressRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.AddressAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowAddressByName() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-address [name]",
		Short: "shows the latest version of an adress",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			name := args[0]

			params := &types.QueryGetAddressByNameRequest{
				Name: name,
			}

			res, err := queryClient.AddressByName(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowAddressByVersion() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-address-by-version [name] [version]",
		Short: "shows a specific version of an address",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			name := args[0]
			version, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryGetAddressByVersionRequest{
				Name:    name,
				Version: version,
			}

			res, err := queryClient.AddressByVersion(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
