package utils

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/astaxie/beego"
)

// get方式获取请
func HttpDoGetStrParams(requestUrl string, requestParams string) (string, error) {
	client := &http.Client{}

	if !strings.HasSuffix(requestUrl, "?") {
		requestUrl += "?"
	}
	if requestParams != "" {
		requestUrl += requestParams
	}
	request, err := http.NewRequest(http.MethodGet, requestUrl, nil)

	if err != nil {
		return "", err
	}

	response, err := client.Do(request)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return "", err
	}
	result, err := url.QueryUnescape(string(data))
	if err != nil {
		return "", err
	}
	return result, nil
}

// get方式获取请
func HttpDoGetMapParams(requestUrl string, requestParams map[string]string) (string, error) {
	if !strings.HasSuffix(requestUrl, "?") {
		requestUrl += "?"
	}
	params := ""
	if requestParams != nil {
		for key, value := range requestParams {
			if len(value) == 0 {
				params += key + "=" + value + "&"
			} else {
				params += key + "=" + url.QueryEscape(value) + "&"
			}
		}
	}
	return HttpDoGetStrParams(requestUrl, params)
}

func HttpDoPost(requestUrl string, params map[string]string) {
	beego.BeeLogger.Info("httpDoPost")
}
