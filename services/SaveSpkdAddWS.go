package services

import (
	"bytes"
	"encoding/xml"
	"errors"
	//"flag"
	"fmt"
	"io"
	//	"io/ioutil"
	//"crypto/tls"
	//"crypto/x509"
	"encoding/base64"
	//"log"
	//"net"
	"net/http"
	"os"
	"time"
	//"strings"
	//	"unicode/utf8"
)

//const entrustCert = `-----BEGIN CERTIFICATE-----
//...
//-----END CERTIFICATE-----`

type Encryption struct {
	Username string
	Password string
}

func (encrypt Encryption) GetEncrypted() string {
	authen := []byte(encrypt.Username + ":" + encrypt.Password)
	encPass := &bytes.Buffer{}
	encrypted := base64.NewEncoder(base64.StdEncoding, encPass)
	encrypted.Write(authen)
	encrypted.Close()

	return string(encPass.Bytes())
}

func (saveSpkdAddWS SaveSpkdAddWS) CreateSoapEnvelope(encrypt Encryption) *SoapEnvelope {
	const layout = "2006-01-02T15:04:05.999Z"
	t := time.Now()

	retval := &SoapEnvelope{}
	retval.SoapEnv = "http://schemas.xmlsoap.org/soap/envelope/"
	retval.Urn = "urn:pack.INSaveSpkdAddWS_typedef.salt11"
	retval.Body.SaveSpkdAddWS = saveSpkdAddWS // SaveSpkdAddWS{Usercode: "LLTHUNYADAP", BLPDINDC: "PCN", CSSPKDPCNCUSTNUMB: "536672462", CSSPKDPCNSUBRNUMB: "66900010040", CSSPKDPCNPACKCODE: "31001501", RDTELPTELPTYPE: "TEL", SAVEFLAG: "1"}
	passwordUsernameToken := PasswordUsernameToken{Password: encrypt.Password, Type: "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordText"}
	nonceUsernameToken := NonceUsernameToken{Nonce: encrypt.GetEncrypted(), EncodingType: "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-soap-message-security-1.0#Base64Binary"}
	userToken := UsernameToken{Id: "UsernameToken-20", Wsu: "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd", Username: encrypt.Username, Password: passwordUsernameToken, Nonce: nonceUsernameToken, Created: t.UTC().Format(layout)}
	retval.Header.Security = Security{MustUnderstand: "1", Wsse: "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd", UsernameToken: userToken}
	return retval
}

func PrintRequest(requestEnvelope *SoapEnvelope) {
	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("  ", "    ")
	if err := enc.Encode(requestEnvelope); err != nil {
		fmt.Printf("error: %v\n", err)
	}
}

type RequestInformation struct {
	BufferOfRequest io.Reader
	Endpoint        string
	SoapAction      string
}

func (requestInfo RequestInformation) GetResponse() *http.Response {
	var resp *http.Response
	var err error
	if resp, err = http.Post(requestInfo.Endpoint, "application/soap+xml; charset=UTF-8; action="+requestInfo.SoapAction, requestInfo.BufferOfRequest); err != nil {
		println(err.Error())
		return nil
	}
	return resp
}

func CallTux(saveSpkdAddWS SaveSpkdAddWS) (*SaveSpkdAddWSResponse, *FaultResponse, error, string) {

	buffer := new(bytes.Buffer) //&bytes.Buffer{}

	var encrypt = Encryption{Username: "LLCALLCENTER", Password: "ae1234"}
	//var saveSpkdAddWS = SaveSpkdAddWS{USER_CODE: "LLTHUNYADAP", BLPD_INDC: "PCN", CS_SPKD_PCN__CUST_NUMB: "536672462", CS_SPKD_PCN__SUBR_NUMB: "66900010040", CS_SPKD_PCN__PACK_CODE: "31001501", RD_TELP__TELP_TYPE: "TEL", SAVE_FLAG: "1"}
	requestEnvelope := saveSpkdAddWS.CreateSoapEnvelope(encrypt)

	encoder := xml.NewEncoder(buffer)
	err := encoder.Encode(requestEnvelope)
	if err != nil {
		println("Error encoding document:", err.Error())
		//return
	}

	//PrintRequest(requestEnvelope)

	var resp *http.Response

	//http://DTH2732042T01:8088/mockINSaveSpkdAddWS_Binding
	//http://athena13:9582/SaveSpkdAddWS
	var reqInfo = RequestInformation{BufferOfRequest: buffer, Endpoint: "http://10.89.75.44:9582/SaveSpkdAddWS", SoapAction: "SaveSpkdAddWS"}
	resp = reqInfo.GetResponse()

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		println("Soap Error:", resp.Status)
	}

	bodyElement, faultElement, err := DecodeResponseBody(resp.Body)

	var state string
	if err != nil {
		//if strings.ContainsAny(err.Error(), "encoding") {
		//	println("Error decoding: ", err.Error())
		//} else {
		//	println("Error: ", err.Error())
		//}
		state = "error"
		//return
	}

	if faultElement != nil {
		//println("This request is fault but we don't know why. We promise you can know this in the future. See you soon.")
		state = "fault"
	}

	if bodyElement != nil {
		state = "body"
		//	println("Empty output")
		//} else {
		//	temp := bodyElement.CS_PKPL_PCN__PACK_DESC
		//	fmt.Println("\r\nResult = ", temp)
	}

	return bodyElement, faultElement, err, state
}

func DecodeResponseBody(body io.Reader) (*SaveSpkdAddWSResponse, *FaultResponse, error) {
	decoder := xml.NewDecoder(body)

	nextElementIsBody := false
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, nil, err
		}
		switch startElement := token.(type) {
		case xml.StartElement:
			if nextElementIsBody {
				responseBody := SaveSpkdAddWSResponse{}
				err = decoder.DecodeElement(&responseBody, &startElement)
				if err != nil {
					responseFault := FaultResponse{}

					//decoder.CharsetReader("UTF-8", body)
					//err = decoder.DecodeElement(&responseFault, &startElement)
					if err != nil {
						return nil, nil, err
					} else {
						return nil, &responseFault, nil
					}
					return nil, nil, err
				}
				return &responseBody, nil, nil
			}
			if startElement.Name.Space == "http://schemas.xmlsoap.org/soap/envelope/" && startElement.Name.Local == "Body" {
				nextElementIsBody = true
			}
		}
	}

	return nil, nil, errors.New("Did not find SOAP body element")
}
