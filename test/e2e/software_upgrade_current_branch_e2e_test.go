package e2e

import (
	"fmt"

	govv1 "cosmossdk.io/api/cosmos/gov/v1"
	"github.com/stretchr/testify/suite"

	"github.com/babylonchain/babylon/test/e2e/configurer"
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

	fpsBeforeUpgrade := nonValidatorNode.QueryFinalityProviders()

	// software upgrade gov prop
	propID := nonValidatorNode.TxGovPropSubmitProposal(vanillaUpgradeFilePath, nonValidatorNode.WalletName)
	s.Equal(1, propID)

	// vote from all nodes
	chainA.TxGovVoteFromAllNodes(propID, govv1.VoteOption_VOTE_OPTION_YES)

	tx := nonValidatorNode.QueryProposal(propID)
	fmt.Printf("\n prop %+v", tx)

	// waits for block to reach from plan + 1
	// tricky to get current heigth and set it in the json, because the file is
	// load at the mounting point of the node it could be created at runtime and
	// stored in the filesystem of the container
	nonValidatorNode.WaitForBlockHeight(21)

	tx = nonValidatorNode.QueryProposal(propID)
	fmt.Printf("\n prop %+v", tx)

	// verifies vanilla upgrade was completed
	fpsAfterUpgrade := nonValidatorNode.QueryFinalityProviders()
	s.Equal(len(fpsBeforeUpgrade)+1, len(fpsAfterUpgrade))

	// docker logs -f
	// [90m2:09AM[0m [31mERR[0m [1mBINARY UPDATED BEFORE TRIGGER! UPGRADE "vanilla" - in binary but not executed on chain. Downgrade your binary[0m [36mmodule=[0mx/upgrade
	// [90m2:09AM[0m [31mERR[0m [1merror in proxyAppConn.FinalizeBlock[0m [36merr=[0m"BINARY UPDATED BEFORE TRIGGER! UPGRADE \"vanilla\" - in binary but not executed on chain. Downgrade your binary" [36mmodule=[0mstate
	tx = nonValidatorNode.QueryProposal(propID)
	fmt.Printf("\n prop %+v", tx)

}
