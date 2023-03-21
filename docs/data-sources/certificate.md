---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sectigo_certificate Data Source - terraform-provider-sectigo"
subcategory: ""
description: |-
  Fetches SSL certificate.
---

# sectigo_certificate (Data Source)

Fetches SSL certificate.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (Number) Certificate ID.

### Read-Only

- `certificate_details` (Attributes) Certificate details. (see [below for nested schema](#nestedatt--certificate_details))
- `serial_number` (String) Certificate serial number.
- `status` (String) Certificate status.

<a id="nestedatt--certificate_details"></a>
### Nested Schema for `certificate_details`

Read-Only:

- `issuer` (String) Certificate issuer.

