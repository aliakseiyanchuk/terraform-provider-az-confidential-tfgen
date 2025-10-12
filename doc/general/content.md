# Encrypted content datasource

The [`az-confidential_general_content`](https://registry.terraform.io/providers/aliakseiyanchuk/az-confidential/latest/docs/data-sources/general_content) datasource is a
mechanism of providing secret strings into the Terraform state file. It can be seen as a functional equivalent of
[vault_generic_secret](https://registry.terraform.io/providers/hashicorp/vault/latest/docs/data-sources/generic_secret)
datasource or [azurerm_key_vault_secret](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/data-sources/key_vault_secret). 
Using this datasource, a similar functionality can be achieved without deploying Vault server.

An important operational characteristic of this data source compared to HashiCorp Vault-based or Azure Key Vault-based
implementations is that it is effectively a one-way write encryption that is accessible to anyone wishing to encrypt
an arbitrary data for the secure deployment. This makes this resource suitable for situations where:
- as a DevOps platform team, the team services a large number of product teams (which makes fine-grained access controls on Vault or Key Vault 
  possible but laborious)
- individual product team require deployments containing lots of specific parameters which are only determined by the
  product team.

An example is a product team requiring integration with multiple third-party APIs. Such team needs to manage a 
large set of varying credentials that periodically rotate, such as e.g API keys and secrets. 

Using this resource, such values can be directly injected into the Terraform state file, thus bypassing the need to 
have these stored in e.g. Vault or Key Vault. This resource decreases the churn of secrets in a vaults that otherwise
would be required.

## Synopsis

```shell
tfgen [common options] general content [datasource option]
```
The command encrypts supplied content and produces the Terraform code (or ciphertext only, if requested by 
the common option). The content is read from the standard input, if present, or from the file the option
`-content-file` specifies. If neither is found, an interactive input is used to prompt for the content to be encrypted.

## Datasource Options
The command accepts the following options:
- `-help` option prints the summary of the available options
- `-base64` input provided is a base-64 string that needs to be encoded
- `-content-file` read the content data from the specific content file. This option *must* be used if the content to 
  be provided is a multi-line string.

> Multi-line string can contain empty lines which could be significant for the program interpreting this. For CLI
> tool, an empty line means an "end-of-input" which could result in curtailed content being encrypted. For this reason,
> multi-line strings *must* be written into a file first and then encrypted using the `-content-file` option

## Examples

### Fully interactive content generation

In this example, the program will interactively prompt for content and for the public key.

```shell
tfgen general content      
Enter content data:
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
- The ciphertext will be marked "expired" after 120 days. After that, ciphertext will need to be created again. 
  This approach implements the requirements for periodic re-authentication
- The number of Terraform plan cycles is effectively limited to 500 times. This limits the enforcement of intended
  fair use of the protected content. After 500 runs are performed, the ciphertext will need to be re-created again.
- The provider unpacking this ciphertext must carry **either** `test`, `demo`, or `acc` constraint label. In practice,
  this measure will prevent the secret being unpacked e.g. into `prod` environment.

> Content datasource does not support create limiting as no resources are actually created from reading the 
> data source.

```shell
tfgen -pubkey <public-key-path> \
  -day-to-expire 120 -num-uses 500 -provider-constraints test,demo,acc \
  general content      
Enter content data:
> ... supply your content here by typing ...
```


### CLI-based generation

In this example, the program will produce an unfolded ciphertext from an environment variable. The public key
is read from the path the `<public-key-path>` variable specifies.

```shell
read CONTENT;
printf $CONTENT | tfgen -pubkey <public-key-path> -ciphertext-only -no-ciphertext-fold general content      
```
> Note: this example will also use default secondary protection settings; consider adding options that
> apply to your case. 

### Print command-line help
```shell
tfgen general content -help
```