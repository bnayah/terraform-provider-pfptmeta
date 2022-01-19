package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"io/ioutil"
	"net/http"
)

const deviceSettingsEndpoint = "v1/settings/device"

type DeviceSettings struct {
	ID                        string   `json:"id,omitempty"`
	Name                      string   `json:"name,omitempty"`
	Description               string   `json:"description"`
	Enabled                   bool     `json:"enabled"`
	ApplyOnOrg                bool     `json:"apply_on_org"`
	ApplyToEntities           []string `json:"apply_to_entities"`
	AutoFqdnDomainNames       []string `json:"auto_fqdn_domain_names,omitempty"`
	DirectSso                 *string  `json:"direct_sso,omitempty"`
	OverlayMfaRefreshPeriod   *int     `json:"overlay_mfa_refresh_period,omitempty"`
	OverlayMfaRequired        *bool    `json:"overlay_mfa_required,omitempty"`
	ProtocolSelectionLifetime *int     `json:"protocol_selection_lifetime,omitempty"`
	ProxyAlwaysOn             *bool    `json:"proxy_always_on,omitempty"`
	SearchDomains             []string `json:"search_domains,omitempty"`
	SessionLifetime           *int     `json:"session_lifetime,omitempty"`
	SessionLifetimeGrace      *int     `json:"session_lifetime_grace,omitempty"`
	TunnelMode                *string  `json:"tunnel_mode,omitempty"`
	VpnLoginBrowser           *string  `json:"vpn_login_browser,omitempty"`
	ZtnaAlwaysOn              *bool    `json:"ztna_always_on,omitempty"`
}

func NewDeviceSettings(d *schema.ResourceData) *DeviceSettings {
	res := &DeviceSettings{}
	if d.HasChange("name") {
		res.Name = d.Get("name").(string)
	}
	res.Description = d.Get("description").(string)
	res.Enabled = d.Get("enabled").(bool)
	res.ApplyOnOrg = d.Get("apply_on_org").(bool)
	res.ApplyToEntities = ConfigToStringSlice("apply_to_entities", d)
	domainNames := d.Get("auto_fqdn_domain_names").([]interface{})
	if len(domainNames) != 0 {
		res.AutoFqdnDomainNames = make([]string, len(domainNames))
		if domainNames[0] != nil {
			res.AutoFqdnDomainNames[0] = domainNames[0].(string)
		}
	}

	ds, exists := d.GetOk("direct_sso")
	if exists {
		directSso := ds.(string)
		res.DirectSso = &directSso
	}

	omrp, exists := d.GetOk("overlay_mfa_refresh_period")
	if exists {
		overlayMfaRefreshPeriod := omrp.(int)
		res.OverlayMfaRefreshPeriod = &overlayMfaRefreshPeriod
	}

	omr, exists := d.GetOk("overlay_mfa_required")
	if exists {
		overlayMfaRequired := omr.(bool)
		res.OverlayMfaRequired = &overlayMfaRequired
	}

	psl, exists := d.GetOk("protocol_selection_lifetime")
	if exists {
		protocolSelectionLifetime := psl.(int)
		res.ProtocolSelectionLifetime = &protocolSelectionLifetime
	}

	pao, exists := d.GetOk("proxy_always_on")
	if exists {
		proxyAlwaysOn := pao.(bool)
		res.ProxyAlwaysOn = &proxyAlwaysOn
	}
	res.SearchDomains = ConfigToStringSlice("search_domains", d)

	slt, exists := d.GetOk("session_lifetime")
	if exists {
		sessionLifetime := slt.(int)
		res.SessionLifetime = &sessionLifetime
	}
	slg, exists := d.GetOk("session_lifetime_grace")
	if exists {
		sessionLifetimeGrace := slg.(int)
		res.SessionLifetimeGrace = &sessionLifetimeGrace
	}
	tm, exists := d.GetOk("tunnel_mode")
	if exists {
		tunnelMode := tm.(string)
		res.TunnelMode = &tunnelMode
	}
	vlb, exists := d.GetOk("vpn_login_browser")
	if exists {
		vpnLoginBrowser := vlb.(string)
		res.VpnLoginBrowser = &vpnLoginBrowser
	}
	ztnaAlwaysOn, exists := d.GetOk("ztna_always_on")
	if exists {
		alwaysOn := ztnaAlwaysOn.(bool)
		res.ZtnaAlwaysOn = &alwaysOn
	}

	return res
}

func parseDeviceSettings(resp *http.Response) (*DeviceSettings, error) {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read device settings response: %v", err)
	}
	ds := &DeviceSettings{}
	err = json.Unmarshal(body, ds)
	if err != nil {
		return nil, fmt.Errorf("could not parse device settings response: %v", err)
	}
	return ds, nil
}

func CreateDeviceSettings(ctx context.Context, c *Client, ds *DeviceSettings) (*DeviceSettings, error) {
	url := fmt.Sprintf("%s/%s", c.BaseURL, deviceSettingsEndpoint)
	body, err := json.Marshal(ds)
	if err != nil {
		return nil, fmt.Errorf("could not convert device settings to json: %v", err)
	}
	resp, err := c.Post(ctx, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	return parseDeviceSettings(resp)
}

func UpdateDeviceSettings(ctx context.Context, c *Client, dsID string, ds *DeviceSettings) (*DeviceSettings, error) {
	url := fmt.Sprintf("%s/%s/%s", c.BaseURL, deviceSettingsEndpoint, dsID)
	body, err := json.Marshal(ds)
	if err != nil {
		return nil, fmt.Errorf("could not convert device settings to json: %v", err)
	}
	resp, err := c.Patch(ctx, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	return parseDeviceSettings(resp)
}

func GetDeviceSettings(ctx context.Context, c *Client, dsID string) (*DeviceSettings, error) {
	url := fmt.Sprintf("%s/%s/%s", c.BaseURL, deviceSettingsEndpoint, dsID)
	resp, err := c.Get(ctx, url, nil)
	if err != nil {
		return nil, err
	}
	return parseDeviceSettings(resp)
}

func DeleteDeviceSettings(ctx context.Context, c *Client, dsID string) (*DeviceSettings, error) {
	url := fmt.Sprintf("%s/%s/%s", c.BaseURL, deviceSettingsEndpoint, dsID)
	resp, err := c.Delete(ctx, url, nil)
	if err != nil {
		return nil, err
	}
	return parseDeviceSettings(resp)
}