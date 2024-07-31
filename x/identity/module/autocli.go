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
				},
				{
					RpcMethod: "UserAll",
					Use:       "list-user",
					Short:     "List all User",
				},
				{
					RpcMethod:      "User",
					Use:            "show-user [id]",
					Short:          "Shows a User",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "did"}},
				},
				{
					RpcMethod: "AddressAll",
					Use:       "list-address",
					Short:     "List all Address",
				},
				{
					RpcMethod:      "Address",
					Use:            "show-address [id]",
					Short:          "Shows a Address",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "owner"}},
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
					RpcMethod:      "CreateUser",
					Use:            "create-user [did] [hash] [owner]",
					Short:          "Create a new User",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "did"}, {ProtoField: "hash"}, {ProtoField: "owner"}},
				},
				{
					RpcMethod:      "UpdateUser",
					Use:            "update-user [did] [hash] [owner]",
					Short:          "Update User",
					Skip:           true,
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "did"}, {ProtoField: "hash"}, {ProtoField: "owner"}},
				},
				{
					RpcMethod:      "DeleteUser",
					Use:            "delete-user [did]",
					Short:          "Delete User",
					Skip:           true,
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "did"}},
				},
				{
					RpcMethod:      "CreateAddress",
					Use:            "create-address [owner]",
					Short:          "Create a new Address",
					Skip:           true,
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "owner"}},
				},
				{
					RpcMethod:      "UpdateAddress",
					Use:            "update-address [owner]",
					Short:          "Update Address",
					Skip:           true,
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "owner"}},
				},
				{
					RpcMethod:      "DeleteAddress",
					Use:            "delete-address [owner]",
					Short:          "Delete Address",
					Skip:           true,
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "owner"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
