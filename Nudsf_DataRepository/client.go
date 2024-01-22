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

package Nudsf_DataRepository

import (
    "crypto/tls"
    "golang.org/x/net/http2"
    "net/http"
)

// APIClient manages communication with the Nudsf_DataRepository API v1.1.0
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
    cfg    *Configuration
    common service // Reuse a single struct instead of allocating one for each service on the heap.

    // API Services
    	BlockCRUDApi *BlockCRUDApiService
    	MetaSchemaCRUDApi *MetaSchemaCRUDApiService
    	NotificationSubscriptionCRUDApi *NotificationSubscriptionCRUDApiService
    	NotificationSubscriptionsCRUDApi *NotificationSubscriptionsCRUDApiService
    	RecordCRUDApi *RecordCRUDApiService
}

type service struct {
    client *APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) *APIClient {
    if cfg.httpClient == nil {
        cfg.httpClient = http.DefaultClient
        cfg.httpClient.Transport = &http2.Transport{
            TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
        }
    }

    c := &APIClient{}
    c.cfg = cfg
    c.common.client = c

    // API Services
    c.BlockCRUDApi = (*BlockCRUDApiService)(&c.common)
    c.MetaSchemaCRUDApi = (*MetaSchemaCRUDApiService)(&c.common)
    c.NotificationSubscriptionCRUDApi = (*NotificationSubscriptionCRUDApiService)(&c.common)
    c.NotificationSubscriptionsCRUDApi = (*NotificationSubscriptionsCRUDApiService)(&c.common)
    c.RecordCRUDApi = (*RecordCRUDApiService)(&c.common)

    return c
}