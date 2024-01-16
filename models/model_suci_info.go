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



// SUCI information containing Routing Indicator and Home Network Public Key ID
type SuciInfo struct {
	RoutingInds []*string `json:"routingInds,omitempty" yaml:"routingInds" bson:"routingInds"`
	HNwPubKeyIds []*int32 `json:"hNwPubKeyIds,omitempty" yaml:"hNwPubKeyIds" bson:"hNwPubKeyIds"`
}
