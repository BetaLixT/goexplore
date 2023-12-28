package notion

import (
	"github.com/Soreing/gent"
)

type Client struct {
	client *gent.Client

	apiKey    string
	apiVer    string
	notionVer string
}

// func (c *Client) Get(
// 	ctx context.Context,
// 	hearingIds []string,
// ) ([]HearingGlance, error) {
// 	key, val, err := c.tknprov.GetToken(ctx)
// 	if err != nil {
// 		return nil, fmt.Errorf("error fetching token: %w", err)
// 	}
//
// 	resp, err := c.client.Get(
// 		ctx,
// 		c.baseUrl+"/api/v1/glances/hearings",
// 		nil, nil,
// 		map[string]string{
// 			key: val,
// 		},
// 		map[string][]string{
// 			"ids": hearingIds,
// 		},
// 	)
// 	if err != nil {
// 		return nil, fmt.Errorf("error making request: %w", err)
// 	}
// 	if resp.StatusCode != 200 {
// 		return nil, &RequestFailed{
// 			Response: resp,
// 		}
// 	}
//
// 	room := []HearingGlance{}
// 	dat, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, err
// 	}
// 	err = json.Unmarshal(dat, &room)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return room, nil
// }
