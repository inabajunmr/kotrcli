package kotrcli

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
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
	uniqueTimestamp := t.Format("20060102150405")
	param := url.Values{}
	param.Add("id", getTypeValue(val))
	param.Add("highAccuracyFlg", "false")
	param.Add("credential_code", "40")
	param.Add("user_token", userToken)
	param.Add("unique_timestamp", uniqueTimestamp)
	param.Add("comment", "")
	param.Add("version", "1.2.7")
	param.Add("token", token)
	request, err := http.NewRequest("POST", "https://s2.kingtime.jp/gateway/bprgateway", bytes.NewBufferString(param.Encode()))
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
