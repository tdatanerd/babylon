package types

import "github.com/babylonchain/babylon/types"

func NewEventSlashedFinalityProvider(evidence *Evidence) *EventSlashedFinalityProvider {
	return &EventSlashedFinalityProvider{
		Evidence: evidence,
	}
}

func NewEventInactiveFinalityProviderDetected(fpPk *types.BIP340PubKey) *EventInactiveFinalityProviderDetected {
	return &EventInactiveFinalityProviderDetected{PublicKey: fpPk.MarshalHex()}
}

func NewEventInactiveFinalityProviderReverted(fpPk *types.BIP340PubKey) *EventInactiveFinalityProviderReverted {
	return &EventInactiveFinalityProviderReverted{PublicKey: fpPk.MarshalHex()}
}
