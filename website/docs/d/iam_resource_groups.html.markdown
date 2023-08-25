---
subcategory : "Account Management"
---

# ovh_iam_resource_groups (Data Source)

Use this data source to list the existing IAM policies of an account.

## Important
-> Using this resource requires that the account is enrolled in the OVHcloud [IAM beta](https://labs.ovhcloud.com/en/iam/) 

## Example Usage

```hcl
data "ovh_iam_resource_groups" "my_groups" {
}
```

## Argument Reference

## Attributes Reference

* `id` - Hash of the list of the resource groups IDs.
* `policies` - List of the resource groups IDs.
