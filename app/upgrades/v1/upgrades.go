package v1

import (
	"context"

	store "cosmossdk.io/store/types"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/babylonchain/babylon/app/keepers"
	"github.com/babylonchain/babylon/app/upgrades"
	bbn "github.com/babylonchain/babylon/types"
	btcstakingkeeper "github.com/babylonchain/babylon/x/btcstaking/keeper"
	bstypes "github.com/babylonchain/babylon/x/btcstaking/types"
	"github.com/btcsuite/btcd/btcec/v2"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
)

var Upgrade = upgrades.Upgrade{
	UpgradeName:          "vanilla",
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades:        store.StoreUpgrades{},
}

func CreateUpgradeHandler(
	mm *module.Manager,
	_ module.Configurator,
	_ upgrades.BaseAppParamManager,
	keepers *keepers.AppKeepers,
) upgradetypes.UpgradeHandler {
	return func(context context.Context, _plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		ctx := sdk.UnwrapSDKContext(context)

		propVanilla(ctx, &keepers.AccountKeeper, &keepers.BTCStakingKeeper)

		return vm, nil
	}
}

func propVanilla(
	ctx sdk.Context,
	accountKeeper *authkeeper.AccountKeeper,
	bsKeeper *btcstakingkeeper.Keeper,
) {
	// remove an account
	allAccounts := accountKeeper.GetAllAccounts(ctx)
	accountKeeper.RemoveAccount(ctx, allAccounts[len(allAccounts)-1])

	// insert a FP
	sk, err := btcec.NewPrivateKey()
	if err != nil {
		panic(err)
	}
	btcPK := bbn.NewBIP340PubKeyFromBTCPK(sk.PubKey())
	fp := &bstypes.FinalityProvider{
		Addr:  allAccounts[0].GetAddress().String(),
		BtcPk: btcPK,
	}
	bsKeeper.SetFinalityProvider(ctx, fp)
}
