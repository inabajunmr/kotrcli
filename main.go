package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Dakoku(id string, userToken string, token string) {

	client := &http.Client{}
	t := time.Now()
	unique_timestamp := t.Format("20060102150405")
	param := fmt.Sprintf("id=%v&highAccuracyFlg=false&credential_code=40&user_token=%v&comment=&unique_timestamp=%v&version=1.2.7&token=%v", id, userToken, unique_timestamp, token)
	buffer := bytes.NewBufferString(param)

	request, err := http.NewRequest("POST", "https://s2.kingtime.jp/gateway/bprgateway", buffer)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	request.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	request.Header.Add("X-Requested-With", "XMLHttpRequest")
	request.Header.Add("Connection", "keep-alive")

	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(string(body))
}

func main() {

	// syukkin := "qmXXCxw9WEWN3X/YrkMWuQ=="
	taikin := "j8ekmJaw6W3M4w3i6hlSIQ=="
	// userToken := ""
	// token := ""

	Dakoku(taikin, userToken, token)
}

