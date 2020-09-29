package tomtom

import (
	"testing"
)

func Test_POIDetails(t *testing.T) {
	c := testClient()

	_, err := c.POIDetails(&POIDetailsRequest{
		ID: "Rm91cnNxdWFyZTo0ZTFiNTM4Y2FlNjBkZGI4ZjY1MWU0M2I=",
	})
	if err != nil {
		t.Fatal(err)
	}
}
