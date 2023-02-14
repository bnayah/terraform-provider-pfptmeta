package client

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const protocolGroupsEndpoint string = "/v1/protocol_groups"

type Protocol struct {
	FromPort int    `json:"from_port"`
	ToPort   int    `json:"to_port"`
	Proto    string `json:"proto"`
}

type ProtocolGroup struct {
	ID          string     `json:"id,omitempty"`
	Name        string     `json:"name,omitempty"`
	Description string     `json:"description,omitempty"`
	Protocols   []Protocol `json:"protocols,omitempty"`
}

func NewProtocolGroup(d *schema.ResourceData) *ProtocolGroup {
	res := &ProtocolGroup{}
	if d.HasChange("name") {
		res.Name = d.Get("name").(string)
	}
	res.Description = d.Get("description").(string)
	rawPs := d.Get("protocols").([]interface{})
	ps := make([]Protocol, len(rawPs))
	for i, v := range rawPs {
		p := v.(map[string]interface{})
		ps[i] = Protocol{FromPort: p["from_port"].(int), ToPort: p["to_port"].(int), Proto: p["proto"].(string)}
	}
	res.Protocols = ps
	return res
}

func parseProtocolGroup(resp []byte) (*ProtocolGroup, error) {
	pg := &ProtocolGroup{}
	err := json.Unmarshal(resp, pg)
	if err != nil {
		return nil, fmt.Errorf("could not parse protocol group response: %v", err)
	}
	return pg, nil
}

func CreateProtocolGroup(ctx context.Context, c *Client, pg *ProtocolGroup) (*ProtocolGroup, error) {
	neUrl := fmt.Sprintf("%s%s", c.BaseURL, protocolGroupsEndpoint)
	body, err := json.Marshal(pg)
	if err != nil {
		return nil, fmt.Errorf("could not convert protocol group to json: %v", err)
	}
	resp, err := c.Post(ctx, neUrl, body)
	if err != nil {
		return nil, err
	}
	return parseProtocolGroup(resp)
}

func UpdateProtocolGroup(ctx context.Context, c *Client, pgID string, pg *ProtocolGroup) (*ProtocolGroup, error) {
	neUrl := fmt.Sprintf("%s%s/%s", c.BaseURL, protocolGroupsEndpoint, pgID)
	body, err := json.Marshal(pg)
	if err != nil {
		return nil, fmt.Errorf("could not convert protocol group to json: %v", err)
	}
	resp, err := c.Patch(ctx, neUrl, body)
	if err != nil {
		return nil, err
	}
	return parseProtocolGroup(resp)
}

func GetProtocolGroupById(ctx context.Context, c *Client, pgID string) (*ProtocolGroup, error) {
	url := fmt.Sprintf("%s%s/%s", c.BaseURL, protocolGroupsEndpoint, pgID)
	resp, err := c.Get(ctx, url, nil)
	if err != nil {
		return nil, err
	}
	return parseProtocolGroup(resp)
}
func GetProtocolGroupByName(ctx context.Context, c *Client, name string) (*ProtocolGroup, error) {
	url := fmt.Sprintf("%s%s", c.BaseURL, protocolGroupsEndpoint)
	resp, err := c.Get(ctx, url, nil)
	if err != nil {
		return nil, err
	}
	pgs := &[]ProtocolGroup{}
	err = json.Unmarshal(resp, pgs)
	if err != nil {
		return nil, fmt.Errorf("could not parse protocol group response: %v", err)
	}
	for _, pg := range *pgs {
		if pg.Name == name {
			return &pg, nil
		}
	}
	return nil, nil
}

func DeleteProtocolGroup(ctx context.Context, c *Client, pgID string) (*ProtocolGroup, error) {
	url := fmt.Sprintf("%s%s/%s", c.BaseURL, protocolGroupsEndpoint, pgID)
	resp, err := c.Delete(ctx, url, nil)
	if err != nil {
		return nil, err
	}
	return parseProtocolGroup(resp)
}
