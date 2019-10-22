package common

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// http get request
func HttpGet(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	s,err:=ioutil.ReadAll(resp.Body)
	return string(s)
}

// http get and convert response into a JSON object
func HttpGetJSON(url string, jsonObj interface{})  {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	s,err:=ioutil.ReadAll(resp.Body)

	json.Unmarshal([]byte(s), &jsonObj)
}
