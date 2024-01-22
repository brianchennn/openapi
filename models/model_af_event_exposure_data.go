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



// AF Event Exposure data managed by a given NEF Instance
type AfEventExposureData struct {
	AfEvents []*AfEvent `json:"afEvents" yaml:"afEvents" bson:"afEvents"`
	AfIds []*string `json:"afIds,omitempty" yaml:"afIds" bson:"afIds"`
	AppIds []*string `json:"appIds,omitempty" yaml:"appIds" bson:"appIds"`
}