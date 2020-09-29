package tomtom

import (
	"errors"
	"net/url"
	"os"
	"testing"
)

func testClient() *Client {
	apiKey, ok := os.LookupEnv("TOMTOM_API_KEY")
	if !ok {
		panic(errors.New("no api keyprovided in environment"))
	}

	c, err := NewClient(&Config{APIKey: apiKey})
	if err != nil {
		panic(err)
	}
	return c

}

func TestCall(t *testing.T) {
	client := testClient()

	failCommand := "doesntexist"
	var resp FuzzySearchResponse
	if err := client.Call(failCommand, "", url.Values{}, &resp); err == nil {
		t.Fatalf("Should have failed with non-supported command of %s", failCommand)
	}

}
