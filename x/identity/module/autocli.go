package identity

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/cosmos/cosmos-sdk/ssiapi/identity"
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
					Skip:      true,
				},
				{
					RpcMethod: "IdAll",
					Use:       "list-id",
					Short:     "List all ID",
				},
				{
					RpcMethod:      "Id",
					Use:            "show-id [id]",
					Short:          "Shows a ID",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "did"}},
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
					RpcMethod:      "CreateId",
					Use:            "create-id [hash] [username]",
					Short:          "Create a new ID",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "hash"}, {ProtoField: "username"}},
				},
				{
					RpcMethod:      "UpdateId",
					Use:            "update-id [did] [hash] [owner] ",
					Short:          "Update ID",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "did"}, {ProtoField: "hash"}, {ProtoField: "owner"}},
				},
				{
					RpcMethod:      "DeleteId",
					Use:            "delete-id [did]",
					Short:          "Delete ID",
					Skip:           true,
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "did"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
