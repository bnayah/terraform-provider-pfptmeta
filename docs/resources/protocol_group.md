---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "pfptmeta_protocol_group Resource - terraform-provider-pfptmeta"
subcategory: ""
description: |-
  Protocol Groups are protocols and ports that must be included into granular policies.
---

# pfptmeta_protocol_group (Resource)

Protocol Groups are protocols and ports that must be included into granular policies.

## Example Usage

```terraform
resource "pfptmeta_protocol_group" "new_protocol" {
  name = "NEW_PROTOCOL"
  protocols {
    from_port = 445
    to_port   = 445
    proto     = "udp"
  }
  protocols {
    from_port = 446
    to_port   = 446
    proto     = "tcp"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **name** (String)
- **protocols** (Block List, Min: 1) A list of protocols (see [below for nested schema](#nestedblock--protocols))

### Optional

- **description** (String)

### Read-Only

- **id** (String) The ID of this resource.

<a id="nestedblock--protocols"></a>
### Nested Schema for `protocols`

Required:

- **from_port** (Number)
- **proto** (String) Protocol type, can be one of: tcp, udp, icmp
- **to_port** (Number)

