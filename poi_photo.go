package tomtom

import "net/url"

type POIPhotoRequest struct {
	ID string `json:"id"`
}

type POIPhotoResponse struct {
	EncodedString string `json:"encoded_string"`
}

func (c *Client) POIPhoto(req *POIPhotoRequest) (*POIPhotoResponse, error) {
	data := url.Values{}
	data.Add("id", req.ID)

	var response POIPhotoResponse
	if err := c.Call(CmdPOIPhoto, "", data, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
