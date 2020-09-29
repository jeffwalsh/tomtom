package tomtom

import (
	"testing"
)

func Test_POIPhoto(t *testing.T) {
	c := testClient()

	_, err := c.POIPhoto(&POIPhotoRequest{
		ID: "c6a048f2-cd8b-3b26-bcc2-4ae3c458ab15",
	})
	if err != nil {
		t.Fatal(err)
	}
}
