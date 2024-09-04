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
					Skip:      true,
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "IdAll",
					Use:       "list-id",
					Short:     "List all Id",
				},
				{
					RpcMethod:      "Id",
					Use:            "show-id [id]",
					Short:          "Shows a Id",
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
					Short:          "Update Id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "did"}, {ProtoField: "hash"}, {ProtoField: "owner"}},
				},
				{
					RpcMethod:      "DeleteId",
					Use:            "delete-id [did]",
					Short:          "Delete Id",
					Skip:           true,
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "did"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
