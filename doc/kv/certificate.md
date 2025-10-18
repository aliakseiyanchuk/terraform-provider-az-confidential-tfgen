# Key Vault Certificate Resource

The [
`az-confidential_keyvault_certificate`](https://registry.terraform.io/providers/aliakseiyanchuk/az-confidential/latest/docs/resources/keyvault_certificate)
resource encrypts a certificate to be stored in the specified Azure Key Vault.

## Synopsis

```shell
tfgen [common options] kv cerificate [resource option]
```

The command encrypts supplied key and produces the Terraform code (or ciphertext only, if requested by
the common option). The content is read from the standard input, if present, or from the file the option
`-cert-file` specifies. If neither is found, an interactive input is used to prompt for the certificate to be encrypted.

Where an input is already encrypted (such as e.g. PKCS12 files), the tool will require supplying a valid
password that will be used to decrypt the supplied key material. This password can be provided via `-password-file`
option or entered via the interactive prompt.

## Example Output
```terraform
resource "az-confidential_keyvault_certificate" "cert" {
  content = <<-CIPHERTEXT
         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
         /2gmcrwbFAAA
         CIPHERTEXT

  # This certificate is enabled for operation. Optionally, there is an option
  # to temporarily disable it.
  enabled = true

  # The certificate version cannot be used before this date
  # Needs to be formatted yyyy-mm-ddTHH:MM:SS'Z'
  # not_before_date = "2025-10-18T15:15:36Z"

  # The certificate version cannot be used after this date
  # Needs to be formatted yyyy-mm-dd'T'HH:MM:SS'Z'
  # not_after_date = "2026-10-18T15:15:36Z"

    tags = {
      # Fill the tags as desired
      # tagName =  "TagValue"
    }

    destination_certificate = {
        # Provider-configured destination vault will be used.
        # If this is not what you want, specify the destination vault explicitly
        # vault_name = "... the name of the vault where you want this certificate ..."
        name = "destination-cert-name"
    }
}
```
## Resource Options

The command accepts the following options:

- `-help` option prints the summary of the available options
- `-base64` input provided is a base-64 string; actual key value *shall be decoded* from base64.
- `-cert-file` read the key material data from the specific file.
- `-password-file` read the certificate material password from the supplied file
- `-no-der-verify` skip checking the correctness of key material/password combination
- `-destination-vault` specifies the destination Azure Key Vault name where the resulting certificate needs to be stored
  > Note: Key Vault is a frequently used resource. The `az-confidential` provider requires the Terraform practitioner
  > to supply the `default_destination_vault_name`. If this option is omitted, then the provide-configured default is implied.
- `-destination-cert-name` specifies the certificate name in the Azure Key Vault (that option `-destination-vault`
  specifies).
  This value must be explicitly supplied when creating Terraform code. The option can be omitted if
  `-ciphertext-only` common option is used without destination lock.

The table below summarises the scenarios which combinations of `-destination-cert-name` and `-destination-vault`
options are valid. ✅ indicates the *requirement* to specify an option.

| Option                   | Terraform output | Terraform output with `-lock-destination` | `-ciphertext-only` | `-ciphertext-only` with `-lock-destination' |
|--------------------------|------------------|-------------------------------------------|--------------------|---------------------------------------------|
| `-destination-cert-name` | ✅                | ✅                                         | *Has no effect*    | ✅                                           |
| `-destination-vault`     | *Optional*       | ✅                                         | *Has no effect*    | ✅                                           |

## Examples

### Fully interactive content generation

In this example, the program will interactively prompt for content and for the public key.

```shell
tfgen kv certificate -destination-key-name myTerrfaormCertiticate    
Enter certificate data (hit Enter twice to end input):
> ... supply your certificate here by typing ...
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

- The Azure Key Vault certificate object must be created within 5 hours. If this doesn't happen, then a new ciphertext
  will need to be created.
- The ciphertext will be marked "expired" after 120 days. After that, ciphertext will need to be created again.
  This approach implements the requirements for periodic re-authentication. Any number of Terraform plans can
  run in this period.
- The Azure Key Vault certificate object may be created maximum 5 times. This allows the deployer to destroy the
  infrastructure
  (e.g. to correct an error) and re-deploy it again **within** 5 hour window the `-time-to-create` option specifies.
  If the infrastructure is not built within 5 attempts, the ciphertext will need to be re-created again.
- The provider unpacking this ciphertext must carry **either** `test`, `demo`, or `acc` constraint label. In practice,
  this measure will prevent the secret being unpacked e.g. into `prod` environment.

```shell
tfgen -pubkey <public-key-path> \
  -time-to-create 5h \
  -day-to-expire 120 -num-uses 5 -provider-constraints test,demo,acc \
  kv certificate      
Enter certificate data (hit Enter twice to end input):
> ... supply your content here by typing ...
```

### Destination Vault Locking

In this example, the ciphertext created will be locked to the destination, that is, the certificate will be only
deployable to the destination specified at the time the ciphertext is created. Destination locking requires
`-destination-vault` to be supplied.

- Command that will ask secret content and public key:
  ```shell
  tfgen -lock-destination \
    kv certificate \
  -destination-vault myTerraformVault \
  -destination-cert-name myTerrfaormCert    
  ```
- Command that will ask to supply the key material interactively and read the public key from supplied file:
  ```shell
  tfgen -lock-destination -pubkey <public-key-path>  \
    kv certificate \
    -destination-vault myTerraformVault \ 
    -destination-cert-name myTerrfaormCert   
  ```

### CLI-based generation

In this example, the program will produce an unfolded ciphertext from an environment variable. The public key
is read from the path the `<public-key-path>` variable specifies.

```shell
read CONTENT;
printf $CONTENT | tfgen -pubkey <public-key-path> \
                         -ciphertext-only -no-ciphertext-fold \
                         kv certificate    
```

> Note: this example will also use default secondary protection settings; consider adding options that
> apply to your case.

### Print command-line help
The CLI arguments summary for this resource can be printed with:
```shell
tfgen kv certificate -help
```