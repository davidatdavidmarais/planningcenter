package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
)

func setBearerToken(req *http.Request, token string) {
	req.Header.Add("Authorization", "Bearer "+token)
}

func BodyBuffer(req interface{}) (io.Reader, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(body), nil
}

func HttpRequestWithBearer(cl *http.Client, req *http.Request, token string, resp interface{}) error {
	setBearerToken(req, token)
	return HttpRequest(cl, req, resp)
}

func HttpRequest(cl *http.Client, req *http.Request, resp interface{}) error {
	req.Header.Add("content-type", "application/json")

	httpResp, err := cl.Do(req)
	if err != nil {
		return err
	}

	defer httpResp.Body.Close()
	respBytes, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return err
	}

	if httpResp.StatusCode != 200 {
		return errors.New(string(respBytes))
	}

	err = json.Unmarshal(respBytes, resp)
	if err != nil {
		return err
	}

	return nil
}
