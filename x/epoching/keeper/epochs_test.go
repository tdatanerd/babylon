package keeper_test

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/babylonchain/babylon/testutil/datagen"
	testhelper "github.com/babylonchain/babylon/testutil/helper"
	"github.com/babylonchain/babylon/x/epoching/types"
)

func FuzzEpochs(f *testing.F) {
	datagen.AddRandomSeedsToFuzzer(f, 10)

	f.Fuzz(func(t *testing.T, seed int64) {
		r := rand.New(rand.NewSource(seed))

		helper := testhelper.NewHelper(t)
		ctx, keeper := helper.Ctx, helper.App.EpochingKeeper
		// ensure that the epoch info is correct at the genesis
		epoch := keeper.GetEpoch(ctx)
		require.Equal(t, epoch.EpochNumber, uint64(1))
		require.Equal(t, epoch.FirstBlockHeight, uint64(1))

		epochInterval := keeper.GetParams(ctx).EpochInterval

		// increment a random number of new blocks
		numIncBlocks := r.Uint64()%100 + 1
		var err error
		for i := uint64(0); i < numIncBlocks-1; i++ {
			// TODO: Figure out why when ctx height is 1, ApplyEmptyBlockWithVoteExtension
			// will still give ctx height 1 once, then start to increment
			ctx, err = helper.ApplyEmptyBlockWithVoteExtension(r)
			require.NoError(t, err)
		}

		// ensure that the epoch info is still correct
		expectedEpochNumber := (numIncBlocks + 1) / epochInterval
		if (numIncBlocks+1)%epochInterval > 0 {
			expectedEpochNumber += 1
		}
		actualNewEpoch := keeper.GetEpoch(ctx)
		require.Equal(t, expectedEpochNumber, actualNewEpoch.EpochNumber)
		require.Equal(t, epochInterval, actualNewEpoch.CurrentEpochInterval)
		require.Equal(t, (expectedEpochNumber-1)*epochInterval+1, actualNewEpoch.FirstBlockHeight)
	})
}

func FuzzEpochs_UpdateEpochInterval(f *testing.F) {
	datagen.AddRandomSeedsToFuzzer(f, 10)

	f.Fuzz(func(t *testing.T, seed int64) {
		r := rand.New(rand.NewSource(seed))

		h := testhelper.NewHelper(t)
		keeper := h.App.EpochingKeeper

		// increment a random number of new blocks
		numIncBlocks := r.Uint64()%100 + 1
		var err error
		for i := uint64(0); i < numIncBlocks-1; i++ {
			// When ctx height is 1, ApplyEmptyBlockWithVoteExtension
			// will still give ctx height 1 once, then start to increment
			_, err = h.ApplyEmptyBlockWithVoteExtension(r)
			require.NoError(t, err)
		}
		// get current epoch metadata
		epoch := keeper.GetEpoch(h.Ctx)

		// update the epoch interval in params
		newEpochInterval := datagen.RandomInt(r, 20) + 2
		newParams := types.Params{EpochInterval: newEpochInterval}
		err = keeper.SetParams(h.Ctx, newParams)
		require.NoError(t, err)

		// ensure the current epoch metadata is not affected
		epoch2 := keeper.GetEpoch(h.Ctx)
		require.Equal(t, epoch, epoch2)

		// enter the last block of the current epoch
		lastHeightOfEpoch := epoch.GetLastBlockHeight()
		for uint64(h.Ctx.HeaderInfo().Height) < lastHeightOfEpoch {
			_, err = h.ApplyEmptyBlockWithVoteExtension(r)
			require.NoError(t, err)
		}
		// enter the next block and thus 1st block of the next epoch
		_, err = h.ApplyEmptyBlockWithVoteExtension(r)
		require.NoError(t, err)
		// ensure
		// - the epoch has incremented
		// - epoch interval is updated
		// - first/last height of the epoch is correct
		newEpoch := keeper.GetEpoch(h.Ctx)
		require.Equal(t, epoch.EpochNumber+1, newEpoch.EpochNumber)
		require.Equal(t, newEpochInterval, newEpoch.CurrentEpochInterval)
		require.Equal(t, uint64(h.Ctx.HeaderInfo().Height), newEpoch.FirstBlockHeight)
		require.Equal(t, uint64(h.Ctx.HeaderInfo().Height)+newEpochInterval-1, newEpoch.GetLastBlockHeight())
	})
}
