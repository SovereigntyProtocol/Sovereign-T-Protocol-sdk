package keeper_test

import (
	"strconv"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "github.com/cosmos/cosmos-sdk/testutil/keeper"
	"github.com/cosmos/cosmos-sdk/x/identity/keeper"
	"github.com/cosmos/cosmos-sdk/x/identity/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestIdMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.IdentityKeeper(t)
	srv := keeper.NewMsgServerImpl(k)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateId{Creator: creator,
			Did: strconv.Itoa(i),
		}
		_, err := srv.CreateId(ctx, expected)
		require.NoError(t, err)
		rst, found := k.GetId(ctx,
			expected.Did,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestIdMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateId
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateId{Creator: creator,
				Did: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateId{Creator: "B",
				Did: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateId{Creator: creator,
				Did: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.IdentityKeeper(t)
			srv := keeper.NewMsgServerImpl(k)
			expected := &types.MsgCreateId{Creator: creator,
				Did: strconv.Itoa(0),
			}
			_, err := srv.CreateId(ctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateId(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetId(ctx,
					expected.Did,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestIdMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteId
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteId{Creator: creator,
				Did: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteId{Creator: "B",
				Did: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteId{Creator: creator,
				Did: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.IdentityKeeper(t)
			srv := keeper.NewMsgServerImpl(k)

			_, err := srv.CreateId(ctx, &types.MsgCreateId{Creator: creator,
				Did: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteId(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetId(ctx,
					tc.request.Did,
				)
				require.False(t, found)
			}
		})
	}
}
