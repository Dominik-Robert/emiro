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

func searchSpecificWithElastic(host string, port int, index string, insecure bool, query string, count int32, all bool) *Answer {
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
		log.Fatal(err)
	}

	var answer *Answer
	jsonData := gjson.ParseBytes(response).Get("hits.hits").Array()

	if len(jsonData) > 1 {
		log.Fatalf("More than one result")
	}

	for _, value := range jsonData {
		answer = &Answer{
			Name:        value.Get("_source.name").String(),
			Description: value.Get("_source.description").String(),
			Command:     value.Get("_source.command").String(),
			Language:    value.Get("_source.language").String(),
			Script:      value.Get("_source.script").Bool(),
			Os:          parseStringArray(value.Get("_source.os")),
			Path:        value.Get("_source.path").String(),
		}
	}
	return answer
}

func parseStringArray(result gjson.Result) []string {
	returnValue := make([]string, len(result.Array()))
	for i, value := range result.Array() {
		returnValue[i] = value.String()
	}

	return returnValue
}

func searchQueryWithElastic(host string, port int, index string, insecure bool, query string, count int32, all bool) []*QueryAnswer {
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
		log.Fatal(err)
	}

	var answer []*QueryAnswer
	jsonData := gjson.ParseBytes(response).Get("hits.hits").Array()

	for _, value := range jsonData {
		entry := &QueryAnswer{
			Name:        value.Get("_source.name").String(),
			Description: value.Get("_source.description").String(),
		}
		answer = append(answer, entry)
	}

	return answer
}

func DoCurl(host string, port int, method string, index string, additionalURL string, insecure bool, query io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, "http://"+host+":"+fmt.Sprint(port)+"/"+index+additionalURL, query)

	if err != nil {
		log.Fatalf("Unable to create request: %s", err)
	}
	req.Header.Add("content-type", "application/json")

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: insecure}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatalf("Failed to send request: %s", err)
	}

	bodyData, _ := ioutil.ReadAll(res.Body)

	isError := gjson.ParseBytes(bodyData).Get("error.root_cause.0.reason").String()
	if isError == strings.Trim("isError", " ") {
		return nil, errors.New(isError)
	}

	return bodyData, nil
}
