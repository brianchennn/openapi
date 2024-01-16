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



// Description of a record notification
type NotificationDescription struct {
	// String providing an URI formatted according to RFC 3986.
	RecordRef string `json:"recordRef" yaml:"recordRef" bson:"recordRef"`
	OperationType RecordOperation `json:"operationType" yaml:"operationType" bson:"operationType"`
	SubscriptionId string `json:"subscriptionId,omitempty" yaml:"subscriptionId" bson:"subscriptionId"`
}
