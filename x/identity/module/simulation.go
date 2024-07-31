package identity

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/cosmos/cosmos-sdk/testutil/sample"
	identitysimulation "github.com/cosmos/cosmos-sdk/x/identity/simulation"
	"github.com/cosmos/cosmos-sdk/x/identity/types"
)

// avoid unused import issue
var (
	_ = identitysimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateUser = "op_weight_msg_user"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateUser int = 100

	opWeightMsgUpdateUser = "op_weight_msg_user"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateUser int = 100

	opWeightMsgDeleteUser = "op_weight_msg_user"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteUser int = 100

	opWeightMsgCreateAddress = "op_weight_msg_address"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateAddress int = 100

	opWeightMsgUpdateAddress = "op_weight_msg_address"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateAddress int = 100

	opWeightMsgDeleteAddress = "op_weight_msg_address"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteAddress int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	identityGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		UserList: []types.User{
			{
				Creator: sample.AccAddress(),
				Did:     "0",
			},
			{
				Creator: sample.AccAddress(),
				Did:     "1",
			},
		},
		AddressList: []types.Address{
			{
				Creator: sample.AccAddress(),
				Owner:   "0",
			},
			{
				Creator: sample.AccAddress(),
				Owner:   "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&identityGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateUser int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateUser, &weightMsgCreateUser, nil,
		func(_ *rand.Rand) {
			weightMsgCreateUser = defaultWeightMsgCreateUser
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateUser,
		identitysimulation.SimulateMsgCreateUser(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateUser int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateUser, &weightMsgUpdateUser, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateUser = defaultWeightMsgUpdateUser
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateUser,
		identitysimulation.SimulateMsgUpdateUser(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteUser int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteUser, &weightMsgDeleteUser, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteUser = defaultWeightMsgDeleteUser
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteUser,
		identitysimulation.SimulateMsgDeleteUser(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateAddress int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateAddress, &weightMsgCreateAddress, nil,
		func(_ *rand.Rand) {
			weightMsgCreateAddress = defaultWeightMsgCreateAddress
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateAddress,
		identitysimulation.SimulateMsgCreateAddress(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateAddress int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateAddress, &weightMsgUpdateAddress, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateAddress = defaultWeightMsgUpdateAddress
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateAddress,
		identitysimulation.SimulateMsgUpdateAddress(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteAddress int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteAddress, &weightMsgDeleteAddress, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteAddress = defaultWeightMsgDeleteAddress
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteAddress,
		identitysimulation.SimulateMsgDeleteAddress(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateUser,
			defaultWeightMsgCreateUser,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				identitysimulation.SimulateMsgCreateUser(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateUser,
			defaultWeightMsgUpdateUser,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				identitysimulation.SimulateMsgUpdateUser(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteUser,
			defaultWeightMsgDeleteUser,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				identitysimulation.SimulateMsgDeleteUser(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateAddress,
			defaultWeightMsgCreateAddress,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				identitysimulation.SimulateMsgCreateAddress(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateAddress,
			defaultWeightMsgUpdateAddress,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				identitysimulation.SimulateMsgUpdateAddress(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteAddress,
			defaultWeightMsgDeleteAddress,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				identitysimulation.SimulateMsgDeleteAddress(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
