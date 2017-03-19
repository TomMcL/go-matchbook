package matchbook

import (
	"crypto/rand"
	"crypto/tls"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

var connectionEndpoints = map[string][]string{
	"login":      {"https://api.matchbook.com/bpapi/rest/security/session", "POST"},
	"logout":     {"https://api.matchbook.com/bpapi/rest/security/session", "DELETE"},
	"getSession": {"https://api.matchbook.com/bpapi/rest/security/session", "GET"},
}

type Config struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Session struct {
	config     *Config
	httpClient *http.Client
}

type RequestSpecification struct {
	Url  string
	Type string
}

// Create a new session.
func (c *Config) NewSession() (*Session, error) {

	s := new(Session)

	// Configuration
	if c.Username == "" {
		return s, errors.New("Config.Username is empty.")
	}
	if c.Password == "" {
		return s, errors.New("Config.Password is empty.")
	}

	s.config = c

	/*
		// HTTP client
		cert, err := tls.LoadX509KeyPair(s.config.CertFile, s.config.KeyFile)
		if err != nil {
			return s, err
		}*/
	ssl := &tls.Config{
		InsecureSkipVerify: true,
	}
	ssl.Rand = rand.Reader
	s.httpClient = &http.Client{
		Transport: &http.Transport{
			Dial: func(network, addr string) (net.Conn, error) {
				return net.DialTimeout(network, addr, time.Duration(time.Second*3))
			},
			TLSClientConfig: ssl,
		},
	}

	return s, nil
}

// Builds URLs for API methods.
func (s *Session) getRequestSpec(key, method string) (RequestSpecification, error) {
	if _, exists := connectionEndpoints[key]; exists == false {
		return RequestSpecification{}, errors.New("Invalid endpoint key: " + key)
	}
	url := connectionEndpoints[key][0] + method
	requestType := connectionEndpoints[key][1]
	if requestType == "GET" {
		url += "/"
	}
	return RequestSpecification{url, requestType}, nil
}

func (s *Session) doRequest(key, method string, body *strings.Reader) ([]byte, error) {

	reqSpec, err := s.getRequestSpec(key, method)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(reqSpec.Type, reqSpec.Url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("X-Authentication", s.token)

	res, err := s.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, errors.New(res.Status)
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *Session) Login() ([]byte, error) {

	body := strings.NewReader("username=" + s.config.Username + "&password=" + s.config.Password)
	resp, err := s.doRequest("login.json", "", body)
	return resp, err

}
