/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

// Package holder enables the Holder: an entity that receives SD-JWTs from the Issuer and has control over them.
package holder

import (
	"github.com/hyperledger/aries-framework-go/component/models/sdjwt/holder"

	"github.com/hyperledger/aries-framework-go/component/kmscrypto/doc/jose"
)

// Claim defines claim.
type Claim = holder.Claim

// ParseOpt is the SD-JWT Parser option.
type ParseOpt = holder.ParseOpt

// WithJWTDetachedPayload option is for definition of JWT detached payload.
func WithJWTDetachedPayload(payload []byte) ParseOpt {
	return holder.WithJWTDetachedPayload(payload)
}

// WithSignatureVerifier option is for definition of JWT detached payload.
func WithSignatureVerifier(signatureVerifier jose.SignatureVerifier) ParseOpt {
	return holder.WithSignatureVerifier(signatureVerifier)
}

// Parse parses issuer SD-JWT and returns claims that can be selected.
// The Holder MUST perform the following (or equivalent) steps when receiving a Combined Format for Issuance:
//
//   - Separate the SD-JWT and the Disclosures in the Combined Format for Issuance.
//
//   - Hash all the Disclosures separately.
//
//   - Find the places in the SD-JWT where the digests of the Disclosures are included.
//
//   - If any of the digests cannot be found in the SD-JWT, the Holder MUST reject the SD-JWT.
//
//   - Decode Disclosures and obtain plaintext of the claim values.
//
//     It is up to the Holder how to maintain the mapping between the Disclosures and the plaintext claim values to
//     be able to display them to the End-User when needed.
func Parse(combinedFormatForIssuance string, opts ...ParseOpt) ([]*Claim, error) {
	return holder.Parse(combinedFormatForIssuance, opts...)
}

// BindingPayload represents holder binding payload.
type BindingPayload = holder.BindingPayload

// BindingInfo defines holder binding payload and signer.
type BindingInfo = holder.BindingInfo

// Option is a holder option.
type Option = holder.Option

// WithHolderBinding option to set optional holder binding.
func WithHolderBinding(info *BindingInfo) Option {
	return holder.WithHolderBinding(info)
}

// CreatePresentation is a convenience method to assemble combined format for presentation
// using selected disclosures (claimsToDisclose) and optional holder binding.
// This call assumes that combinedFormatForIssuance has already been parsed and verified using Parse() function.
//
// For presentation to a Verifier, the Holder MUST perform the following (or equivalent) steps:
//   - Decide which Disclosures to release to the Verifier, obtaining proper End-User consent if necessary.
//   - If Holder Binding is required, create a Holder Binding JWT.
//   - Create the Combined Format for Presentation from selected Disclosures and Holder Binding JWT(if applicable).
//   - Send the Presentation to the Verifier.
func CreatePresentation(combinedFormatForIssuance string, claimsToDisclose []string, opts ...Option) (string, error) {
	return holder.CreatePresentation(combinedFormatForIssuance, claimsToDisclose, opts...)
}

// CreateHolderBinding will create holder binding from binding info.
func CreateHolderBinding(info *BindingInfo) (string, error) {
	return holder.CreateHolderBinding(info)
}

// NoopSignatureVerifier is no-op signature verifier (signature will not get checked).
type NoopSignatureVerifier = holder.NoopSignatureVerifier
