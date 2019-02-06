package kotrcli

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Type int

const (
	SYUKKIN Type = iota
	TAIKIN
)

func Dakoku(val Type, userToken string, token string) (string, error) {

	client := &http.Client{}
	t := time.Now()
	unique_timestamp := t.Format("20060102150405")
	param := fmt.Sprintf("id=%v&highAccuracyFlg=false&credential_code=40&user_token=%v&comment=&unique_timestamp=%v&version=1.2.7&token=%v", getTypeValue(val), userToken, unique_timestamp, token)
	buffer := bytes.NewBufferString(param)

	request, err := http.NewRequest("POST", "https://s2.kingtime.jp/gateway/bprgateway", buffer)
	if err != nil {
		return "", err
	}

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	request.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	request.Header.Add("X-Requested-With", "XMLHttpRequest")
	request.Header.Add("Connection", "keep-alive")

	resp, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func getTypeValue(val Type) string {
	switch val {
	case SYUKKIN:
		return "qmXXCxw9WEWN3X/YrkMWuQ=="
	case TAIKIN:
		return "j8ekmJaw6W3M4w3i6hlSIQ=="
	}
	return ""
}
