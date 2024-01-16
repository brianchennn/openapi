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



// Range of TMGIs
type TmgiRange struct {
	MbsServiceIdStart string `json:"mbsServiceIdStart" yaml:"mbsServiceIdStart" bson:"mbsServiceIdStart"`
	MbsServiceIdEnd string `json:"mbsServiceIdEnd" yaml:"mbsServiceIdEnd" bson:"mbsServiceIdEnd"`
	PlmnId *PlmnId `json:"plmnId" yaml:"plmnId" bson:"plmnId"`
	// This represents the Network Identifier, which together with a PLMN ID is used to identify an SNPN (see 3GPP TS 23.003 and 3GPP TS 23.501 clause 5.30.2.1).  
	Nid string `json:"nid,omitempty" yaml:"nid" bson:"nid"`
}
