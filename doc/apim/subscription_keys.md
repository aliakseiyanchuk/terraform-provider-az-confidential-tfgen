# API Management Named Value Resource

The [
`az-confidential_apim_subscription`](https://registry.terraform.io/providers/aliakseiyanchuk/az-confidential/latest/docs/resources/apim_subscription)
resource encrypts a subscription keys parameters that need to be created or changed in the Azure API Management
instance.

## Synopsis

```shell
tfgen [common options] apim subscription [resource option]
```

The command encrypts supplied named value and produces the Terraform code (or ciphertext only, if requested by
the common option). The content is read from the standard input, if present, or from the file the option
`-named-value-file` specifies. If neither is found, an interactive input is used to prompt for the certificate to be
encrypted.

## Resource Options

The command accepts the following options:

- `-help` option prints the summary of the available options
- `-subscription-id` id of the description to be created. Must be unique in the target API management instance.
  If none supplied, a UUID-based subscription Id will be generated.
- `-api` specifies the name API for which the subscription should be associated. Where no explicit API is specified,
  the subscription will be bound with all APIs. Mutually exclusive with `-product`.
- `-product` specifies the API Management product to be bound with this subscription. Mutually exclusive with `-api`
- `-owner` the identity that should be set as an owner identity.
- `-primary-key-file` specifies the file containing the primary subscription key
- `-secondary-key-file` specifies the file containing the secondary subscription key
- `-az-subscription-id` subscription Id where the target API Management instance is deployed
- `-resource-group-name` the Azure resource group name where the target API Management instance is deployed
- `-service-name` the name of the target API Management instance with subscription (specified by `-az-subscription-id`)
  and resource group (specified by `-resource-group-name`).

The table below summarises the scenarios which combinations of
options are valid. ✅ indicates the *requirement* to specify an option.

| Option                 | Terraform output | Terraform output with `-lock-destination` | `-ciphertext-only` | `-ciphertext-only` with `-lock-destination' |
|------------------------|------------------|-------------------------------------------|--------------------|---------------------------------------------|
| `-az-subscription-id`  | *Optional*       | ✅                                         | *Has no effect*    | ✅                                           |
| `-resource-group-name` | *Optional*       | ✅                                         | *Has no effect*    | ✅                                           |
| `-service-name`        | *Optional*       | ✅                                         | *Has no effect*    | ✅                                           |
| `-api` or `-product`   | *Optional*       | If omitted, implies all APIs              | *Has no effect*    | If omitted, implies all APIs                |
| `-subscription-id`     | *Optional*       | If omitted, will be generated             | *Has no effect*    | If omitted, will be generated               |
| `-owner`               | *Optional*       | If omitted, implies no owner is set.      | *Has no effect*    | If omitted, implies no owner is set.        |

## Examples

### Fully interactive content generation

In this example, the program will interactively prompt for content and for the public key. The command specified the
intended
API Management named value name for added clarity of the generated Terraform configuration

```shell
tfgen apim subscription -subscription-id myTerrfaormApimSubscriptionId    
Enter primary subscription key:
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

- The Azure API Management named value object must be created within 5 hours. If this doesn't happen, then a new
  ciphertext
  will need to be created.
- The ciphertext will be marked "expired" after 120 days. After that, ciphertext will need to be created again.
  This approach implements the requirements for periodic re-authentication. Any number of Terraform plans can
  run in this period.
- The Azure API Management named value object may be created maximum 5 times. This allows the deployer to destroy the
  infrastructure
  (e.g. to correct an error) and re-deploy it again **within** 5 hour window the `-time-to-create` option specifies.
  If the infrastructure is not built within 5 attempts, the ciphertext will need to be re-created again.
- The provider unpacking this ciphertext must carry **either** `test`, `demo`, or `acc` constraint label. In practice,
  this measure will prevent the secret being unpacked e.g. into `prod` environment.

```shell
tfgen -pubkey <public-key-path> \
  -time-to-create 5h \
  -day-to-expire 120 -num-uses 5 -provider-constraints test,demo,acc \
  apim subscription      
Enter primary subscription key:
> ... supply your certificate here by typing ...
Enter secondary subscription key:
> ... supply your content here by typing ...
```

### Destination API Management Service Locking

In this example, the ciphertext created will be locked to the destination, that is, the named value will be only
deployable to the destination specified at the time the ciphertext is created. Destination locking requires
multiple options to be supplied.

- Command that will ask secret content and public key:
  ```shell
  tfgen -lock-destination \
    apim subscription \
  -az-subscription-id <subscription id> \
  -resource-group-name <resource group> \    
  -service-name <APIM service name> \    
  -namned-value <APIM service named value name>     
  ```
- Command that will ask to supply the key material interactively and read the public key from supplied file:
  ```shell
  tfgen -lock-destination -pubkey <public-key-path>  \
    apim subscription \
   -az-subscription-id <subscription id> \
   -resource-group-name <resource group> \    
   -service-name <APIM service name> \    
   -namned-value <APIM service named value name>
  ```

### CLI-based generation

In this example, the program will produce an unfolded ciphertext from an environment variable. The public key
is read from the path the `<public-key-path>` variable specifies. Subscription requires primary and secondary keys
which **Must** be different.

```shell
read PRIMARY_KEY; read SECONDARY_KEY; 
printf "%s\n%s\n" "$PRIMARY_KEY" "$SECONDARY_KEY"| tfgen \
  -pubkey <public-key-path> \
  -ciphertext-only -no-ciphertext-fold \
  apim subscription    
```

> Note: this example will also use default secondary protection settings; consider adding options that
> apply to your case.

### Print command-line help
The CLI arguments summary for this resource can be printed with:
```shell
tfgen apim subscription -help
```