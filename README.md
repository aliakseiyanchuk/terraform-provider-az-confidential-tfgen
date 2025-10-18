# `az-confidential` Terraform Code Generator Tool

`tfgen` is a complimentary tool for [`az-confidendtial` Terraform provider](https://registry.terraform.io/providers/aliakseiyanchuk/az-confidential/latest).
The tool encrypts the (interactively) provided confidential material and generates Terraform code
that can readily be deployed by the `az-confidential` provider. Depending on the resource, the Terraform output
may contain optional parameters (such as e.g. Azure portal display names) which are not essential for the encryption
process, where the Terraform practitioner may wish nevertheless to adjust the generated Terraform code.

## Command Syntax

The command line syntax is as follows:
`tfgen [common options] [group] [resource] [resource options]`

### Command Options

A build-in help can be obtained using the `-help` option. The option can be supplied to the tool itself, to groups, 
and to individual resource, e.g.:
```shell
tfgen -help # Print common options, list available groups
tfgen kv -help # List resources in the group
tfgen kv secret -help # Print options this resource supports
```

#### Key-Encrypting Key related options

The Key-Encrypting Key (KEK) is the "master" RSA key underpinning the encryption procedure used in creating the
ciphertext. 

- `-pubkey` public key of the KEK. If not supplied explicitly, the tool will prompt for an interactive input
- `-wrapping-key-vault` the specific vault name containing the KEK that needs to be included in Terraform code. The option
  can be omitted; in that case, the provider defaults are used.
- `-wrapping-key-name` the name of KEK in the vault. The option
  can be omitted; in that case, the provider defaults are used.
- `-wrapping-key-version` the version of KEK used. The option
  can be omitted; in that case, the latest version of hte key would be used.

For long-term maintainability of the Terraform code, it is recommended to include `wrapping-key-vault`,  `wrapping-key-name`,
and  `wrapping-key-version` parameters always in the CLI arguments to ensure that the generated ciphertext is "pinned"
with the specific KEK version. This simplifies procedures around rotating the KEK and re-encrypting assets in the 
Terraform code periodically.

#### Secondary ciphertext protection

Secondary ciphertext protection options embed instructions into the ciphertext to limit the ciphertext use by the
`az-confidendtial` provider. (See [provider configuration](https://github.com/aliakseiyanchuk/terraform-provider-az-confidential/blob/main/docs/index.md)
for the description of primary and secondary protection measures.)

- `-provider-constraints` a comma-separated list of the labels that need to be associated with the `az-confidential` provider
  instance in order to perform the ciphertext decryption (as a pre-requisite to the actual resource deployment)
- `-lock-destination` adds a "lock" of the target Azure resource (e.g. resource group, name, URL, etc.) into which the
  confidential material encrypted in this ciphertext may be unpacked.
- `time-to-create` sets a specific time-frame for the resource to be created. If omitted, then
  a constraint of 3 calendar days will be added automatically. A constraint can be removed by specifying `-no-create-limit` option.
- `days-to-expire` sets a specific number of days before the ciphertext expires completely. Expired ciphertext cannot be used;  
  the owner of the confidential material has to re-encrypt it. If omitted, then
  a constraint of 365 calendar days will be added automatically. A constraint can be removed by specifying `-no-expiry-limit` option.
- `num-uses` sets a specific number of times this ciphertext can be used to create a resource. Depleted ciphertext cannot be used;  
  the owner of the confidential material has to re-encrypt it. If omitted, then
  a constraint of 10 uses will be added automatically. Option `-create-once` can be specified to allow the resource to be created
  only one time, which is shortcut for `-num-uses 1` A constraint can be removed by specifying `-no-usage-limit` option.

#### CLI integration options

- `-ciphertext-only` instructs to output only ciphertext
- `-no-ciphertext-fold` instructs to output a ciphertext as a single string, not as a folded string. This option could
  is intended for cases where the ciphertext author needs to produce a text string which is later on interpreted by
  automation tools.

> Note: Ciphertext by default is a multi-line string that is folded at 80 characters per row for the readability
> purposes. Folded ciphertext can be "unfolded" by simply removing all new lines (and any whitespace).

### Groups and resources
The tool supports the following groups and Azure resources for which the ciphertext is generated:
- [general](./doc/general/index.md)
    - [content](./doc/general/content.md) creates encrypted string values injectable into Terraform state
- [kv](./doc/kv/index.md) for Azure Key Vault resources
    - [secret](./doc/kv/secret.md) Key Vault secret
    - [key](./doc/kv/key.md) Key Vault key
    - [certificate](./doc/kv/certificate.md) Key Vault certificate
- [apim](./doc/apim/index.md)
    - [named value](./doc/apim/named_value.md) encrypts a value of a sensitive API Management named value
    - [subscription keys](./doc/apim/subscription_keys.md) encrypts the subscription keys for a API Management subscription.

Common options can be divided into three categories: key-encryption-key related options, secondary ciphertext protection,
and CLI integration options.

> Where no specific secondary protection options are specified, the tool will assume the following defaults:
> - expiry time in 365 days
> - number of uses: 10
>
> Provider and destination Azure object will remain unlocked. Consider specifying secondary ciphertext proection
> parameters explicitly when the ciphertext is being created.

## Reporting bugs

Please raise a new issue in [this GitHub project](https://github.com/aliakseiyanchuk/terraform-provider-az-confidential-tfgen/issues)
describing the nature of the project as fully as practical.
