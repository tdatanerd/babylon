package e2e

import (
	"github.com/stretchr/testify/suite"

	"github.com/babylonchain/babylon/test/e2e/configurer"
)

var (
// r = rand.New(rand.NewSource(time.Now().Unix()))
// net = &chaincfg.SimNetParams
// // finality provider
// fpBTCSK, _, _ = datagen.GenRandomBTCKeyPair(r)
// cacheFP       *bstypes.FinalityProvider
// // BTC delegation
// delBTCSK, delBTCPK, _ = datagen.GenRandomBTCKeyPair(r)
// // covenant
// covenantSKs, _, covenantQuorum = bstypes.DefaultCovenantCommittee()

// stakingValue = int64(2 * 10e8)
)

const (
	// Mount path in container is fmt.Sprintf("%s/upgrades:/upgrades", pwd)
	vanillaUpgradeFilePath = "/upgrades/vanilla.json"
)

type SoftwareUpgradeCurrentBranchTestSuite struct {
	suite.Suite

	configurer configurer.Configurer
}

func (s *SoftwareUpgradeCurrentBranchTestSuite) SetupSuite() {
	s.T().Log("setting up e2e integration test suite...")
	var err error

	// The e2e test flow is as follows:
	//
	// 1. Configure 1 chain with some validator nodes
	// 2. Execute various e2e tests
	s.configurer, err = configurer.NewSoftwareUpgradeTest(s.T(), true)
	s.NoError(err)
	err = s.configurer.ConfigureChains()
	s.NoError(err)
	err = s.configurer.RunSetup()
	s.NoError(err)
}

func (s *SoftwareUpgradeCurrentBranchTestSuite) TearDownSuite() {
	err := s.configurer.ClearResources()
	s.Require().NoError(err)
}

// Test1UpgradeVanilla is an end-to-end test for
// running a software upgrade proposal
func (s *SoftwareUpgradeCurrentBranchTestSuite) Test1UpgradeVanilla() {
	// chain is already start the chain with software upgrade available
	chainA := s.configurer.GetChainConfig(0)
	chainA.WaitUntilHeight(1)

	nonValidatorNode, err := chainA.GetNodeAtIndex(2)
	s.NoError(err)

	propID := nonValidatorNode.TxGovPropSubmitProposal(vanillaUpgradeFilePath, nonValidatorNode.WalletName)
	s.Equal(1, propID)
	// run software upgrade gov prop
	// waits for block to reach
	// verifies if vanilla update was done
}
