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



// NSACF service capabilities (e.g. to monitor and control the number of registered UEs or established PDU sessions per network slice) 
type NsacfCapability struct {
	// Indicates the service capability of the NSACF to monitor and control the number of registered UEs per network slice for the network slice that is subject to NSAC   true: Supported   false (default): Not Supported 
	SupportUeSAC bool `json:"supportUeSAC,omitempty" yaml:"supportUeSAC" bson:"supportUeSAC"`
	// Indicates the service capability of the NSACF to monitor and control the number of established PDU sessions per network slice for the network slice that is subject to NSAC   true: Supported   false (default): Not Supported 
	SupportPduSAC bool `json:"supportPduSAC,omitempty" yaml:"supportPduSAC" bson:"supportPduSAC"`
}