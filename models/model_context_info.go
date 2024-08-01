/*
 * Nudm_UECM
 *
 * Nudm Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * Source file: 3GPP TS 29.503 Unified Data Management Services, version 17.7.0
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.503/
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models



type ContextInfo struct {
	OrigHeaders []*string `json:"origHeaders,omitempty" yaml:"origHeaders" bson:"origHeaders"`
	RequestHeaders []*string `json:"requestHeaders,omitempty" yaml:"requestHeaders" bson:"requestHeaders"`
}
