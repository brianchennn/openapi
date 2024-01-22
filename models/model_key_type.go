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


type KeyType string

// List of KeyType
const (
        KeyType_UNIQUE_KEY KeyType = "UNIQUE_KEY"
        KeyType_SEARCH_KEY KeyType = "SEARCH_KEY"
        KeyType_COUNT_KEY KeyType = "COUNT_KEY"
        KeyType_SEARCH_AND_COUNT_KEY KeyType = "SEARCH_AND_COUNT_KEY"
        KeyType_OTHER_TAG KeyType = "OTHER_TAG"
)

