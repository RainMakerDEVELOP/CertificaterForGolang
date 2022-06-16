package main

import (
	"fmt"
	"goproject/certificater/certparser"
	"log"
	"net/http"
	"os"
)

// Global 변수
var g_clientInfo map[string]any
var g_strDataType string
var g_byteData []byte

// Argument
// 1 : 클라이언트 정보 (JSON 구조)
// 2 : 형식 (FILE, DER ...)
// 3 : 데이터 ([]byte ...)
func main() {
	if len(os.Args) < 4 {
		log.Println("잘못된 호출. 인자값 부족")
		return
	}

	// 실행 인자값 글로벌 변수에 할당
	setValidFromArgs(os.Args)

	// http 서버 구동 및 핸들 등록
	http.ListenAndServe(":443", initHttpHandleFunc())

	pCp := certparser.NewCertInfo()
	// C:\Program Files\NPKI\KICA\USER\cn=밝은미소의료소비자생활협동조합참미소치과,ou=KNET,ou=조달연구원,ou=등록기관,o=SignKorea,c=KR
	// strDerFilePath := "C:\\Program Files\\NPKI\\KICA\\USER\\cn=이수본치과의원,ou=건강보험,ou=mohw ra센터,ou=등록기관,ou=licensedca,o=kica,c=kr\\signcert.der"
	strDerFilePath := "C:\\Program Files\\NPKI\\KICA\\USER\\cn=밝은미소의료소비자생활협동조합참미소치과,ou=KNET,ou=조달연구원,ou=등록기관,o=SignKorea,c=KR\\signcert.der"
	pCertData, err := pCp.ParseCertDataThread(strDerFilePath)

	if err != nil {
		return
	}

	fmt.Println(pCertData.NotAfter)

	fmt.Println("종료")
}

// setValidFromArgs
// 실행 인자값 저장
func setValidFromArgs(arg []string) {
	setClientInfoFromValid(arg[1])
	g_strDataType = arg[2]
	g_byteData = []byte(arg[3])

	fmt.Println("'", g_strDataType, "', '", g_byteData, "'")
}

func setClientInfoFromValid(val string) {
	g_clientInfo = make(map[string]any)

}

// initHttpHandleFunc
// Http 처리용 핸들 초기화
func initHttpHandleFunc() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome! Certificate Parser")
	})

	return mux
}
