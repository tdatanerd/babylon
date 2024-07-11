package containers

// ImageConfig contains all images and their respective tags
// needed for running e2e tests.
type ImageConfig struct {
	InitRepository string
	InitTag        string

	BabylonRepository string
	BabylonTag        string

	RelayerRepository string
	RelayerTag        string
}

//nolint:deadcode
const (
	// name of babylon container produced by running `make localnet-build-env`
	CurrentBranchBabylonRepository = "babylonchain/babylond"
	CurrentBranchBabylonTag        = "latest"

	// Pre-upgrade babylon repo/tag to pull
	// TODO: replace with the prior version
	previousVersionBabylonRepository = "babylonchain/babylond"
	previousVersionBabylonTag        = "latest"
	// Pre-upgrade repo/tag for initialization (this should be one version below upgradeVersion)
	// TODO: replace with the prior version
	previousVersionInitRepository = "babylonchain/babylond"
	previousVersionInitTag        = "latest"

	hermesRelayerRepository = "informalsystems/hermes"
	hermesRelayerTag        = "v1.8.2"
	// Built using the `build-cosmos-relayer-docker` target on an Intel (amd64) machine and pushed to ECR
	cosmosRelayerRepository = "public.ecr.aws/t9e9i3h0/cosmos-relayer"
	// TODO: Replace with version tag once we have a working version
	cosmosRelayerTag = "main"
)

// NewImageConfig returns ImageConfig needed for running e2e test.
// If isUpgrade is true, returns images for running the upgrade
// If isFork is true, utilizes provided fork height to initiate fork logic
func NewImageConfig(isCosmosRelayer bool, isUpgrade bool, isFork bool) ImageConfig {
	config := ImageConfig{}

	// set relayer image name / tag
	if isCosmosRelayer {
		config.RelayerRepository = cosmosRelayerRepository
		config.RelayerTag = cosmosRelayerTag
	} else {
		config.RelayerRepository = hermesRelayerRepository
		config.RelayerTag = hermesRelayerTag
	}

	if !isUpgrade {
		// If upgrade is not tested, we do not need InitRepository and InitTag
		// because we directly call the initialization logic without
		// the need for Docker.
		config.BabylonRepository = CurrentBranchBabylonRepository
		config.BabylonTag = CurrentBranchBabylonTag
		return config
	}

	if isFork {
		// Forks are state compatible with earlier versions before fork height.
		// Normally, validators switch the binaries pre-fork height
		// Then, once the fork height is reached, the state breaking-logic
		// is run.
		config.BabylonRepository = CurrentBranchBabylonRepository
		config.BabylonTag = CurrentBranchBabylonTag
	} else {
		// Upgrades are run at the time when upgrade height is reached
		// and are submitted via a governance proposal. Therefore, we
		// must start running the previous Osmosis version. Then, the node
		// should auto-upgrade, at which point we can restart the updated
		// Osmosis validator container.
		config.BabylonRepository = previousVersionBabylonRepository
		config.BabylonTag = previousVersionBabylonTag
	}

	// If upgrade is tested, we need to utilize InitRepository and InitTag
	// to initialize older state with Docker
	config.InitRepository = previousVersionInitRepository
	config.InitTag = previousVersionInitTag

	return config
}
