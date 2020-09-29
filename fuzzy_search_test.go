package tomtom

import (
	"testing"
)

func Test_FuzzySearch(t *testing.T) {
	c := testClient()

	_, err := c.FuzzySearch(&FuzzySearchRequest{
		Position: Position{
			Lat: 49.3102797,
			Lon: -123.2653601,
		},
		Radius: 100000,
		Query:  "Kojima Sushi",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func Test_DataSourceID(t *testing.T) {
	r := &FuzzySearchResponse{
		Results: []Result{
			{
				DataSources: DataSources{
					Geometry: Geometry{
						ID: "geometry",
					},
				},
			},
		},
	}

	if r.Results[0].DataSourceID() != "geometry" {
		t.Fatal("expected geometry")
	}
}
