---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Configuration: https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/iap/TunnelDestGroup.yaml
#     Template:      https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/datasource_iam.html.markdown.tmpl
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Identity-Aware Proxy"
description: |-
  A datasource to retrieve the IAM policy state for Identity-Aware Proxy TunnelDestGroup
---


# google_iap_tunnel_dest_group_iam_policy

Retrieves the current IAM policy data for tunneldestgroup


## Example Usage


```hcl
data "google_iap_tunnel_dest_group_iam_policy" "policy" {
  project = google_iap_tunnel_dest_group.dest_group.project
  region = google_iap_tunnel_dest_group.dest_group.region
  dest_group = google_iap_tunnel_dest_group.dest_group.group_name
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional) The region of the tunnel group. Must be the same as the network resources in the group.
 Used to find the parent resource to bind the IAM policy to. If not specified,
  the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
  region is specified, it is taken from the provider configuration.
* `dest_group` - (Required) Used to find the parent resource to bind the IAM policy to

* `project` - (Optional) The ID of the project in which the resource belongs.
    If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.

## Attributes Reference

The attributes are exported:

* `etag` - (Computed) The etag of the IAM policy.

* `policy_data` - (Required only by `google_iap_tunnel_dest_group_iam_policy`) The policy data generated by
  a `google_iam_policy` data source.
