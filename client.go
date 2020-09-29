package tomtom

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	baseURL    = "https://api.tomtom.com"
	versionTwo = "2"
	extJson    = "json"
)

var (
	CmdFuzzySearch = "search"
	CmdPOIDetails  = "poiDetails"
	CmdPOIPhoto    = "poiPhoto"
)

type Client struct {
	HTTPClient *http.Client
	key        string
	baseURL    string
	version    string
	ext        string
	commands   []string
}

type Config struct {
	APIKey string
}

// SupportedCommands returns a slice of strings with all the available commands
func SupportedCommands() []string {
	return []string{CmdFuzzySearch, CmdPOIDetails, CmdPOIPhoto}
}
func NewClient(cfg *Config, options ...func(c *Client)) (*Client, error) {
	if cfg.APIKey == "" {
		return nil, ErrNoAPIKey
	}

	c := &Client{
		HTTPClient: &http.Client{},
		key:        cfg.APIKey,
		baseURL:    baseURL,
		version:    versionTwo,
		ext:        extJson,
		commands:   SupportedCommands(),
	}

	for _, option := range options {
		option(c)
	}

	return c, nil
}

// Call sends a request with the given cmd and data, and then unmarshals the response into the given responseStruct.
// If you don't want to pass a Reader to the individual functions, you can use Call to call any command directly.
// You just need to build out your own data with an url.Values, and create an instance of the responsestruct expected by your command.
// ie: data := url.Values{}
// data.Add("lat", "42")
// data.Add("lon", "40")
// data.Add("radius", "10000")
// resp := tomtom.FuzzySearchResponse{}
// err := c.Call(tomtom.CmdFuzzySearch, data, &resp)
func (c *Client) Call(cmd string, query string, data url.Values, responseStruct interface{}) error {
	if !stringExistsInSlice(c.commands, cmd) {
		return ErrCommandDoesntExist
	}
	return c.call(cmd, query, data, responseStruct)
}

// call sends a request with the given cmd and data, and then unmarshals the response into the given responseStruct.
func (c *Client) call(cmd string, query string, data url.Values, responseStruct interface{}) error {
	data.Add("key", c.key)

	var url string
	if cmd == CmdFuzzySearch {
		url = fmt.Sprintf("%v/search/%v/%v/%v.%v", c.baseURL, c.version, cmd, query, c.ext)
	} else if cmd == CmdPOIPhoto {
		url = fmt.Sprintf("%v/search/%v/%v", c.baseURL, c.version, cmd)
	} else {

		url = fmt.Sprintf("%v/search/%v/%v.%v", c.baseURL, c.version, cmd, c.ext)
	}
	// create the request using the url and url values
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.URL.RawQuery = data.Encode()

	// do the actual request
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		ret := fmt.Sprintf("failed to make api call: expected status %v, got %v", 200, resp.StatusCode)
		return errors.New(ret)
	}

	// return the unmarshalled response and an error if it occurred
	if cmd != CmdPOIPhoto {
		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			return err
		}

		return json.Unmarshal(body, responseStruct)
	} else {
		var buff bytes.Buffer
		// nolint
		buff.ReadFrom(resp.Body)

		encodedString := base64.StdEncoding.EncodeToString(buff.Bytes())

		return json.Unmarshal([]byte(fmt.Sprintf(`{"encoded_string":"%v"}`, encodedString)), responseStruct)

	}
}

// contains is a helper to determine whether an element exists in a slice of string
func stringExistsInSlice(strings []string, e string) bool {
	for _, a := range strings {
		if a == e {
			return true
		}
	}
	return false
}
