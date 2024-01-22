/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * Source file: 3GPP TS 29.510 V18.4.0; 5G System; Network Function Repository Services; Stage 3
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.510/
 *
 * API version: 1.3.0-alpha.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models



type RuleSet struct {
	Priority int32 `json:"priority" yaml:"priority" bson:"priority"`
	Plmns []*PlmnId `json:"plmns,omitempty" yaml:"plmns" bson:"plmns"`
	Snpns []*PlmnIdNid `json:"snpns,omitempty" yaml:"snpns" bson:"snpns"`
	NfTypes []*NfType `json:"nfTypes,omitempty" yaml:"nfTypes" bson:"nfTypes"`
	NfDomains []*string `json:"nfDomains,omitempty" yaml:"nfDomains" bson:"nfDomains"`
	Nssais []*ExtSnssai `json:"nssais,omitempty" yaml:"nssais" bson:"nssais"`
	NfInstances []*string `json:"nfInstances,omitempty" yaml:"nfInstances" bson:"nfInstances"`
	Scopes []*string `json:"scopes,omitempty" yaml:"scopes" bson:"scopes"`
	Action RuleSetAction `json:"action" yaml:"action" bson:"action"`
}