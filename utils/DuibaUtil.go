package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

func BuildUrlWithSign(duibaurl string, appKey string, appSecret string, params map[string]string) string {

	if _, ok := params["timestamp"]; !ok {
		params["timestamp"] = strconv.FormatInt((time.Now().UnixNano() / 1e6), 10)
	}

	params["appKey"] = appKey
	params["appSecret"] = appSecret
	signResult := Sign(params)
	params["sign"] = signResult
	delete(params, "appSecret")
	return assembleUrl(duibaurl, params)
}

func Sign(params map[string]string) string {
	keys := make([]string, len(params))
	i := 0
	for key, _ := range params {
		keys[i] = key
		i++
	}
	sort.Strings(keys)
	var buf bytes.Buffer
	for _, key := range keys {
		buf.WriteString(params[key])
	}
	result := hex.EncodeToString(EncryptMD5(buf.Bytes()))
	return result
}

func EncryptMD5(data []byte) []byte {
	h := md5.New()
	h.Write(data)
	result := h.Sum(nil)
	return result
}

func assembleUrl(duibaurl string, params map[string]string) string {
	if !strings.HasSuffix(duibaurl, "?") {
		duibaurl += "?"
	}

	for key, value := range params {
		if len(value) == 0 {
			duibaurl += key + "=" + value + "&"
		} else {
			duibaurl += key + "=" + url.QueryEscape(value) + "&"
		}
	}
	return duibaurl
}
