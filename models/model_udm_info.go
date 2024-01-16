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



// Information of an UDM NF Instance
type UdmInfo struct {
	// Identifier of a group of NFs.
	GroupId string `json:"groupId,omitempty" yaml:"groupId" bson:"groupId"`
	SupiRanges []*SupiRange `json:"supiRanges,omitempty" yaml:"supiRanges" bson:"supiRanges"`
	GpsiRanges []*IdentityRange `json:"gpsiRanges,omitempty" yaml:"gpsiRanges" bson:"gpsiRanges"`
	ExternalGroupIdentifiersRanges []*IdentityRange `json:"externalGroupIdentifiersRanges,omitempty" yaml:"externalGroupIdentifiersRanges" bson:"externalGroupIdentifiersRanges"`
	RoutingIndicators []*string `json:"routingIndicators,omitempty" yaml:"routingIndicators" bson:"routingIndicators"`
	InternalGroupIdentifiersRanges []*InternalGroupIdRange `json:"internalGroupIdentifiersRanges,omitempty" yaml:"internalGroupIdentifiersRanges" bson:"internalGroupIdentifiersRanges"`
	SuciInfos []*SuciInfo `json:"suciInfos,omitempty" yaml:"suciInfos" bson:"suciInfos"`
}
