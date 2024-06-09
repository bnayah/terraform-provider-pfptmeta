---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "Data Source pfptmeta_aac_rule - terraform-provider-pfptmeta"
subcategory: "Adaptive Access Control Rule"
description: |-
  Adaptive access control rule for protecting users connecting to service provider application under risky conditions
---

# Data Source (pfptmeta_aac_rule)

Adaptive access control rule for protecting users connecting to service provider application under risky conditions

## Example Usage

```terraform
data "pfptmeta_aac_rule" "aac_rule" {
  id = "arl-XWU8BQmGbevn7"
}

output "aac_rule" {
  value = data.pfptmeta_aac_rule.aac_rule
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `action` (String) The action to enforce when rule is matched to a connection
- `app_ids` (List of String) IDs of the apps that the AAC rule is applied to
- `apply_all_apps` (Boolean) Indicates whether this rule applies to all apps of the org, regardless whether such apps are specified in app_ids. Note: this attribute overrides app_ids
- `certificate_id` (String) The root/intermediate certificate ID of managed devices that the rule is applied to
- `description` (String)
- `enabled` (Boolean)
- `exempt_sources` (List of String) Subgroup of 'sources' to which the AAC rule is not applied
- `filter_expression` (String) Defines filtering expressions to to provide user granularity in AAC rule application
- `id` (String) The ID of this resource.
- `ip_reputations` (List of String) List of IP reputations that the rule is applied to
- `locations` (List of String) List of locations that the rule is applied to. Each country is represented by an Alpha-2 code (ISO-3166). Enum: `AD`,`AE`,`AF`,`AG`,`AI`,`AL`,`AM`,`AO`,`AQ`,`AR`,`AS`,`AT`,`AU`,`AW`,`AX`,`AZ`,`BA`,`BB`,`BD`,`BE`,`BF`,`BG`,`BH`,`BI`,`BJ`,`BL`,`BM`,`BN`,`BO`,`BQ`,`BR`,`BS`,`BT`,`BV`,`BW`,`BY`,`BZ`,`CA`,`CC`,`CD`,`CF`,`CG`,`CH`,`CI`,`CK`,`CL`,`CM`,`CN`,`CO`,`CR`,`CU`,`CV`,`CW`,`CX`,`CY`,`CZ`,`DE`,`DJ`,`DK`,`DM`,`DO`,`DZ`,`EC`,`EE`,`EG`,`EH`,`ER`,`ES`,`ET`,`FI`,`FJ`,`FK`,`FM`,`FO`,`FR`,`GA`,`GB`,`GD`,`GE`,`GF`,`GG`,`GH`,`GI`,`GL`,`GM`,`GN`,`GP`,`GQ`,`GR`,`GS`,`GT`,`GU`,`GW`,`GY`,`HK`,`HM`,`HN`,`HR`,`HT`,`HU`,`ID`,`IE`,`IL`,`IM`,`IN`,`IO`,`IQ`,`IR`,`IS`,`IT`,`JE`,`JM`,`JO`,`JP`,`KE`,`KG`,`KH`,`KI`,`KM`,`KN`,`KP`,`KR`,`KW`,`KY`,`KZ`,`LA`,`LB`,`LC`,`LI`,`LK`,`LR`,`LS`,`LT`,`LU`,`LV`,`LY`,`MA`,`MC`,`MD`,`ME`,`MF`,`MG`,`MH`,`MK`,`ML`,`MM`,`MN`,`MO`,`MP`,`MQ`,`MR`,`MS`,`MT`,`MU`,`MV`,`MW`,`MX`,`MY`,`MZ`,`NA`,`NC`,`NE`,`NF`,`NG`,`NI`,`NL`,`NO`,`NP`,`NR`,`NU`,`NZ`,`OM`,`PA`,`PE`,`PF`,`PG`,`PH`,`PK`,`PL`,`PM`,`PN`,`PR`,`PS`,`PT`,`PW`,`PY`,`QA`,`RE`,`RO`,`RS`,`RU`,`RW`,`SA`,`SB`,`SC`,`SD`,`SE`,`SG`,`SH`,`SI`,`SJ`,`SK`,`SL`,`SM`,`SN`,`SO`,`SR`,`SS`,`ST`,`SV`,`SX`,`SY`,`SZ`,`TC`,`TD`,`TF`,`TG`,`TH`,`TJ`,`TK`,`TL`,`TM`,`TN`,`TO`,`TR`,`TT`,`TV`,`TW`,`TZ`,`UA`,`UG`,`UM`,`US`,`UY`,`UZ`,`VA`,`VC`,`VE`,`VG`,`VI`,`VN`,`VU`,`WF`,`WS`,`YE`,`YT`,`ZA`,`ZM`,`ZW`
- `name` (String)
- `networks` (List of String) List of IP network IDs that the rule is applied to
- `notification_channels` (List of String) List of notification channel IDs
- `priority` (Number) Determines the order in which the aac rules are being matched. Lower priority indicates that the AAC rule is matched earlier
- `sources` (List of String) Users and groups that the rule is applied to