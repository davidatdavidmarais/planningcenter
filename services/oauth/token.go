package oauth

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	EXCHANGE_TOKEN_GRANT_TYPE = "authorization_code"
	REFRESH_TOKEN_GRANT_TYPE  = "refresh_token"

	PATH = "/oauth/token"
)

type OAuthClient struct {
	config Config
	cl     *http.Client
}

func New(cl *http.Client, c Config) *OAuthClient {
	return &OAuthClient{
		cl:     cl,
		config: c,
	}
}

func (a *OAuthClient) ExchangeToken(code, redirectURI string) (*ExchangeTokenResponse, error) {
	req := &ExchangeTokenRequest{
		GrantType:    EXCHANGE_TOKEN_GRANT_TYPE,
		Code:         code,
		ClientID:     a.config.ClientID,
		ClientSecret: a.config.ClientSecret,
		RedirectURI:  redirectURI,
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(body)

	httpReq, err := http.NewRequest(http.MethodPost, a.config.BaseURL+PATH, buf)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Add("content-type", "application/json")

	resp, err := a.cl.Do(httpReq)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(string(respBytes))
	}

	respObj := new(ExchangeTokenResponse)
	err = json.Unmarshal(respBytes, respObj)
	if err != nil {
		return nil, err
	}

	return respObj, nil
}

func (a *OAuthClient) RefreshToken(refreshToken string) (*ExchangeTokenResponse, error) {
	req := &RefreshTokenRequest{
		GrantType:    REFRESH_TOKEN_GRANT_TYPE,
		ClientID:     a.config.ClientID,
		ClientSecret: a.config.ClientSecret,
		RefreshToken: refreshToken,
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(body))

	buf := bytes.NewBuffer(body)

	httpReq, err := http.NewRequest(http.MethodPost, a.config.BaseURL+PATH, buf)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Add("content-type", "application/json")

	resp, err := a.cl.Do(httpReq)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(string(respBytes))
	}

	respObj := new(ExchangeTokenResponse)
	err = json.Unmarshal(respBytes, respObj)
	if err != nil {
		return nil, err
	}

	return respObj, nil
}
