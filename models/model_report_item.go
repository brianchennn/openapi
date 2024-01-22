/*
 * Nudsf_DataRepository
 *
 * Nudsf Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * Source file: 3GPP TS 29.598 UDSF Services, V17.6.0.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.598/
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models



// indicates performed modivications.
type ReportItem struct {
	// Contains a JSON pointer value (as defined in IETF RFC 6901) that references a  location of a resource to which the modification is subject. 
	Path string `json:"path" yaml:"path" bson:"path"`
	// A human-readable reason providing details on the reported modification failure.  The reason string should identify the operation that failed using the operation's  array index to assist in correlation of the invalid parameter with the failed  operation, e.g. \"Replacement value invalid for attribute (failed operation index= 4)\". 
	Reason string `json:"reason,omitempty" yaml:"reason" bson:"reason"`
}