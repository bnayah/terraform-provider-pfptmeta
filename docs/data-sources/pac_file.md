---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "Data Source pfptmeta_pac_file - terraform-provider-pfptmeta"
subcategory: "Web Security Resources"
description: |-
  Web traffic inspection is further enhanced by means of traffic steering rules (Implemented as Proxy Auto Config file), installed by the web security engine after the user has been onboarded. It is a JavaScript-based file that uses logical statements to determine which traffic is routed through the proxy and which traffic bypasses it. Each tenant has a default rule supplied by Proofpoint, based on the best practice recommendations. However, the administrators can decide to override the default rule with customized traffic steering rules that are better suited for their organization. Once created, the traffic steering rule is distributed to intended users to be hosted locally on their machines. Afterwards, the rule can be updated at any time.
---

# Data Source (pfptmeta_pac_file)

Web traffic inspection is further enhanced by means of traffic steering rules (Implemented as Proxy Auto Config file), installed by the web security engine after the user has been onboarded. It is a JavaScript-based file that uses logical statements to determine which traffic is routed through the proxy and which traffic bypasses it. Each tenant has a default rule supplied by Proofpoint, based on the best practice recommendations. However, the administrators can decide to override the default rule with customized traffic steering rules that are better suited for their organization. Once created, the traffic steering rule is distributed to intended users to be hosted locally on their machines. Afterwards, the rule can be updated at any time.

## Example Usage

```terraform
data "pfptmeta_pac_file" "pac_file" {
  id = "alr-123abc"
}

output "pac_file" {
  value = data.pfptmeta_pac_file.pac_file
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `apply_to_org` (Boolean) Indicates whether this PAC file applies to the org.
- `content` (String) The content of the PAC file
- `description` (String)
- `enabled` (Boolean)
- `exempt_sources` (List of String) Subgroup of `sources` on which the PAC file should not be applied.
- `has_content` (Boolean) Whether the PAC file object has content associated with it.
- `id` (String) The ID of this resource.
- `name` (String)
- `priority` (Number) Determines the order in which the PAC files are being matched. Lower priority value means the PAC file will be matched earlier.
- `sources` (List of String) Users and groups on which the PAC file should be applied.