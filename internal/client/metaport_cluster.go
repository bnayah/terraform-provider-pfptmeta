package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"io/ioutil"
	"net/http"
	u "net/url"
)

const (
	metaportClusterEndpoint string = "v1/metaport_clusters"
)

type MetaportCluster struct {
	ID             string   `json:"id,omitempty"`
	Name           string   `json:"name"`
	Description    string   `json:"description,omitempty"`
	MappedElements []string `json:"mapped_elements"`
	Metaports      []string `json:"metaports"`
}

func NewMetaportCluster(d *schema.ResourceData) *MetaportCluster {
	res := &MetaportCluster{}
	if d.HasChange("name") {
		res.Name = d.Get("name").(string)
	}
	res.Description = d.Get("description").(string)

	mes := d.Get("mapped_elements")
	res.MappedElements = ResourceTypeSetToStringSlice(mes.(*schema.Set))

	mps := d.Get("metaports")
	res.Metaports = ResourceTypeSetToStringSlice(mps.(*schema.Set))

	return res
}

func parseMetaportCluster(resp *http.Response) (*MetaportCluster, error) {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	mc := &MetaportCluster{}
	err = json.Unmarshal(body, mc)
	if err != nil {
		return nil, fmt.Errorf("could not parse metaport cluster response: %v", err)
	}
	return mc, nil
}

func CreateMetaportCluster(c *Client, m *MetaportCluster) (*MetaportCluster, error) {
	neUrl := fmt.Sprintf("%s/%s", c.BaseURL, metaportClusterEndpoint)
	body, err := json.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("could not convert metaport cluster to json: %v", err)
	}
	resp, err := c.Post(neUrl, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	return parseMetaportCluster(resp)
}

func GetMetaportCluster(c *Client, mId string) (*MetaportCluster, error) {
	url := fmt.Sprintf("%s/%s/%s", c.BaseURL, metaportClusterEndpoint, mId)
	resp, err := c.Get(url, u.Values{"expand": {"true"}})
	if err != nil {
		return nil, err
	}
	return parseMetaportCluster(resp)
}

func UpdateMetaportCluster(c *Client, mId string, m *MetaportCluster) (*MetaportCluster, error) {
	neUrl := fmt.Sprintf("%s/%s/%s", c.BaseURL, metaportClusterEndpoint, mId)
	body, err := json.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("could not convert metaport cluster to json: %v", err)
	}
	resp, err := c.Patch(neUrl, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	return parseMetaportCluster(resp)
}

func DeleteMetaportCluster(c *Client, neID string) (*MetaportCluster, error) {
	url := fmt.Sprintf("%s/%s/%s", c.BaseURL, metaportClusterEndpoint, neID)
	resp, err := c.Delete(url, nil)
	if err != nil {
		return nil, err
	}
	return parseMetaportCluster(resp)
}