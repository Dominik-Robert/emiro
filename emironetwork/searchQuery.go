package emironetwork

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/tidwall/gjson"
)

func searchSpecificWithElastic(host string, port int, index string, insecure bool, query string, count int32, all bool) (*Answer, error) {
	elkQuery := `
	{
		"size": ` + fmt.Sprint(count) + `,
		"query": {
			"match_phrase": {
				"name": "` + query + `"
			}
		}
	}
	`

	body := bytes.NewBufferString(elkQuery)
	response, err := DoCurl(host, port, "GET", index, "/_search", insecure, body)

	if err != nil {
		return nil, errors.New("error at http Request:" + err.Error())
	}

	var answer *Answer
	jsonData := gjson.ParseBytes(response).Get("hits.hits").Array()
	for _, value := range jsonData {
		answer = &Answer{
			Name:        value.Get("_source.name").String(),
			Description: value.Get("_source.description").String(),
			Command:     value.Get("_source.command").String(),
			Language:    value.Get("_source.language").String(),
			Script:      value.Get("_source.script").Bool(),
			Os:          parseStringArray(value.Get("_source.os")),
			Path:        value.Get("_source.path").String(),
			Interactive: value.Get("_source.interactive").Bool(),
			Params:      parseMap(value.Get("_source.params")),
		}
	}
	return answer, nil
}

func parseMap(result gjson.Result) map[string]string {
	returnValue := make(map[string]string)

	for key, value := range result.Map() {
		returnValue[key] = value.String()
	}
	return returnValue
}

func parseStringArray(result gjson.Result) []string {
	returnValue := make([]string, len(result.Array()))
	for i, value := range result.Array() {
		returnValue[i] = value.String()
	}

	return returnValue
}

func searchQueryWithElastic(host string, port int, index string, insecure bool, query string, count int32, all bool) ([]*Answer, error) {
	elkQuery := `
	{
		"size": ` + fmt.Sprint(count) + `,
		"query": {
			"match_phrase": {
				"name": "` + query + `"
			}
		}
	}
	`

	if all {
		elkQuery = `
		{
			"size": ` + fmt.Sprint(count) + `,
			"query": {
				"match_all": {}
			}
		}
		`
	}

	body := bytes.NewBufferString(elkQuery)
	response, err := DoCurl(host, port, "GET", index, "/_search", insecure, body)

	if err != nil {
		return nil, errors.New("error at http Request:" + err.Error())
	}

	var answer []*Answer
	jsonData := gjson.ParseBytes(response).Get("hits.hits").Array()

	for _, value := range jsonData {
		entry := &Answer{
			Name:        value.Get("_source.name").String(),
			Description: value.Get("_source.description").String(),
			Command:     value.Get("_source.command").String(),
			Language:    value.Get("_source.language").String(),
			Script:      value.Get("_source.script").Bool(),
			Os:          parseStringArray(value.Get("_source.os")),
			Path:        value.Get("_source.path").String(),
			Interactive: value.Get("_source.interactive").Bool(),
		}
		answer = append(answer, entry)
	}

	return answer, nil
}

func createNewWithElastic(host string, port int, index string, insecure bool, query []byte) (bool, error) {
	elkQuery := string(query)

	body := bytes.NewBufferString(elkQuery)
	response, err := DoCurl(host, port, "POST", index, "/_doc", insecure, body)

	if err != nil {
		return false, errors.New("error at http Request:" + err.Error())
	}

	if gjson.ParseBytes(response).Get("result").String() == "created" {
		return true, nil
	} else {
		log.Println(gjson.ParseBytes(response).Get("result").String())
		return false, nil
	}

}

func DoCurl(host string, port int, method string, index string, additionalURL string, insecure bool, query io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, "http://"+host+":"+fmt.Sprint(port)+"/"+index+additionalURL, query)

	if err != nil {
		log.Println("Unable to create Request:" + err.Error())
		return nil, errors.New("Unable to create Request:" + err.Error())
	}
	req.Header.Add("content-type", "application/json")

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: insecure}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Println("Failed to send request:" + err.Error())
		return nil, errors.New("Failed to send request:" + err.Error())
	}

	bodyData, _ := ioutil.ReadAll(res.Body)

	isError := gjson.ParseBytes(bodyData).Get("error.root_cause.0.reason").String()
	if isError == strings.Trim("isError", " ") {
		return nil, errors.New(isError)
	}

	return bodyData, nil
}
