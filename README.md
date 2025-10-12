# `az-confidential` Terraform Code Generator Tool

`tfgen` is a complimentary tool for [`az-confidendtial` Terraform provider](https://registry.terraform.io/providers/aliakseiyanchuk/az-confidential/latest).
The tool encrypts the provided confidential material and generates Terraform code
that can readily be deployed by the provider.

> This project is a distribution assembly-only project that publishes the 
> `tfgen` tool separately from the Terraform provider distributions.

## General Syntax

The command line syntax is as follows:
`tfgen [common options] [group] [resource] [resource options]`

Common options can be divided into three categories: key-encryption-key related options, secondary ciphertext protection,
and CLI integration options.

> Where no specific secondary protection options are specified, the tool will assume the following defaults:
> - expiry time in 365 days
> - number of uses: 10
>
> Provider and destination Azure object will remain unlocked. Consider specifying secondary ciphertext proection
> parameters explicitly when the ciphertext is being created.

A build-in help can be obtained using the `-help` option. The option can be supplied to the tool itself, to groups, 
and to individual resource, e.g.:
```shell
tfgen -help # Print common options, list available groups
tfgen kv -help # List resources in the group
tfgen kv secret -help # Print options this resource supports
```

### Key-Encrypting Key related options

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

### Secondary ciphertext protection

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

### CLI integration options

- `-ciphertext-only` instructs to output only ciphertext
- `-no-ciphertext-fold` instructs to output a ciphertext as a single string, not as a folded string. This option could
  is intended for cases where the ciphertext author needs to produce a text string which is later on interpreted by
  automation tools.

> Note: Ciphertext by default is a multi-line string that is folded at 80 characters per row for the readability
> purposes. Folded ciphertext can be "unfolded" by simply removing all new lines (and any whitespace).

## Sub-commands:
- `password`: generates a password that **will be** in the state file. This datasource
  is added to support the feature of providing very long passwords (i.e. those that cannot
  be encrypted by an RSA key alone)
- `secret`: encrypt a secret
- `key`: encrypt an RSA, EC, or a symmetric key
- `certificate` encrypt a certificate

## Encrypting Password

The command synopsis is:
`tfgen [common options] password [password options]`

Command options are:
- `-password-file`: read the password for a specified file
- `-base64`: the file (or input) being read is base-64 encoded

Where the password file would not be specified, the program will attempt to read input from the standard input.
If not available, interactive prompt will be used to prompt to enter the password information.

Example to label the password to with provides configured to accept the `demo` and `testing` labels:
```shell
tfgen -pubkey PATH_TO_PUB_KEY -fixed-labels demo,testing password
```

A successful execution of this command should produce a template that looks as follows:
```terraform
#
# Minimal terraform code for your example:
# ----------------------------------------------------------------------------

data "az-confidential_password" "confidential_password" {
  content = "....a very long, base-64 encoded ciphertext ciphertext..."
}
```
## Encrypt secret

The command synopsis is:
`tfgen [common options] secret [secret options]`

Command options are:
- `-secret-file`: read the secret for a specified file
- `-base64`: the file (or input) being read is base-64 encoded

Where the secret would not be specified, the program will attempt to read input from the standard input.
If not available, interactive prompt will be used to prompt to enter the password information.

Example to label the secret to be used with a provider configured to accept the `demo` and `testing` labels:
```shell
tfgen -pubkey PATH_TO_PUB_KEY \
  -output-vault demo-vault -output-vault-object demo \
  -fixed-labels demo,testing \
  secret
```
A successful execution of this command should produce a template that looks as follows:
```terraform
#
# Minimal terraform code for your example:
# ----------------------------------------------------------------------------

resource "az-confidential_secret" "secret" {
  content = "...a very long, base-64 encoded ciphertext ciphertext..."

  # This secret is enabled for operation. Optionally, there is an option
  # to temporarily disable it.
  enabled = true

  # The secret version cannot be used before this date
  # Needs to be formatted yyyy-mm-ddTHH:MM:SS'Z'
  # not_before_date = "2025-06-21T18:04:00Z"

  # The secret version cannot be used after this date
  # Needs to be formatted yyyy-mm-dd'T'HH:MM:SS'Z'
  # not_after_date = "2026-06-21T18:04:00Z"

  tags = {
    # Fill the tags as desired
    # tagName =  "TagValue"
  }

  destination_secret = {
    vault_name = "demo-vault"
    name = "demo"
  }
}
```
> You may need to modify this template as required for your project.

## Encrypt a key

The command synopsis is:
`tfgen [common options] key [secret options]`

Command options are:
- `-key-file`: read the secret for a specified file
- `-password-file`: read the password for the encrypted private key from the specified file
- `-base64`: the file (or input) being read is base-64 encoded
- `-symmetric` The key being created is a symmetric key

Where the key would not be specified, the program will attempt to read input from the standard input.
If not available, interactive prompt will be used to prompt to enter the password information.
The same logic applied for passwords.

Example to label the key to be used with a provider configured to accept the `demo` and `testing` labels:
```shell
tfgen -pubkey PATH_TO_PUB_KEY \
  -output-vault demo-vault -output-vault-object demo \
  -fixed-labels demo,testing \
  key
```
A successful execution of this command should produce a template that looks as follows:
```terraform
#
# Minimal terraform code for your example:
# ----------------------------------------------------------------------------

resource "az-confidential_key" "key" {
  content = "...a very long, base-64 encoded ciphertext ciphertext..."

  # This secret is enabled for operation. Optionally, there is an option
  # to temporarily disable it.
  enabled = true

  key_opts = toset([
    "decrypt",
    "encrypt",
    "import",
    "sign",
    "unwrapKey",
    "verify",
    "wrapKey"
  ])

  # The secret version cannot be used before this date
  # Needs to be formatted yyyy-mm-ddTHH:MM:SS'Z'
  # not_before_date = "2025-06-21T18:10:21Z"

  # The secret version cannot be used after this date
  # Needs to be formatted yyyy-mm-dd'T'HH:MM:SS'Z'
  # not_after_date = "2026-06-21T18:10:21Z"

  tags = {
        # Fill the tags as desired
        # tagName =  "TagValue"
      }

  destination_key = {
        vault_name = "demo-vault"
        name = "demo"
  }
}
```
> You may need to tweak the parameters of the allowed key operations as required to your use case.

## Encrypt a certificate

The command synopsis is:
`tfgen [common options] certificate [secret options]`

Command options are:
- `-cert-file`: read the secret for a specified file
- `-password-file`: read the password for the encrypted private key from the specified file
- `-base64`: the file (or input) being read is base-64 encoded
- `-no-der-verify` Skip verification of DER-encoded certificate

Where the key would not be specified, the program will attempt to read input from the standard input.
If not available, interactive prompt will be used to prompt to enter the password information.
The same logic applied for passwords.

Example to label the certificate to be used with a provider configured to accept the `demo` and `testing` labels:
```shell
tfgen -pubkey PATH_TO_PUB_KEY \
  -output-vault demo-vault -output-vault-object demo \
  -fixed-labels demo,testing \
  key
```
A successful execution of this command should produce a template that looks as follows:
```terraform
#
# Minimal terraform code for your example:
# ----------------------------------------------------------------------------

resource "az-confidential_key" "key" {
  content = "...a very long, base-64 encoded ciphertext ciphertext..."

  # This secret is enabled for operation. Optionally, there is an option
  # to temporarily disable it.
  enabled = true

  key_opts = toset([
    "decrypt",
    "encrypt",
    "import",
    "sign",
    "unwrapKey",
    "verify",
    "wrapKey"
  ])

  # The secret version cannot be used before this date
  # Needs to be formatted yyyy-mm-ddTHH:MM:SS'Z'
  # not_before_date = "2025-06-21T18:10:21Z"

  # The secret version cannot be used after this date
  # Needs to be formatted yyyy-mm-dd'T'HH:MM:SS'Z'
  # not_after_date = "2026-06-21T18:10:21Z"

  tags = {
        # Fill the tags as desired
        # tagName =  "TagValue"
      }

  destination_key = {
        vault_name = "demo-vault"
        name = "demo"
  }
}
```
> You may need to tweak the parameters of the allowed key operations as required to your use case.