package v1_test

import (
	"fmt"
	"testing"
	"time"

	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/core/header"
	"cosmossdk.io/x/upgrade"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/babylonchain/babylon/app"
	v1 "github.com/babylonchain/babylon/app/upgrades/v1"
	bstypes "github.com/babylonchain/babylon/x/btcstaking/types"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
)

const (
	DummyUpgradeHeight = 5
)

type UpgradeTestSuite struct {
	suite.Suite

	ctx       sdk.Context
	app       *app.BabylonApp
	preModule appmodule.HasPreBlocker
}

func (s *UpgradeTestSuite) SetupTest() {
	s.app = app.Setup(s.T(), false)
	s.ctx = s.app.BaseApp.NewContextLegacy(false, tmproto.Header{Height: 1, ChainID: "babylon-1", Time: time.Now().UTC()})
	s.preModule = upgrade.NewAppModule(s.app.UpgradeKeeper, s.app.AccountKeeper.AddressCodec())
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(UpgradeTestSuite))
}

func (s *UpgradeTestSuite) TestUpgradePayments() {
	oldAcctNum := 0

	testCases := []struct {
		msg         string
		pre_update  func()
		update      func()
		post_update func()
		expPass     bool
	}{
		{
			"Test vanilla software upgrade gov prop",
			func() {
				allAccounts := s.app.AccountKeeper.GetAllAccounts(s.ctx)
				oldAcctNum = len(allAccounts)
			},
			func() {
				// inject upgrade plan
				s.ctx = s.ctx.WithBlockHeight(DummyUpgradeHeight - 1)
				plan := upgradetypes.Plan{Name: v1.Upgrade.UpgradeName, Height: DummyUpgradeHeight}
				s.app.UpgradeKeeper.SetUpgradeHandler(
					v1.Upgrade.UpgradeName,
					v1.Upgrade.CreateUpgradeHandler(
						s.app.ModuleManager,
						nil,
						nil,
						s.app.AppKeepers,
					),
				)

				// run upgrade
				err := s.app.UpgradeKeeper.ScheduleUpgrade(s.ctx, plan)
				s.NoError(err)
				_, err = s.app.UpgradeKeeper.GetUpgradePlan(s.ctx)
				s.NoError(err)

				s.ctx = s.ctx.WithHeaderInfo(header.Info{Height: DummyUpgradeHeight, Time: s.ctx.BlockTime().Add(time.Second)}).WithBlockHeight(DummyUpgradeHeight)
				s.NotPanics(func() {
					_, err := s.preModule.PreBlock(s.ctx)
					s.NoError(err)
				})
			},
			func() {
				// ensure the account is removed
				allAccounts := s.app.AccountKeeper.GetAllAccounts(s.ctx)
				newAcctNum := len(allAccounts)
				s.Equal(newAcctNum, oldAcctNum-1)

				// ensure finality provider is inserted
				resp, err := s.app.BTCStakingKeeper.FinalityProviders(s.ctx, &bstypes.QueryFinalityProvidersRequest{})
				s.NoError(err)
				s.Len(resp.FinalityProviders, 1)
			},
			true,
		},
	}

	for _, tc := range testCases {
		s.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			s.SetupTest() // reset

			tc.pre_update()
			tc.update()
			tc.post_update()
		})
	}
}
