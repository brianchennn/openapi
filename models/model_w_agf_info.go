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



// Information of the W-AGF end-points
type WAgfInfo struct {
	Ipv4EndpointAddresses []*string `json:"ipv4EndpointAddresses,omitempty" yaml:"ipv4EndpointAddresses" bson:"ipv4EndpointAddresses"`
	Ipv6EndpointAddresses []*string `json:"ipv6EndpointAddresses,omitempty" yaml:"ipv6EndpointAddresses" bson:"ipv6EndpointAddresses"`
	// Fully Qualified Domain Name
	EndpointFqdn string `json:"endpointFqdn,omitempty" yaml:"endpointFqdn" bson:"endpointFqdn"`
}