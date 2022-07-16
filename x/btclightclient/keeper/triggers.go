package keeper

import (
	"github.com/babylonchain/babylon/x/btclightclient/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) triggerRollBack(ctx sdk.Context, headerInfo *types.BTCHeaderInfo) {
	// Trigger AfterBTCRollBack hook
	k.AfterBTCRollBack(ctx, headerInfo)
	// Emit BTCRollBack event
	ctx.EventManager().EmitTypedEvent(&types.EventBTCRollBack{Header: headerInfo})
}

func (k Keeper) triggerRollForward(ctx sdk.Context, headerInfo *types.BTCHeaderInfo) {
	// Trigger AfterBTCRollForward hook
	k.AfterBTCRollForward(ctx, headerInfo)
	// Emit BTCRollForward event
	ctx.EventManager().EmitTypedEvent(&types.EventBTCRollForward{Header: headerInfo})
}