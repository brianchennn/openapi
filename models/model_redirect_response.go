/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * Source file: 3GPP TS 29.510 V17.11.0; 5G System; Network Function Repository Services; Stage 3
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.510/
 *
 * API version: 1.2.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models



// The response shall include a Location header field containing a different URI  (pointing to a different URI of an other service instance), or the same URI if a request  is redirected to the same target resource via a different SCP. 
type RedirectResponse struct {
	Cause string `json:"cause,omitempty" yaml:"cause" bson:"cause"`
	// String providing an URI formatted according to RFC 3986.
	TargetScp string `json:"targetScp,omitempty" yaml:"targetScp" bson:"targetScp"`
	// String providing an URI formatted according to RFC 3986.
	TargetSepp string `json:"targetSepp,omitempty" yaml:"targetSepp" bson:"targetSepp"`
}