package client

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

func Post(client *http.Client, apiURL string, bodyMap map[string]interface{}) (*Result, error) {
	bodyRaw, _ := json.Marshal(bodyMap)

	resp, err := client.Post(apiURL, "application/json", bytes.NewBuffer(bodyRaw))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Errorf("status: %v  body: %v", resp.StatusCode, string(respBody))
	}

	result := new(Result)
	err = json.Unmarshal(respBody, result)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return result, nil
}
