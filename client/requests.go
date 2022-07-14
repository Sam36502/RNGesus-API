package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Sam36502/RNGesus-API/dto"
)

type RNGesusClient struct {
	BaseURL string
}

func (c *RNGesusClient) GetRandomFloat() (float64, error) {
	resp, err := c.executeRequest(http.MethodGet, "/v1/rand/float", nil)
	if err != nil {
		return 0, err
	}

	var numMsg dto.FloatResponse
	err = json.Unmarshal(resp, &numMsg)
	if err != nil {
		return 0, err
	}

	return numMsg.Number, nil
}

func (c *RNGesusClient) GetRandomInt() (int64, error) {
	resp, err := c.executeRequest(http.MethodGet, "/v1/rand/int", nil)
	if err != nil {
		return 0, err
	}

	var numMsg dto.IntResponse
	err = json.Unmarshal(resp, &numMsg)
	if err != nil {
		return 0, err
	}

	return numMsg.Number, nil
}

func (c *RNGesusClient) GetRandomIntInRange(min, max int64) (int64, error) {
	params := map[string]string{
		dto.PARAM_INT_MIN: fmt.Sprint(min),
		dto.PARAM_INT_MAX: fmt.Sprint(max),
	}

	resp, err := c.executeRequest(http.MethodGet, "/v1/rand/int", params)
	if err != nil {
		return 0, err
	}

	var numMsg dto.IntResponse
	err = json.Unmarshal(resp, &numMsg)
	if err != nil {
		return 0, err
	}

	return numMsg.Number, nil
}

func (c *RNGesusClient) PrayForFloat(float float64) (string, error) {
	params := map[string]string{
		dto.PARAM_PRAYER: fmt.Sprint(float),
	}

	resp, err := c.executeRequest(http.MethodGet, "/v1/pray/float", params)
	if err != nil {
		return string(resp), err
	}

	var msg dto.MessageResponse
	err = json.Unmarshal(resp, &msg)
	if err != nil {
		return "", err
	}

	return msg.Message, nil
}

func (c *RNGesusClient) executeRequest(method, path string, params map[string]string) ([]byte, error) {
	req, err := http.NewRequest(method, fmt.Sprint(c.BaseURL, path), nil)
	if err != nil {
		return nil, err
	}

	// Add params
	if params != nil {
		q := req.URL.Query()
		for k, v := range params {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	// Execute request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read response
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Check status
	if resp.StatusCode != dto.STATUS_OK {
		var msg dto.MessageResponse
		err := json.Unmarshal(respBody, &msg)
		if err != nil {
			return []byte(fmt.Sprint(resp.Status, "\n", string(respBody))), err
		} else {
			return []byte(fmt.Sprint(resp.Status, "\n", msg.Message)), err
		}
	}

	return respBody, nil
}
