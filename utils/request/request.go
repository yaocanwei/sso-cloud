/**
  @author: cheney
  @date: 2021/5/9
  @note:
 **/
package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func ApiGet(host, path, method string, req urlRequestParse, respObj interface{}) error {
	url := composeApiUrl(host, path, req)
	urlStr := url.String()

	httpreq, err := http.NewRequest(method, urlStr, nil)
	//TODO: composeWithAuthorization

	client := &http.Client{}
	resp, err := client.Do(httpreq)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	return json.Unmarshal(body, respObj)
}

func ApiPost(host, path, method string, req bodyer, respObj interface{}) error {
	url := composeApiUrl(host, path, req)
	urlStr := url.String()

	body, err := req.intoBody()
	if err != nil {
		return err
	}
	httpreq, err := http.NewRequest(method, urlStr, bytes.NewReader(body))
	if err != nil {
		return err
	}

	//TODO: composeWithAuthorization
	client := &http.Client{}
	resp, err := client.Do(httpreq)
	if err != nil {
		return err
	}
	strBody, _ := ioutil.ReadAll(resp.Body)
	return json.Unmarshal(strBody, respObj)
}

func composeApiUrl(host, path string, req interface{}) *url.URL {
	base, err := url.Parse(host)
	base.Path = path
	values := url.Values{}

	if valuer, ok := req.(urlRequestParse); ok {
		values = valuer.intoURLValues()
		valuer.intoURLPathParams(base)
	}

	if err != nil {
		panic(fmt.Sprintf("apiHost invalid host=%s err=%+v", host, err))
	}

	base.RawQuery = values.Encode()

	return base
}


