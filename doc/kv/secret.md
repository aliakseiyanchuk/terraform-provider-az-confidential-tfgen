# Key Vault Secret Resource

The [
`az-confidential_keyvault_secret`](https://registry.terraform.io/providers/aliakseiyanchuk/az-confidential/latest/docs/resources/keyvault_secret)
resource encrypts a secret to be stored in the specified Azure Key Vault.

## Synopsis

```shell
tfgen [common options] kv secret [resource option]
```

The command encrypts supplied secret and produces the Terraform code (or ciphertext only, if requested by
the common option). The content is read from the standard input, if present, or from the file the option
`-secret-file` specifies. If neither is found, an interactive input is used to prompt for the secret to be encrypted.

## Example Output
```terraform
resource "az-confidential_keyvault_secret" "secret" {
  content = <<-CIPHERTEXT
       H4sIAAAAAAAA/1TTtxKqzgKA8X6fwp65A0gSZ26hJEGygKEjrLikJawEn/4/53TnK3/9978/nTXDdHeK
       50aaG/0VoIwwJdBGLSLHHSuJLCPIgiACbe3RuP0hec9xIieIwMEFbEL4hiPscnjc5XiE9ERG1JX0zAL3
       28YTnI47lgF+k+awhR1RcDeRMUUdmY474I94RgUc/9Vo6+FxV8/0BPMREgC8nxTkTlirT7VZyujZP+Cy
       +tP3Yzszjn7OkGdt8HmXNBSr7KIURV7Ye6dMJl/7PSrWAmQK1NdU7jW/2N/0+3rdWGW8Pehz6KyHkK56
       vz4LivKCatY9dNgoav17NTaF3pqse98Y6EiuCCfrhvVcXoUauZEypCXjwyQ1Oid1rOuvPki9rLOyixnb
       +oqoMUu2JbYWqWqMwFzYTUZRXGe+1NI18ic3bSEyyjSI28P29a3xcXauS6J0sXv/SAeuP73rKz8azp7Q
       FwcCqk8pHru05HWdEracFwXxQztlJXmFwcPw0MVNMhJkZoclRLNp/ZwNK1IeJkMF7+ZwAMiTVMhortdk
       M91Uzal5R9SnvKh6jAXH9Z2ZqwcROcWJvVnmh2IDWCzGRLuZdrnzSgjoYCpjWwi7FM9X5E/bHD2kE9M3
       WDGH/n0fPlZc+Ye9XQcBP/TUdWDuZ7O8rNaC8V3CwJKo8aQeqrdOUGRMoRAmXtJS3Z7XLpPVefN6V7bx
       er79Bi099AEpRSM1lbxelvqJVwOIPN6s9MN0V/OGQ3GfQKJQedbOTbwkS4Y9XixZfdPm8KoHzEdtiZDc
       Y2mEVvDzuTQCt1N5eUmY1Es12pSQc/lbqqt9FK034mWt+A0k+RQJjUKLtRBvn2ltuMRmBo1tmNZ3L2CR
       4S+e+SpxxkQdBWqxm7OFEp4tNOF5Wl65bK3ik0eVXP4f/B1Kc9V/B/svAAD//wBP6cV5AwAA
       CIPHERTEXT
 
  # This secret is enabled for operation. Optionally, there is an option
  # to temporarily disable it.
  enabled = true

  # If the secret version should nto used before a specific date,
  # it needs to be formatted yyyy-mm-ddTHH:MM:SS'Z'
  # not_before_date = "2025-10-18T15:12:36Z"

  # If the secret version cannot be used after this date
  # it needs to be formatted yyyy-mm-dd'T'HH:MM:SS'Z'
  # not_after_date = "2026-10-18T15:12:36Z"
  
  tags = {
      # Fill the tags as desired
      # tagName = "TagValue"
    }

  destination_secret = {
      # Provider-configured destination vault will be used.
      # If this is not what you want, specify the destination vault explicitly
      # vault_name = "... the name of the vault where you want this secret ..."
      name = "destination-secret-name"
  }
}
```

## Resource Options

The command accepts the following options:

- `-help` option prints the summary of the available options
- `-base64` input provided is a base-64 string; actual secret value *shall be decoded* from base64.
  > Note: Key Vault secrets are text strings, not binary data. If your application requires a secret which is binary,
  > then it needs to be stored in the key vault in the Base-64 encoded form. Then this option *should not* be used
  > when creating a ciphertext.
- `-secret-file` read the content data from the specific content file. This option *must* be used if the content to
  be provided is a multi-line string.
- `-destination-vault` specifies the destination Azure Key Vault name where the resulting secret needs to be stored,
  > Note: Key Vault is a frequently used resource. The `az-confidential` provider requires the Terraform practitioner
  > to supply the `default_destination_vault_name`. If this option is omitted, then the provide-configured default is implied.
- `-destination-secret-name` specifies the secret name in the Azure Key Vault (that option `-destination-vault`
  specifies).
  This value must be explicitly supplied when creating Terraform code. The option can be omitted if
  `-ciphertext-only` common option is used without destination lock.

The table below summarises the scenarios which combinations of `-destination-secret-name` and `-destination-vault`
options are valid. ✅ indicates the *requirement* to specify an option.

| Option                     | Terraform output | Terraform output with `-lock-destination` | `-ciphertext-only` | `-ciphertext-only` with `-lock-destination' |
|----------------------------|------------------|-------------------------------------------|--------------------|---------------------------------------------|
| `-destination-secret-name` | ✅                | ✅                                         | *Has no effect*    | ✅                                           |
| `-destination-vault`       | *Optional*       | ✅                                         | *Has no effect*    | ✅                                           |

## Examples

### Fully interactive content generation

In this example, the program will interactively prompt for content and for the public key.

```shell
tfgen kv secret -destination-secret-name myTerrfaormSecret    
Enter secret data:
> ... supply your content here by typing ...
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

- The Azure Key Vault secret object must be created within 5 hours. If this doesn't happen, then a new ciphertext
  will need to be created.
- The ciphertext will be marked "expired" after 120 days. After that, ciphertext will need to be created again.
  This approach implements the requirements for periodic re-authentication. Any number of Terraform plans can
  run in this period.
- The Azure Key Vault secret object may be created maximum 5 times. This allows the deployer to destroy the
  infrastructure
  (e.g. to correct an error) and re-deploy it again **within** 5 hour window the `-time-to-create` option specifies.
  If the infrastructure is not built within 5 attempts, the ciphertext will need to be re-created again.
- The provider unpacking this ciphertext must carry **either** `test`, `demo`, or `acc` constraint label. In practice,
  this measure will prevent the secret being unpacked e.g. into `prod` environment.

```shell
tfgen -pubkey <public-key-path> \
  -time-to-create 5h \
  -day-to-expire 120 -num-uses 5 -provider-constraints test,demo,acc \
  kv secret      
Enter secret data:
> ... supply your content here by typing ...
```

### Destination Vault Locking

In this example, the ciphertext created will be locked to the destination, that is, the secret will be only
deployable to the destination specified at the time the ciphertext is created. Destination locking requires
`-destination-vault` to be supplied.

- Command that will ask secret content and public key:
  ```shell
  tfgen -lock-destination \
    kv secret \
  -destination-vault myTerraformVault \
  -destination-secret-name myTerrfaormSecret    
  ```
- Command that will ask to supply the secret content interactively and read the public key from supplied file:
  ```shell
  tfgen -lock-destination -pubkey <public-key-path>  \
    kv secret \
    -destination-vault myTerraformVault \ 
    -destination-secret-name myTerrfaormSecret    
  ```

### CLI-based generation

In this example, the program will produce an unfolded ciphertext from an environment variable. The public key
is read from the path the `<public-key-path>` variable specifies.

```shell
read CONTENT;
printf $CONTENT | tfgen -pubkey <public-key-path> \
                         -ciphertext-only -no-ciphertext-fold \
                         kv secret    
```

> Note: this example will also use default secondary protection settings; consider adding options that
> apply to your case.

### Print command-line help
The CLI arguments summary for this resource can be printed with:
```shell
tfgen kv secret -help
```