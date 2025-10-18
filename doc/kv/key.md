# Key Vault Key Resource

The [
`az-confidential_keyvault_key`](https://registry.terraform.io/providers/aliakseiyanchuk/az-confidential/latest/docs/resources/keyvault_key)
resource encrypts a key to be stored in the specified Azure Key Vault.

## Synopsis

```shell
tfgen [common options] kv key [resource option]
```

The command encrypts supplied key and produces the Terraform code (or ciphertext only, if requested by
the common option). The content is read from the standard input, if present, or from the file the option
`-key-file` specifies. If neither is found, an interactive input is used to prompt for the key to be encrypted.

Where an input is already encrypted (such as e.g. encrypted RSA key), the tool will require supplying a valid
password that will be used to decrypt the supplied key material. This password can be provided via `-password-file`
option or entered via the interactive prompt.

## Example Output
```terraform
resource "az-confidential_keyvault_key" "key" {
  content = <<-CIPHERTEXT
         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
         CIPHERTEXT

  # This key is enabled for operation. Optionally, there is an option
  # to temporarily disable it.
  enabled = true
  key_opts = toset([
    "encrypt",
    "decrypt",
    "sign",
    "verify",
    "wrapKey",
    "unwrapKey",
  ])

  # The key version cannot be used before this date
  # Needs to be formatted yyyy-mm-ddTHH:MM:SS'Z'
  # not_before_date = "2025-10-18T15:14:36Z"

  # The key version cannot be used after this date
  # Needs to be formatted yyyy-mm-dd'T'HH:MM:SS'Z'
  # not_after_date = "2026-10-18T15:14:36Z"

  tags = {
        # Fill the tags as desired
        # tagName =  "TagValue"
      }

  destination_key = {
        # Provider-configured destination vault will be used.
        # If this is not what you want, specify the destination vault explicitly
        # vault_name = "... the name of the vault where you want this key ..."
        name = "destination-key-name"
  }
}
```

## Resource Options

The command accepts the following options:

- `-help` option prints the summary of the available options
- `-base64` input provided is a base-64 string; actual key value *shall be decoded* from base64.
- `-key-file` read the key material data from the specific file. 
- `-password-file` read the key material password from the supplied file 
- `-destination-vault` specifies the destination Azure Key Vault name where the resulting key needs to be stored
  > Note: Key Vault is a frequently used resource. The `az-confidential` provider requires the Terraform practitioner
  > to supply the `default_destination_vault_name`. If this option is omitted, then the provide-configured default is implied.
- `-destination-key-name` specifies the key name in the Azure Key Vault (that option `-destination-vault`
  specifies).
  This value must be explicitly supplied when creating Terraform code. The option can be omitted if
  `-ciphertext-only` common option is used without destination lock.

The table below summarises the scenarios which combinations of `-destination-key-name` and `-destination-vault`
options are valid. ✅ indicates the *requirement* to specify an option.

| Option                  | Terraform output | Terraform output with `-lock-destination` | `-ciphertext-only` | `-ciphertext-only` with `-lock-destination' |
|-------------------------|------------------|-------------------------------------------|--------------------|---------------------------------------------|
| `-destination-key-name` | ✅                | ✅                                         | *Has no effect*    | ✅                                           |
| `-destination-vault`    | *Optional*       | ✅                                         | *Has no effect*    | ✅                                           |

## Examples

### Fully interactive content generation

In this example, the program will interactively prompt for content and for the public key.

```shell
tfgen kv key -destination-key-name myTerrfaormKey    
Enter key data (hit Enter twice to end input):
> ... supply your key here by typing ...
Please provide public key of the key wrapping key:
>-----BEGIN PUBLIC KEY-----
... public key contgent .....
-----END PUBLIC KEY-----
```

> Note: the above-mentioned example uses default secondary ciphertext protection parameters. Consider overriding
> these as fit to your case using secondary protection parameters options.

### Customized secondary protection parameters

In this example, the program will interactively prompt for content and apply the following secondary protection
parameters:

- The Azure Key Vault key object must be created within 5 hours. If this doesn't happen, then a new ciphertext
  will need to be created.
- The ciphertext will be marked "expired" after 120 days. After that, ciphertext will need to be created again.
  This approach implements the requirements for periodic re-authentication. Any number of Terraform plans can
  run in this period.
- The Azure Key Vault key object may be created maximum 5 times. This allows the deployer to destroy the
  infrastructure
  (e.g. to correct an error) and re-deploy it again **within** 5 hour window the `-time-to-create` option specifies.
  If the infrastructure is not built within 5 attempts, the ciphertext will need to be re-created again.
- The provider unpacking this ciphertext must carry **either** `test`, `demo`, or `acc` constraint label. In practice,
  this measure will prevent the secret being unpacked e.g. into `prod` environment.

```shell
tfgen -pubkey <public-key-path> \
  -time-to-create 5h \
  -day-to-expire 120 -num-uses 5 -provider-constraints test,demo,acc \
  kv key      
Enter key data (hit Enter twice to end input):
> ... supply your content here by typing ...
```

### Destination Vault Locking

In this example, the ciphertext created will be locked to the destination, that is, the key will be only
deployable to the destination specified at the time the ciphertext is created. Destination locking requires
`-destination-vault` to be supplied.

- Command that will ask secret content and public key:
  ```shell
  tfgen -lock-destination \
    kv key \
  -destination-vault myTerraformVault \
  -destination-key-name myTerrfaormKey    
  ```
- Command that will ask to supply the key material interactively and read the public key from supplied file:
  ```shell
  tfgen -lock-destination -pubkey <public-key-path>  \
    kv key \
    -destination-vault myTerraformVault \ 
    -destination-key-name myTerrfaormKey    
  ```

### CLI-based generation

In this example, the program will produce an unfolded ciphertext from an environment variable. The public key
is read from the path the `<public-key-path>` variable specifies.

```shell
read CONTENT;
printf $CONTENT | tfgen -pubkey <public-key-path> \
                         -ciphertext-only -no-ciphertext-fold \
                         kv key    
```

> Note: this example will also use default secondary protection settings; consider adding options that
> apply to your case.

### Print command-line help
The CLI arguments summary for this resource can be printed with:
```shell
tfgen kv key -help
```