# Key Vault Secret Resource

The [
`az-confidential_keyvault_key`](https://registry.terraform.io/providers/aliakseiyanchuk/az-confidential/latest/docs/resources/keyvault_key)
resource encrypts a key to be stored in the specified Azure Key Vault.

## Synopsis

```shell
tfgen [common options] kv key [resource option]
```

The command encrypts supplied key and produces the Terraform code (or ciphertext only, if requested by
the common option). The content is read from the standard input, if present, or from the file the option
`-key-file` specifies. If neither is found, an interactive input is used to prompt for the secret to be encrypted.

Where an input is already encrypted (such as e.g. encrypted RSA key), the tool will require supplying a valid
password that will be used to decrypt the supplied key material. This password can be provided via `-password-file`
option or entered via the interactive prompt.

## Datasource Options

The command accepts the following options:

- `-help` option prints the summary of the available options
- `-base64` input provided is a base-64 string; actual key value *shall be decoded* from base64.
- `-key-file` read the key material data from the specific file. 
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

In this example, the ciphertext created will be locked to the destination, that is, the secret will be only
deployable to the destination specified at the time the secret is created. Destination locking requires
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
    -destination-secret-name myTerrfaormKey    
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

teh
```shell
tfgen kv key -help
```