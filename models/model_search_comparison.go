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



// A comparison to apply on tag/values pairs.
type SearchComparison struct {
	Op ComparisonOperator `json:"op" yaml:"op" bson:"op"`
	Tag string `json:"tag" yaml:"tag" bson:"tag"`
	Value string `json:"value" yaml:"value" bson:"value"`
}
