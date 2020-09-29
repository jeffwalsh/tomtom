package tomtom

import "net/url"

type POIDetailsResponse struct {
	ID     string           `json:"id"`
	Result POIDetailsResult `json:"result"`
}

type POIDetailsResult struct {
	Rating       POIDetailRating   `json:"rating"`
	PriceRange   PriceRange        `json:"priceRange"`
	Photos       []POIDetailPhoto  `json:"photos"`
	Reviews      []POIDetailReview `json:"reviews"`
	PopularHours []PopularHours    `json:"popularHours"`
}

type PopularHours struct {
	DayOfWeek  int         `json:"dayOfWeek"`
	TimeRanges []TimeRange `json:"timeRanges"`
}

type TimeRange struct {
	StartTime Time `json:"startTime"`
	EndTime   Time `json:"endTime"`
}

type Time struct {
	Hour   int `json:"hour"`
	Minute int `json:"minute"`
}

type POIDetailRating struct {
	TotalRatings int     `json:"totalRatings"`
	Value        float64 `json:"value"`
	MinValue     float64 `json:"minValue"`
	MaxValue     float64 `json:"maxValue"`
}

type PriceRange struct {
	Value    int    `json:"value"`
	Label    string `json:"label"`
	MinValue int    `json:"minValue"`
	MaxValue int    `json:"maxValue"`
}

type POIDetailPhoto struct {
	ID string `json:"id"`
}

type POIDetailReview struct {
	Text string `json:"text"`
	Date string `json:"date"`
}

type POIDetailsRequest struct {
	ID string `json:"id"`
}

func (c *Client) POIDetails(req *POIDetailsRequest) (*POIDetailsResponse, error) {
	data := url.Values{}
	data.Add("id", req.ID)

	var response POIDetailsResponse
	if err := c.Call(CmdPOIDetails, "", data, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
