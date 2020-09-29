package tomtom

import (
	"fmt"
	"net/url"
	"strconv"
)

type FuzzySearchResponse struct {
	Summary Summary  `json:"summary"`
	Results []Result `json:"results"`
}
type Summary struct {
	Query        string   `json:"query"`
	QueryType    string   `json:"queryType"`
	QueryTime    int      `json:"queryTime"`
	NumResults   int      `json:"numResults"`
	Offset       int      `json:"offset"`
	TotalResults int      `json:"totalResults"`
	FuzzyLevel   int      `json:"fuzzyLevel"`
	GeoBias      Position `json:"geoBias"`
}

type Result struct {
	Type        string       `json:"type"`
	ID          string       `json:"id"`
	Score       float64      `json:"score"`
	Dist        float64      `json:"dist"`
	Info        string       `json:"info"`
	Poi         Poi          `json:"poi"`
	Address     Address      `json:"address"`
	Position    Position     `json:"position"`
	Viewport    Viewport     `json:"viewport"`
	EntryPoints []EntryPoint `json:"entryPoints"`
	DataSources DataSources  `json:"dataSources"`
}

type Poi struct {
	Name            string           `json:"name"`
	Phone           string           `json:"phone"`
	CategorySet     []CategorySet    `json:"categorySet"`
	Categories      []string         `json:"categories"`
	Classifications []Classification `json:"classifications"`
}

type Classification struct {
	Code  string `json:"code"`
	Names []Name `json:"names"`
}

type Name struct {
	NameLocale string `json:"nameLocale"`
	Name       string `json:"name"`
}

type CategorySet struct {
	ID int `json:"id"`
}

type Address struct {
	StreetNumber           string `json:"streetNumber"`
	StreetName             string `json:"streetName"`
	Municipality           string `json:"municipality"`
	CountrySubdivision     string `json:"countrySubdivision"`
	CountrySubdivisionName string `json:"countrySubdivisionName"`
	PostalCode             string `json:"postalCode"`
	ExtendedPostalCode     string `json:"extendedPostalCode"`
	CountryCode            string `json:"countryCode"`
	Country                string `json:"country"`
	CountryCodeISO3        string `json:"countryCodeISO3"`
	FreeformAddress        string `json:"freeformAddress"`
	LocalName              string `json:"localName"`
}

type Position struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type Viewport struct {
	TopLeftPoint  Position `json:"topLeftPoint"`
	BtmRightPoint Position `json:"btmRightPoint"`
}

type PoiDetail struct {
	ID         string `json:"id"`
	SourceName string `json:"sourceName"`
}

type EntryPoint struct {
	Type     string   `json:"type"`
	Position Position `json:"position"`
}

type DataSources struct {
	PoiDetails []PoiDetail `json:"poiDetails"`
	Geometry   Geometry    `json:"geometry"`
}

type Geometry struct {
	ID string `json:"id"`
}

type FuzzySearchRequest struct {
	Position Position
	Radius   int
	Query    string
}

func (r Result) DataSourceID() string {
	if r.DataSources.PoiDetails == nil {
		return r.DataSources.Geometry.ID
	}
	return r.DataSources.PoiDetails[0].ID
}

func (c *Client) FuzzySearch(req *FuzzySearchRequest) (*FuzzySearchResponse, error) {
	data := url.Values{}
	data.Add("lat", fmt.Sprintf("%g", req.Position.Lat))
	data.Add("lon", fmt.Sprintf("%g", req.Position.Lon))
	data.Add("radius", strconv.Itoa(req.Radius))
	data.Add("typeahead", "false")
	data.Add("idxSet", "Geo,POI")

	var response FuzzySearchResponse
	if err := c.Call(CmdFuzzySearch, req.Query, data, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
