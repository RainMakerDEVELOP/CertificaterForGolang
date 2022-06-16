package certparser

import (
	"crypto/x509"
	"crypto/x509/pkix"
	"fmt"
	"io/ioutil"
)

type CertInfo struct {
	Id       int
	CertData x509.Certificate
}

func Test1() {
	fmt.Println("Test1")
}

func NewCertInfo() *CertInfo {
	var ci = &CertInfo{}
	return ci.NewCertInfo()
}

// Test2 CertInfo 객체 테스트
func (ci *CertInfo) Test2() {
	fmt.Println(ci.Id)
}

// NewCertInfo CertInfo 객체 생성 함수
func (ci *CertInfo) NewCertInfo() *CertInfo {
	return &CertInfo{}
}

// ParseCertData DER 데이터 파싱 후 데이터 리턴 처리 함수
func (ci *CertInfo) ParseCertData(derData []byte) (*x509.Certificate, error) {
	pCertData, err := x509.ParseCertificate(derData)

	if err != nil {
		return pCertData, err
	}

	return pCertData, nil
}

func (ci *CertInfo) ParseCertExtData(ct *x509.Certificate) []pkix.Extension {
	var extensions []pkix.Extension
	for _, ext := range ct.Extensions {
		// if strings.Contains(ext.Id.String(), customOIDPrefix) {
		extensions = append(extensions, ext)
		// }
	}

	return extensions
}

func (ci *CertInfo) ParseCertDataThread(derFileName string) (*x509.Certificate, error) {
	data, err := ioutil.ReadFile(derFileName)
	if err != nil {
		fmt.Println("파일을 찾을 수 없습니다.", derFileName)
		return nil, err
	}

	// DER 파일 파싱
	pCertData, err := ci.ParseCertData(data)

	if err != nil {
		return nil, err
	}

	// OID값 추출
	for nLen := 0; nLen < len(pCertData.PolicyIdentifiers); nLen++ {
		strOID := pCertData.PolicyIdentifiers[nLen].String()
		fmt.Println(strOID)
	}

	// extensions 정보 추출
	extensions := ci.ParseCertExtData(pCertData)

	// fmt.Println(extensions)
	// fmt.Println(extensions[0].Value[0])
	// extData := cryptobyte.String(data)

	fmt.Sprintf("%s", extensions[0].Value)

	// certificates, err := x509.ParseCertificates(extensions[0].Value)

	// if err != nil {
	// 	return nil, err
	// }

	// fmt.Println(certificates)

	return pCertData, nil
}
