---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Configuration: https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/securesourcemanager/Repository.yaml
#     Template:      https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/datasource_iam.html.markdown.tmpl
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Secure Source Manager"
description: |-
  A datasource to retrieve the IAM policy state for Secure Source Manager Repository
---


# google_secure_source_manager_repository_iam_policy

Retrieves the current IAM policy data for repository


## Example Usage


```hcl
data "google_secure_source_manager_repository_iam_policy" "policy" {
  project = google_secure_source_manager_repository.default.project
  location = google_secure_source_manager_repository.default.location
  repository_id = google_secure_source_manager_repository.default.repository_id
}
```

## Argument Reference

The following arguments are supported:

* `location` - (Optional) The location for the Repository.
 Used to find the parent resource to bind the IAM policy to. If not specified,
  the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
  location is specified, it is taken from the provider configuration.
* `repository_id` - (Required) Used to find the parent resource to bind the IAM policy to

* `project` - (Optional) The ID of the project in which the resource belongs.
    If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.

## Attributes Reference

The attributes are exported:

* `etag` - (Computed) The etag of the IAM policy.

* `policy_data` - (Required only by `google_secure_source_manager_repository_iam_policy`) The policy data generated by
  a `google_iam_policy` data source.
