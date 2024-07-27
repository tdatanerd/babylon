//go:build e2e

package app

import "github.com/babylonchain/babylon/app/upgrades/vanilla"

func init() {
	Upgrades = append(Upgrades, vanilla.Upgrade)
}
