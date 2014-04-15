package services

import (
	"encoding/xml"
	"math/big"
)

var _ = big.MaxBase // Avoid potential unused-import error

type SoapEnvelope struct {
	XMLName xml.Name   `xml:"soapenv:Envelope"`
	SoapEnv string     `xml:"xmlns:soapenv,attr"`
	Urn     string     `xml:"xmlns:urn,attr"`
	Header  SoapHeader `xml:"soapenv:Header"`
	Body    SoapBody   `xml:"soapenv:Body"`
}

type SoapHeader struct {
	Security Security `xml:"wsse:Security"`
}

type SoapBody struct {
	SaveSpkdAddWS interface{}
}

type SaveSpkdAddWS struct {
	XMLName                xml.Name `xml:"SaveSpkdAddWS"`
	USER_CODE              string   `xml:"urn:inbuf>USER_CODE"`
	BLPD_INDC              string   `xml:"urn:inbuf>BLPD_INDC"`
	CS_SPKD_PCN__CUST_NUMB string   `xml:"urn:inbuf>CS_SPKD_PCN__CUST_NUMB"`
	CS_SPKD_PCN__SUBR_NUMB string   `xml:"urn:inbuf>CS_SPKD_PCN__SUBR_NUMB"`
	CS_SPKD_PCN__PACK_CODE string   `xml:"urn:inbuf>CS_SPKD_PCN__PACK_CODE"`
	//CS_SUBR_PCN__SUBR_TYPE string   `xml:"urn:inbuf>CS_SUBR_PCN__SUBR_TYPE"`
	RD_TELP__TELP_TYPE string `xml:"urn:inbuf>RD_TELP__TELP_TYPE"`
	SAVE_FLAG          string `xml:"urn:inbuf>SAVE_FLAG"`
}

type Security struct {
	XMLName        xml.Name      `xml:"wsse:Security"`
	MustUnderstand string        `xml:"soapenv:mustUnderstand,attr"`
	Wsse           string        `xml:"xmlns:wsse,attr"`
	UsernameToken  UsernameToken `xml:"wsse:UsernameToken"`
}

type UsernameToken struct {
	XMLName  xml.Name              `xml:"wsse:UsernameToken"`
	Id       string                `xml:"wsu:Id,attr"`
	Wsu      string                `xml:"xmlns:wsu,attr"`
	Username string                `xml:"wsse:Username"`
	Password PasswordUsernameToken `xml:"wsse:Password"`
	Nonce    NonceUsernameToken    `xml:"wsse:Nonce"`
	Created  string                `xml:"wsse:Created"`
}

type PasswordUsernameToken struct {
	XMLName  xml.Name `xml:""`
	Password string   `xml:",chardata"`
	Type     string   `xml:",attr"`
}
type NonceUsernameToken struct {
	XMLName      xml.Name `xml:""`
	Nonce        string   `xml:",chardata"`
	EncodingType string   `xml:",attr"`
}

type SoapEnvelopeResponse struct {
	XMLName xml.Name         `xml:"SOAP-ENV:Envelope"`
	SoapEnv string           `xml:"xmlns:SOAP-ENV,attr"`
	Tuxedo  string           `xml:"xmlns:tuxedo,attr"`
	Header  string           `xml:"SOAP-ENV:Header"`
	Body    SoapBodyResponse `xml:"SOAP-ENV:Body"`
}

type SoapBodyResponse struct {
	SaveSpkdAddWSResponse interface{}
}

type SaveSpkdAddWSResponse struct {
	XMLName                      xml.Name `xml:"SaveSpkdAddWSResponse"`
	CS_SPKD_PCN__PACK_CODE       string   `xml:"outbuf>CS_SPKD_PCN__PACK_CODE"`
	CS_PKPL_PCN__PACK_DESC       string   `xml:"outbuf>CS_PKPL_PCN__PACK_DESC"`
	CS_PACK_TYPE__PACK_TYPE_DESC string   `xml:"outbuf>CS_PACK_TYPE__PACK_TYPE_DESC"`
	CS_SPKD_PCN__PACK_STRT_DTTM  string   `xml:"outbuf>CS_SPKD_PCN__PACK_STRT_DTTM"`
	CS_SPKD_PCN__PACK_END_DTTM   string   `xml:"outbuf>CS_SPKD_PCN__PACK_END_DTTM"`
	CS_SPKD_PCN__DISC_CODE       string   `xml:"outbuf>CS_SPKD_PCN__DISC_CODE"`
	TBL_OCCR                     string   `xml:"outbuf>TBL_OCCR"`
	MESSAGE_TEXT_ENG             string   `xml:"errbuf>MESSAGE_TEXT_ENG"`
	ESSAGE_TEXT_THA              string   `xml:"errbuf>ESSAGE_TEXT_THA"`
	MESSAGE_SQLCODE              string   `xml:"errbuf>MESSAGE_SQLCODE"`
	MESSAGE_NATURE               string   `xml:"errbuf>MESSAGE_NATURE"`
}

type FaultResponse struct {
	XMLName     xml.Name     `xml:"Fault"`
	Faultcode   string       `xml:"faultcode"`
	Faultstring string       `xml:"faultstring"`
	Detail      DetailStruct `xml:"detail"`
}

type DetailStruct struct {
	XMLName            xml.Name                 `xml:"detail"`
	SaveSpkdAddWSFault SaveSpkdAddWSFaultStruct `xml:"tuxedo:SaveSpkdAddWSFault"`
}

type SaveSpkdAddWSFaultStruct struct {
	XMLName xml.Name     `xml:"tuxedo:SaveSpkdAddWSFault"`
	Errbuf  ErrbufStruct `xml:"tuxedo:errbuf"`
}

type ErrbufStruct struct {
	XMLName          xml.Name `xml:"tuxedo:errbuf"`
	MESSAGE_TEXT_ENG string   `xml:"tuxedo:MESSAGE_TEXT_ENG"`
	MESSAGE_TEXT_THA string   `xml:"tuxedo:MESSAGE_TEXT_THA"`
	MESSAGE_SQLCODE  string   `xml:"tuxedo:MESSAGE_SQLCODE"`
	MESSAGE_NATURE   string   `xml:"tuxedo:MESSAGE_NATURE"`
}
