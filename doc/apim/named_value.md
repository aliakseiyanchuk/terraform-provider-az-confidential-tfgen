# API Management Named Value Resource

The [`az-confidential_apim_named_value`](https://registry.terraform.io/providers/aliakseiyanchuk/az-confidential/latest/docs/resources/apim_named_value)
resource encrypts a named value to be stored in the specified Azure API Management instance.

## Synopsis

```shell
tfgen [common options] apim named_value [resource option]
```

The command encrypts supplied named value and produces the Terraform code (or ciphertext only, if requested by
the common option). The content is read from the standard input, if present, or from the file the option
`-named-value-file` specifies. If neither is found, an interactive input is used to prompt for the certificate to be
encrypted.

## Example Output
```terraform
resource "az-confidential_apim_named_value" "named_value" {
   content = <<-CIPHERTEXT
            H4sIAAAAAAAA/1TTubKqSgCF4Zyn8AVuISItWnUDmWFDy6QiWQPN2MyzT3/q7Oys8Kta4f/f3wmyqsOD
            +IC+DP1focQBowmbRV1MtwNzAcyRu/IMQ8lbVwz7X7qeWBYAhqGsNsHExSkecBPj2yFuB0yP01A0Gb0w
            FJzr54jH24E5UjZBMa5xM4ltM04DKpppvB0oe2iXIsHDv+rvHb4dUFccatSg7PdHN6jGyWFBZMYU1SYP
            JLP3tEO9W81J/pN3b6ItYWqQEhxPbLZHwSilWxOOXNHhRMmBB52UW4NUlKRUoxSVy6tjbzHVd5XqODtd
            rw+hsOXMYOt+0UsxkaSTYtQSqqpBmQKZPbdJMaVYbLfLzAMqMrnwuCWSddKeoPRA/yRvTeS6ig/Ez8a+
            SJonlegHasibtKi4W+93nTgbgS4M+k+LKG7iHf5l9G7+wNcwiQpIes+qXkrzAxkuA4mh0eUXin5Ud+l6
            JvapezhkcYM6NGeyuFRg25WqmE9aYM31Hq4qDdOe9qxZtosoROQVvF0trVF+QboRiN6Vv0yzKTQ0gP6H
            SyHlGLDbSWYoNOuB7M7SvZuowN/KT75F3ZM+i0uVhE3jCrI67VaUh+DsaP13s3mv/OKSeqXxWxZscyfX
            sfVJIN7tzrU7dILQm7+7/66Rz1wNN2unYenBB6dvcdyDoFhfutAOFaVUsqMzQ6w6JILgKeRP8rJr9WQW
            BHZRSfO8Wl9K6xOKDhfc9+6JEkko732y8cDy443inCvIlod/ec9megoucAiB/A6Trv2sUyvU0vc8xs9H
            tGpuFBUs9yksR3ujPtGteBY+MWXa23o5K6peEAOzcf59mUxvjj0rhTsHNCtUoXDJezHOZHUH3uTelVDS
            x6odC6TIO0+dzny6IB+dq6j/VMd4rfyvfk4VcaXLUpukGXBeufKqxdz/p34Lk6H0b3F/AgAA//8ZTvZa
            igMAAA==
            CIPHERTEXT
  
    tags = [
      # Fill the tags as desired
      # "TagValue"
    ]

    # Confidential values created from the sensitive ciphertext are best kept secret, meaning that
    # these will be hidden on the Portal display
    secret = true

    # A display name of this named value. It doesn't play a role in the actual operation;
    # whereas it's great to give it a descriptive name. If it is not specified, it is
    # inferred from the name of the destination named value.
    # display_name = "confidential named value"

    destination_named_value = {
        # Specify a Azure subscription id where the APIM instance is created
        az_subscription_id = "...specify the subscription..."
        # Specify a Azure resource group  id where the APIM instance is created
        resource_group = "...specify the resource group name..."
        # Specify a Azure APIM service name
        api_management_name = "...specify the APIM service name..."
        # Specify the name this named value should use; may contain
        # only letters, digits, periods, dashes and underscores
        name = "...specify the APIM named value..."
    }
}
```

## Resource Options

The command accepts the following options:

- `-help` option prints the summary of the available options
- `-base64` input provided is a base-64 string; actual key value *shall be decoded* from base64.
- `-az-subscription-id` subscription Id where the target API Management instance is deployed
- `-resource-group-name` the Azure resource group name where the target API Management instance is deployed
- ` -service-name` the name of the target API Management instance with subscription (specified by `-az-subscription-id`)
  and resource group (specified by `-resource-group-name`).
- `-named-value` the name of the named value to be created within the API Management instance.
- `-named-value-file` read the named value from the specific file.

The table below summarises the scenarios which combinations of
options are valid. ✅ indicates the *requirement* to specify an option.

| Option                 | Terraform output | Terraform output with `-lock-destination` | `-ciphertext-only` | `-ciphertext-only` with `-lock-destination' |
|------------------------|------------------|-------------------------------------------|--------------------|---------------------------------------------|
| `-az-subscription-id`  | *Optional*       | ✅                                         | *Has no effect*    | ✅                                           |
| `-resource-group-name` | *Optional*       | ✅                                         | *Has no effect*    | ✅                                           |
| `-service-name`        | *Optional*       | ✅                                         | *Has no effect*    | ✅                                           |
| `-named-value`         | *Optional*       | ✅                                         | *Has no effect*    | ✅                                           |

## Examples

### Fully interactive content generation

In this example, the program will interactively prompt for content and for the public key. The command specified the intended
API Management named value name for added clarity of the generated Terraform configuration

```shell
tfgen apim named_value -named-value myTerrfaormNamedValue    
Enter named value data:
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

- The Azure API Management named value object must be created within 5 hours. If this doesn't happen, then a new ciphertext
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
  apim named_value      
Enter certificate data (hit Enter twice to end input):
> ... supply your content here by typing ...
```

### Destination API Management Service  Locking

In this example, the ciphertext created will be locked to the destination, that is, the named value will be only
deployable to the destination specified at the time the ciphertext is created. Destination locking requires
multiple options to be supplied.

- Command that will ask secret content and public key:
  ```shell
  tfgen -lock-destination \
    apim named_value \
  -az-subscription-id <subscription id> \
  -resource-group-name <resource group> \    
  -service-name <APIM service name> \    
  -namned-value <APIM service named value name>     
  ```
- Command that will ask to supply the key material interactively and read the public key from supplied file:
  ```shell
  tfgen -lock-destination -pubkey <public-key-path>  \
    apim named_value \
   -az-subscription-id <subscription id> \
   -resource-group-name <resource group> \    
   -service-name <APIM service name> \    
   -namned-value <APIM service named value name>
  ```

### CLI-based generation

In this example, the program will produce an unfolded ciphertext from an environment variable. The public key
is read from the path the `<public-key-path>` variable specifies.

```shell
read CONTENT;
printf $CONTENT | tfgen -pubkey <public-key-path> \
                         -ciphertext-only -no-ciphertext-fold \
                         apim named_value    
```

> Note: this example will also use default secondary protection settings; consider adding options that
> apply to your case.

### Print command-line help
The CLI arguments summary for this resource can be printed with:
```shell
tfgen apim named_value -help
```