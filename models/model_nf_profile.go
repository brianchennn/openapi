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
import (
	"time"
)



// Information of an NF Instance registered in the NRF
type NfProfile struct {
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	NfInstanceId string `json:"nfInstanceId" yaml:"nfInstanceId" bson:"nfInstanceId"`
	NfInstanceName string `json:"nfInstanceName,omitempty" yaml:"nfInstanceName" bson:"nfInstanceName"`
	NfType NfType `json:"nfType" yaml:"nfType" bson:"nfType"`
	NfStatus NfStatus `json:"nfStatus" yaml:"nfStatus" bson:"nfStatus"`
	CollocatedNfInstances []*CollocatedNfInstance `json:"collocatedNfInstances,omitempty" yaml:"collocatedNfInstances" bson:"collocatedNfInstances"`
	HeartBeatTimer int32 `json:"heartBeatTimer,omitempty" yaml:"heartBeatTimer" bson:"heartBeatTimer"`
	PlmnList []*PlmnId `json:"plmnList,omitempty" yaml:"plmnList" bson:"plmnList"`
	SnpnList []*PlmnIdNid `json:"snpnList,omitempty" yaml:"snpnList" bson:"snpnList"`
	SNssais []*ExtSnssai `json:"sNssais,omitempty" yaml:"sNssais" bson:"sNssais"`
	PerPlmnSnssaiList []*PlmnSnssai `json:"perPlmnSnssaiList,omitempty" yaml:"perPlmnSnssaiList" bson:"perPlmnSnssaiList"`
	NsiList []*string `json:"nsiList,omitempty" yaml:"nsiList" bson:"nsiList"`
	// Fully Qualified Domain Name
	Fqdn string `json:"fqdn,omitempty" yaml:"fqdn" bson:"fqdn"`
	// Fully Qualified Domain Name
	InterPlmnFqdn string `json:"interPlmnFqdn,omitempty" yaml:"interPlmnFqdn" bson:"interPlmnFqdn"`
	Ipv4Addresses []*string `json:"ipv4Addresses,omitempty" yaml:"ipv4Addresses" bson:"ipv4Addresses"`
	Ipv6Addresses []*string `json:"ipv6Addresses,omitempty" yaml:"ipv6Addresses" bson:"ipv6Addresses"`
	AllowedPlmns []*PlmnId `json:"allowedPlmns,omitempty" yaml:"allowedPlmns" bson:"allowedPlmns"`
	AllowedSnpns []*PlmnIdNid `json:"allowedSnpns,omitempty" yaml:"allowedSnpns" bson:"allowedSnpns"`
	AllowedNfTypes []*NfType `json:"allowedNfTypes,omitempty" yaml:"allowedNfTypes" bson:"allowedNfTypes"`
	AllowedNfDomains []*string `json:"allowedNfDomains,omitempty" yaml:"allowedNfDomains" bson:"allowedNfDomains"`
	AllowedNssais []*ExtSnssai `json:"allowedNssais,omitempty" yaml:"allowedNssais" bson:"allowedNssais"`
	Priority int32 `json:"priority,omitempty" yaml:"priority" bson:"priority"`
	Capacity int32 `json:"capacity,omitempty" yaml:"capacity" bson:"capacity"`
	Load int32 `json:"load,omitempty" yaml:"load" bson:"load"`
	// string with format 'date-time' as defined in OpenAPI.
	LoadTimeStamp *time.Time `json:"loadTimeStamp,omitempty" yaml:"loadTimeStamp" bson:"loadTimeStamp"`
	Locality string `json:"locality,omitempty" yaml:"locality" bson:"locality"`
	UdrInfo *UdrInfo `json:"udrInfo,omitempty" yaml:"udrInfo" bson:"udrInfo"`
	// A map (list of key-value pairs) where a (unique) valid JSON string serves as key of UdrInfo 
	UdrInfoList map[string]*UdrInfo `json:"udrInfoList,omitempty" yaml:"udrInfoList" bson:"udrInfoList"`
	UdmInfo *UdmInfo `json:"udmInfo,omitempty" yaml:"udmInfo" bson:"udmInfo"`
	// A map (list of key-value pairs) where a (unique) valid JSON string serves as key of UdmInfo 
	UdmInfoList map[string]*UdmInfo `json:"udmInfoList,omitempty" yaml:"udmInfoList" bson:"udmInfoList"`
	AusfInfo *AusfInfo `json:"ausfInfo,omitempty" yaml:"ausfInfo" bson:"ausfInfo"`
	// A map (list of key-value pairs) where a (unique) valid JSON string serves as key of AusfInfo 
	AusfInfoList map[string]*AusfInfo `json:"ausfInfoList,omitempty" yaml:"ausfInfoList" bson:"ausfInfoList"`
	AmfInfo *AmfInfo `json:"amfInfo,omitempty" yaml:"amfInfo" bson:"amfInfo"`
	// A map (list of key-value pairs) where a (unique) valid JSON string serves as key of AmfInfo 
	AmfInfoList map[string]*AmfInfo `json:"amfInfoList,omitempty" yaml:"amfInfoList" bson:"amfInfoList"`
	SmfInfo *SmfInfo `json:"smfInfo,omitempty" yaml:"smfInfo" bson:"smfInfo"`
	// A map (list of key-value pairs) where a (unique) valid JSON string serves as key of SmfInfo 
	SmfInfoList map[string]*SmfInfo `json:"smfInfoList,omitempty" yaml:"smfInfoList" bson:"smfInfoList"`
	UpfInfo *UpfInfo `json:"upfInfo,omitempty" yaml:"upfInfo" bson:"upfInfo"`
	// A map (list of key-value pairs) where a (unique) valid JSON string serves as key of UpfInfo 
	UpfInfoList map[string]*UpfInfo `json:"upfInfoList,omitempty" yaml:"upfInfoList" bson:"upfInfoList"`
	PcfInfo *PcfInfo `json:"pcfInfo,omitempty" yaml:"pcfInfo" bson:"pcfInfo"`
	// A map (list of key-value pairs) where a (unique) valid JSON string serves as key of PcfInfo 
	PcfInfoList map[string]*PcfInfo `json:"pcfInfoList,omitempty" yaml:"pcfInfoList" bson:"pcfInfoList"`
	BsfInfo *BsfInfo `json:"bsfInfo,omitempty" yaml:"bsfInfo" bson:"bsfInfo"`
	// A map (list of key-value pairs) where a (unique) valid JSON string serves as key of BsfInfo 
	BsfInfoList map[string]*BsfInfo `json:"bsfInfoList,omitempty" yaml:"bsfInfoList" bson:"bsfInfoList"`
	ChfInfo *ChfInfo `json:"chfInfo,omitempty" yaml:"chfInfo" bson:"chfInfo"`
	// A map (list of key-value pairs) where a (unique) valid JSON string serves as key of ChfInfo 
	ChfInfoList map[string]*ChfInfo `json:"chfInfoList,omitempty" yaml:"chfInfoList" bson:"chfInfoList"`
	NefInfo *NefInfo `json:"nefInfo,omitempty" yaml:"nefInfo" bson:"nefInfo"`
	NrfInfo *NrfInfo `json:"nrfInfo,omitempty" yaml:"nrfInfo" bson:"nrfInfo"`
	UdsfInfo *UdsfInfo `json:"udsfInfo,omitempty" yaml:"udsfInfo" bson:"udsfInfo"`
	// A map (list of key-value pairs) where a (unique) valid JSON string serves as key of UdsfInfo 
	UdsfInfoList map[string]*UdsfInfo `json:"udsfInfoList,omitempty" yaml:"udsfInfoList" bson:"udsfInfoList"`
	NwdafInfo *NwdafInfo `json:"nwdafInfo,omitempty" yaml:"nwdafInfo" bson:"nwdafInfo"`
	// A map (list of key-value pairs) where a (unique) valid JSON string serves as key of NwdafInfo 
	NwdafInfoList map[string]*NwdafInfo `json:"nwdafInfoList,omitempty" yaml:"nwdafInfoList" bson:"nwdafInfoList"`
	// A map (list of key-value pairs) where a (unique) valid JSON string serves as key of PcscfInfo 
	PcscfInfoList map[string]*PcscfInfo `json:"pcscfInfoList,omitempty" yaml:"pcscfInfoList" bson:"pcscfInfoList"`
	// A map (list of key-value pairs) where a (unique) valid JSON string serves as key of HssInfo 
	HssInfoList map[string]*HssInfo `json:"hssInfoList,omitempty" yaml:"hssInfoList" bson:"hssInfoList"`
	CustomInfo map[string]interface{} `json:"customInfo,omitempty" yaml:"customInfo" bson:"customInfo"`
	// string with format 'date-time' as defined in OpenAPI.
	RecoveryTime *time.Time `json:"recoveryTime,omitempty" yaml:"recoveryTime" bson:"recoveryTime"`
	NfServicePersistence bool `json:"nfServicePersistence,omitempty" yaml:"nfServicePersistence" bson:"nfServicePersistence"`
	NfServices []*NfService `json:"nfServices,omitempty" yaml:"nfServices" bson:"nfServices"`
	// A map (list of key-value pairs) where serviceInstanceId serves as key of NFService 
	NfServiceList map[string]*NfService `json:"nfServiceList,omitempty" yaml:"nfServiceList" bson:"nfServiceList"`
	NfProfileChangesSupportInd bool `json:"nfProfileChangesSupportInd,omitempty" yaml:"nfProfileChangesSupportInd" bson:"nfProfileChangesSupportInd"`
	NfProfileChangesInd bool `json:"nfProfileChangesInd,omitempty" yaml:"nfProfileChangesInd" bson:"nfProfileChangesInd"`
	DefaultNotificationSubscriptions []*DefaultNotificationSubscription `json:"defaultNotificationSubscriptions,omitempty" yaml:"defaultNotificationSubscriptions" bson:"defaultNotificationSubscriptions"`
	LmfInfo *LmfInfo `json:"lmfInfo,omitempty" yaml:"lmfInfo" bson:"lmfInfo"`
	GmlcInfo *GmlcInfo `json:"gmlcInfo,omitempty" yaml:"gmlcInfo" bson:"gmlcInfo"`
	NfSetIdList []*string `json:"nfSetIdList,omitempty" yaml:"nfSetIdList" bson:"nfSetIdList"`
	ServingScope []*string `json:"servingScope,omitempty" yaml:"servingScope" bson:"servingScope"`
	LcHSupportInd bool `json:"lcHSupportInd,omitempty" yaml:"lcHSupportInd" bson:"lcHSupportInd"`
	OlcHSupportInd bool `json:"olcHSupportInd,omitempty" yaml:"olcHSupportInd" bson:"olcHSupportInd"`
	// A map (list of key-value pairs) where NfSetId serves as key of DateTime
	NfSetRecoveryTimeList map[string]*time.Time `json:"nfSetRecoveryTimeList,omitempty" yaml:"nfSetRecoveryTimeList" bson:"nfSetRecoveryTimeList"`
	// A map (list of key-value pairs) where NfServiceSetId serves as key of DateTime 
	ServiceSetRecoveryTimeList map[string]*time.Time `json:"serviceSetRecoveryTimeList,omitempty" yaml:"serviceSetRecoveryTimeList" bson:"serviceSetRecoveryTimeList"`
	ScpDomains []*string `json:"scpDomains,omitempty" yaml:"scpDomains" bson:"scpDomains"`
	ScpInfo *ScpInfo `json:"scpInfo,omitempty" yaml:"scpInfo" bson:"scpInfo"`
	SeppInfo *SeppInfo `json:"seppInfo,omitempty" yaml:"seppInfo" bson:"seppInfo"`
	// Vendor ID of the NF Service instance (Private Enterprise Number assigned by IANA)
	VendorId string `json:"vendorId,omitempty" yaml:"vendorId" bson:"vendorId"`
	// The key of the map is the IANA-assigned SMI Network Management Private Enterprise Codes 
	SupportedVendorSpecificFeatures map[string]*[]*VendorSpecificFeature `json:"supportedVendorSpecificFeatures,omitempty" yaml:"supportedVendorSpecificFeatures" bson:"supportedVendorSpecificFeatures"`
	// A map (list of key-value pairs) where a (unique) valid JSON string serves as key of AanfInfo 
	AanfInfoList map[string]*AanfInfo `json:"aanfInfoList,omitempty" yaml:"aanfInfoList" bson:"aanfInfoList"`
	Var5gDdnmfInfo *Model5GDdnmfInfo `json:"5gDdnmfInfo,omitempty" yaml:"5gDdnmfInfo" bson:"5gDdnmfInfo"`
	MfafInfo *MfafInfo `json:"mfafInfo,omitempty" yaml:"mfafInfo" bson:"mfafInfo"`
	// A map (list of key-value pairs) where a (unique) valid JSON string serves as key of EasdfInfo 
	EasdfInfoList map[string]*EasdfInfo `json:"easdfInfoList,omitempty" yaml:"easdfInfoList" bson:"easdfInfoList"`
	DccfInfo *DccfInfo `json:"dccfInfo,omitempty" yaml:"dccfInfo" bson:"dccfInfo"`
	// A map (list of key-value pairs) where a (unique) valid JSON string serves as key of NsacfInfo 
	NsacfInfoList map[string]*NsacfInfo `json:"nsacfInfoList,omitempty" yaml:"nsacfInfoList" bson:"nsacfInfoList"`
	// A map (list of key-value pairs) where a (unique) valid JSON string serves as key of MbSmfInfo 
	MbSmfInfoList map[string]*MbSmfInfo `json:"mbSmfInfoList,omitempty" yaml:"mbSmfInfoList" bson:"mbSmfInfoList"`
	// A map (list of key-value pairs) where a (unique) valid JSON string serves as key of TsctsfInfo 
	TsctsfInfoList map[string]*TsctsfInfo `json:"tsctsfInfoList,omitempty" yaml:"tsctsfInfoList" bson:"tsctsfInfoList"`
	// A map (list of key-value pairs) where a (unique) valid JSON string serves as key of MbUpfInfo 
	MbUpfInfoList map[string]*MbUpfInfo `json:"mbUpfInfoList,omitempty" yaml:"mbUpfInfoList" bson:"mbUpfInfoList"`
	TrustAfInfo *TrustAfInfo `json:"trustAfInfo,omitempty" yaml:"trustAfInfo" bson:"trustAfInfo"`
	NssaafInfo *NssaafInfo `json:"nssaafInfo,omitempty" yaml:"nssaafInfo" bson:"nssaafInfo"`
	HniList []*string `json:"hniList,omitempty" yaml:"hniList" bson:"hniList"`
	IwmscInfo *IwmscInfo `json:"iwmscInfo,omitempty" yaml:"iwmscInfo" bson:"iwmscInfo"`
	MnpfInfo *MnpfInfo `json:"mnpfInfo,omitempty" yaml:"mnpfInfo" bson:"mnpfInfo"`
}
