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



// Information of an HSS NF Instance
type HssInfo struct {
	// Identifier of a group of NFs.
	GroupId string `json:"groupId,omitempty" yaml:"groupId" bson:"groupId"`
	ImsiRanges []*ImsiRange `json:"imsiRanges,omitempty" yaml:"imsiRanges" bson:"imsiRanges"`
	ImsPrivateIdentityRanges []*IdentityRange `json:"imsPrivateIdentityRanges,omitempty" yaml:"imsPrivateIdentityRanges" bson:"imsPrivateIdentityRanges"`
	ImsPublicIdentityRanges []*IdentityRange `json:"imsPublicIdentityRanges,omitempty" yaml:"imsPublicIdentityRanges" bson:"imsPublicIdentityRanges"`
	MsisdnRanges []*IdentityRange `json:"msisdnRanges,omitempty" yaml:"msisdnRanges" bson:"msisdnRanges"`
	ExternalGroupIdentifiersRanges []*IdentityRange `json:"externalGroupIdentifiersRanges,omitempty" yaml:"externalGroupIdentifiersRanges" bson:"externalGroupIdentifiersRanges"`
	HssDiameterAddress *NetworkNodeDiameterAddress `json:"hssDiameterAddress,omitempty" yaml:"hssDiameterAddress" bson:"hssDiameterAddress"`
}
