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



// Subscription to a set of NFs based on their Group Ids
type NfGroupListCond struct {
	ConditionType string `json:"conditionType" yaml:"conditionType" bson:"conditionType"`
	NfType string `json:"nfType" yaml:"nfType" bson:"nfType"`
	NfGroupIdList []*string `json:"nfGroupIdList" yaml:"nfGroupIdList" bson:"nfGroupIdList"`
}