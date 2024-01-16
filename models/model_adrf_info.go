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



// Information of an ADRF NF Instance
type AdrfInfo struct {
	MlModelStorageCapability bool `json:"mlModelStorageCapability,omitempty" yaml:"mlModelStorageCapability" bson:"mlModelStorageCapability"`
}
