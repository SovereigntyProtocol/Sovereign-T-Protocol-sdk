package tokenfactory

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/cosmos/cosmos-sdk/tokenfactory"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "DenomAll",
					Use:       "list-denom",
					Short:     "List all Denom",
				},
				{
					RpcMethod:      "Denom",
					Use:            "show-denom [id]",
					Short:          "Shows a Denom",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "denom"}},
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateDenom",
					Use:            "create-denom [denom] [ticker] [precision] [maxSupply]",
					Short:          "Create a new Denom",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "denom"}, {ProtoField: "ticker"}, {ProtoField: "precision"}, {ProtoField: "maxSupply"}},
				},
				// {
				// 	RpcMethod:      "UpdateDenom",
				// 	Use:            "update-denom [denom] [ticker] [precision] [maxSupply]",
				// 	Short:          "Update Denom",
				// 	PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "denom"}, {ProtoField: "ticker"}, {ProtoField: "precision"}, {ProtoField: "maxSupply"}},
				// },
				// {
				// 	RpcMethod:      "DeleteDenom",
				// 	Use:            "delete-denom [denom]",
				// 	Short:          "Delete Denom",
				// 	PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "denom"}},
				// },
				{
					RpcMethod:      "UpdateOwner",
					Use:            "update-owner [denom] [new-owner]",
					Short:          "update the owner of denom, UpdateOwner tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "denom"}, {ProtoField: "newOwner"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
