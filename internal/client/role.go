package client

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const rolesEndpoint string = "/v1/roles"

type Role struct {
	ID                string   `json:"id,omitempty"`
	Name              string   `json:"name,omitempty"`
	Description       string   `json:"description,omitempty"`
	SubOrgsExpression string   `json:"suborgs_expression,omitempty"`
	Privileges        []string `json:"privileges"`
	ApplyToOrgs       []string `json:"apply_to_orgs"`
	AllReadPrivs      *bool    `json:"all_read_privileges,omitempty"`
	AllWritePrivs     *bool    `json:"all_write_privileges,omitempty"`
	AllSubOrgs        *bool    `json:"all_suborgs,omitempty"`
}

func NewRole(d *schema.ResourceData) *Role {
	res := &Role{}
	if d.HasChange("name") {
		res.Name = d.Get("name").(string)
	}
	res.Description = d.Get("description").(string)

	s := d.Get("suborgs_expression").(string)
	res.SubOrgsExpression = s

	p := d.Get("privileges").(*schema.Set)
	res.Privileges = ResourceTypeSetToStringSlice(p)

	res.ApplyToOrgs = ConfigToStringSlice("apply_to_orgs", d)

	allRead := d.Get("all_read_privileges").(bool)
	res.AllReadPrivs = &allRead

	allWrite := d.Get("all_write_privileges").(bool)
	res.AllWritePrivs = &allWrite

	AllSubOrgs := d.Get("all_suborgs").(bool)
	res.AllSubOrgs = &AllSubOrgs
	return res
}

func parseRole(resp []byte) (*Role, error) {
	pg := &Role{}
	err := json.Unmarshal(resp, pg)
	if err != nil {
		return nil, fmt.Errorf("could not parse role response: %v", err)
	}
	return pg, nil
}

func CreateRole(ctx context.Context, c *Client, r *Role) (*Role, error) {
	neUrl := fmt.Sprintf("%s%s", c.BaseURL, rolesEndpoint)
	body, err := json.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not convert role to json: %v", err)
	}
	resp, err := c.Post(ctx, neUrl, body)
	if err != nil {
		return nil, err
	}
	return parseRole(resp)
}

func UpdateRole(ctx context.Context, c *Client, rID string, r *Role) (*Role, error) {
	neUrl := fmt.Sprintf("%s%s/%s", c.BaseURL, rolesEndpoint, rID)
	body, err := json.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not convert role to json: %v", err)
	}
	resp, err := c.Patch(ctx, neUrl, body)
	if err != nil {
		return nil, err
	}
	return parseRole(resp)
}

func GetRoleByID(ctx context.Context, c *Client, rID string) (*Role, error) {
	url := fmt.Sprintf("%s%s/%s", c.BaseURL, rolesEndpoint, rID)
	resp, err := c.Get(ctx, url, nil)
	if err != nil {
		return nil, err
	}
	return parseRole(resp)
}
func GetRoleByName(ctx context.Context, c *Client, name string) (*Role, error) {
	url := fmt.Sprintf("%s%s", c.BaseURL, rolesEndpoint)
	resp, err := c.Get(ctx, url, nil)
	if err != nil {
		return nil, err
	}
	roles := &[]Role{}
	err = json.Unmarshal(resp, roles)
	if err != nil {
		return nil, fmt.Errorf("could not parse role response: %v", err)
	}
	for _, r := range *roles {
		if r.Name == name {
			return &r, nil
		}
	}
	return nil, nil
}

func DeleteRole(ctx context.Context, c *Client, rID string) (*Role, error) {
	url := fmt.Sprintf("%s%s/%s", c.BaseURL, rolesEndpoint, rID)
	resp, err := c.Delete(ctx, url, nil)
	if err != nil {
		return nil, err
	}
	return parseRole(resp)
}
