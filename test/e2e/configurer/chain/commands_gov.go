package chain

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/babylonchain/babylon/test/e2e/util"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govv1types "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	"github.com/stretchr/testify/require"
)

func (n *NodeConfig) QueryGovModuleAccount() string {
	cmd := []string{"babylond", "query", "auth", "module-accounts", "--output=json"}

	out, _, err := n.containerManager.ExecCmd(n.t, n.Name, cmd, "")
	require.NoError(n.t, err)
	var result map[string][]interface{}
	err = json.Unmarshal(out.Bytes(), &result)
	require.NoError(n.t, err)
	for _, acc := range result["accounts"] {
		account, ok := acc.(map[string]interface{})
		require.True(n.t, ok)
		if account["name"] == "gov" {
			moduleAccount, ok := account["base_account"].(map[string]interface{})["address"].(string)
			require.True(n.t, ok)
			return moduleAccount
		}
	}
	require.True(n.t, false, "gov module account not found")
	return ""
}

func (n *NodeConfig) QueryGovParams() *govv1types.Params {
	bz, err := n.QueryGRPCGateway("/babylon/btcstaking/v1/params", url.Values{})
	require.NoError(n.t, err)

	var resp govv1types.QueryParamsResponse
	err = util.Cdc.UnmarshalJSON(bz, &resp)
	require.NoError(n.t, err)

	return resp.Params
}

func (n *NodeConfig) SubmitParamChangeProposal(proposalJson, from string, isLegacy bool) int {
	n.LogActionF("submitting param change proposal %s", proposalJson)
	// ToDo: Is there a better way to do this?
	wd, err := os.Getwd()
	require.NoError(n.t, err)
	currentTime := time.Now().Format("20060102-150405.000")
	localProposalFile := wd + fmt.Sprintf("/scripts/param_change_proposal_%s.json", currentTime)
	f, err := os.Create(localProposalFile)
	require.NoError(n.t, err)
	_, err = f.WriteString(proposalJson)
	require.NoError(n.t, err)
	err = f.Close()
	require.NoError(n.t, err)

	var cmd []string
	if isLegacy {
		cmd = []string{"babylond", "tx", "gov", "submit-legacy-proposal", "param-change", fmt.Sprintf("/babylon/param_change_proposal_%s.json", currentTime), fmt.Sprintf("--from=%s", from)}
	} else {
		cmd = []string{"babylond", "tx", "gov", "submit-proposal", "param-change", fmt.Sprintf("/babylon/param_change_proposal_%s.json", currentTime), fmt.Sprintf("--from=%s", from)}
	}

	resp, _, err := n.containerManager.ExecTxCmd(n.t, n.chainId, n.Name, cmd)
	require.NoError(n.t, err)

	os.Remove(localProposalFile)

	proposalID, err := extractProposalIdFromResponse(resp.String())
	require.NoError(n.t, err)

	n.LogActionF("successfully submitted param change proposal")

	return proposalID
}

func (n *NodeConfig) SubmitNewV1ProposalType(proposalJson, from string) int {
	n.LogActionF("submitting new v1 proposal type %s", proposalJson)
	// ToDo: Is there a better way to do this?
	wd, err := os.Getwd()
	require.NoError(n.t, err)
	currentTime := time.Now().Format("20060102-150405.000")
	localProposalFile := wd + fmt.Sprintf("/scripts/new_v1_prop_%s.json", currentTime)
	f, err := os.Create(localProposalFile)
	require.NoError(n.t, err)
	_, err = f.WriteString(proposalJson)
	require.NoError(n.t, err)
	err = f.Close()
	require.NoError(n.t, err)

	cmd := []string{"babylond", "tx", "gov", "submit-proposal", fmt.Sprintf("/osmosis/new_v1_prop_%s.json", currentTime), fmt.Sprintf("--from=%s", from)}

	resp, _, err := n.containerManager.ExecTxCmd(n.t, n.chainId, n.Name, cmd)
	require.NoError(n.t, err)

	os.Remove(localProposalFile)

	proposalID, err := extractProposalIdFromResponse(resp.String())
	require.NoError(n.t, err)

	n.LogActionF("successfully submitted new v1 proposal type")

	return proposalID
}

func (n *NodeConfig) SubmitProposal(cmdArgs []string, isExpedited bool, propDescriptionForLogs string, isLegacy bool) int {
	n.LogActionF("submitting proposal: %s", propDescriptionForLogs)
	var cmd []string
	if isLegacy {
		cmd = append([]string{"babylond", "tx", "gov", "submit-legacy-proposal"}, cmdArgs...)
	} else {
		cmd = append([]string{"babylond", "tx", "gov", "submit-proposal"}, cmdArgs...)
	}

	param := n.QueryGovParams()

	depositAmt := param.MinDeposit
	if isExpedited {
		cmd = append(cmd, "--is-expedited=true")
		depositAmt = param.ExpeditedMinDeposit
	}
	depositAmtStr := sdk.NewCoins(depositAmt...).String()
	cmd = append(cmd, fmt.Sprintf("--deposit=%s", depositAmtStr))

	resp, _, err := n.containerManager.ExecTxCmd(n.t, n.chainId, n.Name, cmd)
	require.NoError(n.t, err)

	proposalID, err := extractProposalIdFromResponse(resp.String())
	require.NoError(n.t, err)

	n.LogActionF("successfully submitted proposal: %s", propDescriptionForLogs)

	return proposalID
}

func (n *NodeConfig) SubmitUpgradeProposal(upgradeVersion string, upgradeHeight int64, initialDeposit sdk.Coin, isLegacy bool) int {
	cmd := []string{"software-upgrade", upgradeVersion, fmt.Sprintf("--title=\"%s upgrade\"", upgradeVersion), "--description=\"upgrade proposal submission\"", fmt.Sprintf("--upgrade-height=%d", upgradeHeight), "--upgrade-info=\"\"", "--no-validate", "--from=val"}
	return n.SubmitProposal(cmd, false, fmt.Sprintf("upgrade proposal %s for height %d", upgradeVersion, upgradeHeight), true)
}

func (n *NodeConfig) SubmitTextProposal(text string, isExpedited, isLegacy bool) int {
	cmd := []string{"--type=text", fmt.Sprintf("--title=\"%s\"", text), "--description=\"test text proposal\"", "--from=val"}
	return n.SubmitProposal(cmd, isExpedited, "text proposal", isLegacy)
}

func (n *NodeConfig) DepositProposal(proposalNumber int, isExpedited bool) {
	n.LogActionF("depositing on proposal: %d", proposalNumber)

	cmd := []string{"babylond", "tx", "gov", "deposit", fmt.Sprintf("%d", proposalNumber)}

	// set deposit amount
	param := n.QueryGovParams()
	depositAmt := param.MinDeposit
	if isExpedited {
		depositAmt = param.ExpeditedMinDeposit
	}
	depositAmtStr := sdk.NewCoins(depositAmt...).String()
	cmd = append(cmd, depositAmtStr)

	// set account
	cmd = append(cmd, "--from=val")

	_, _, err := n.containerManager.ExecTxCmd(n.t, n.chainId, n.Name, cmd)
	require.NoError(n.t, err)
	n.LogActionF("successfully deposited on proposal %d", proposalNumber)
}

func (n *NodeConfig) VoteYesProposal(from string, proposalNumber int) {
	n.LogActionF("voting yes on proposal: %d", proposalNumber)
	cmd := []string{"babylond", "tx", "gov", "vote", fmt.Sprintf("%d", proposalNumber), "yes", fmt.Sprintf("--from=%s", from)}
	_, _, err := n.containerManager.ExecTxCmd(n.t, n.chainId, n.Name, cmd)
	require.NoError(n.t, err)
	n.LogActionF("successfully voted yes on proposal %d", proposalNumber)
}

func (n *NodeConfig) VoteNoProposal(from string, proposalNumber int) {
	n.LogActionF("voting no on proposal: %d", proposalNumber)
	cmd := []string{"babylond", "tx", "gov", "vote", fmt.Sprintf("%d", proposalNumber), "no", fmt.Sprintf("--from=%s", from)}
	_, _, err := n.containerManager.ExecTxCmd(n.t, n.chainId, n.Name, cmd)
	require.NoError(n.t, err)
	n.LogActionF("successfully voted no on proposal: %d", proposalNumber)
}

func extractProposalIdFromResponse(response string) (int, error) {
	// Extract the proposal ID from the response
	startIndex := strings.Index(response, `[{"key":"proposal_id","value":"`) + len(`[{"key":"proposal_id","value":"`)
	endIndex := strings.Index(response[startIndex:], `"`)

	// Extract the proposal ID substring
	proposalIDStr := response[startIndex : startIndex+endIndex]

	// Convert the proposal ID from string to int
	proposalID, err := strconv.Atoi(proposalIDStr)
	if err != nil {
		return 0, err
	}

	return proposalID, nil
}
