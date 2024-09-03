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
	opWeightMsgCreateId = "op_weight_msg_id"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateId int = 100

	opWeightMsgUpdateId = "op_weight_msg_id"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateId int = 100

	opWeightMsgDeleteId = "op_weight_msg_id"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteId int = 100

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
		IdList: []types.Id{
			{
				Creator: sample.AccAddress(),
				Did:     "0",
			},
			{
				Creator: sample.AccAddress(),
				Did:     "1",
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

	var weightMsgCreateId int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateId, &weightMsgCreateId, nil,
		func(_ *rand.Rand) {
			weightMsgCreateId = defaultWeightMsgCreateId
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateId,
		identitysimulation.SimulateMsgCreateId(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateId int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateId, &weightMsgUpdateId, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateId = defaultWeightMsgUpdateId
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateId,
		identitysimulation.SimulateMsgUpdateId(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteId int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteId, &weightMsgDeleteId, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteId = defaultWeightMsgDeleteId
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteId,
		identitysimulation.SimulateMsgDeleteId(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateId,
			defaultWeightMsgCreateId,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				identitysimulation.SimulateMsgCreateId(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateId,
			defaultWeightMsgUpdateId,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				identitysimulation.SimulateMsgUpdateId(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteId,
			defaultWeightMsgDeleteId,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				identitysimulation.SimulateMsgDeleteId(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
