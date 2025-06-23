package main

import (
	"flag"
	"fmt"
	"github.com/aliakseiyanchuk/terraform-provider-az-confidential/tfgen"
	"os"
)

var baseParams = tfgen.ContentWrappingParams{
	NoLabels:              false,
	Labels:                "",
	TargetCoordinateLabel: false,
}

var outputCiphertextOnly = false

var baseFlags = flag.NewFlagSet("tfgen", flag.ExitOnError)

var subcommands []string

func init() {
	subcommands = []string{
		"secret",
		"password",
		"key",
		"certificate",
	}

	baseFlags.SetOutput(os.Stdout)
	baseFlags.StringVar(&baseParams.DestinationCoordinate.VaultName,
		"output-vault",
		"",
		"Output vault name")

	baseFlags.StringVar(&baseParams.DestinationCoordinate.Name,
		"output-vault-object",
		"",
		"Output vault object name")

	baseFlags.StringVar(&baseParams.WrappingKeyCoordinate.VaultName,
		"wrapping-key-vault",
		"",
		"Vault containing the wrapping key")

	baseFlags.StringVar(&baseParams.WrappingKeyCoordinate.KeyName,
		"wrapping-key-name",
		"",
		"Wrapping/encrypting key name")

	baseFlags.StringVar(&baseParams.WrappingKeyCoordinate.KeyVersion,
		"wrapping-key-version",
		"",
		"Wrapping/encrypting key version")

	baseFlags.StringVar(&baseParams.RSAPublicKeyFile,
		"pubkey",
		"",
		"RSA public key to encrypt secrets/content encryption keys",
	)

	baseFlags.BoolVar(&baseParams.NoLabels,
		"no-labels",
		true,
		"No not use any labels",
	)

	baseFlags.StringVar(&baseParams.Labels,
		"fixed-labels",
		"",
		"Fixed labels to associate with the ciphertext. Use comma to separate individual labels",
	)

	baseFlags.BoolVar(&baseParams.TargetCoordinateLabel,
		"target-only-label",
		true,
		"Label the ciphertext to be expandable only into specified vault and object",
	)

	baseFlags.BoolVar(&outputCiphertextOnly,
		"ciphertext-only",
		false,
		"Output only ciphertext (i.e. do not output associated Terraform code template)`",
	)
}

func main() {
	if parseErr := baseFlags.Parse(os.Args[1:]); parseErr != nil {
		_, _ = fmt.Printf("Invalid command line: %s", parseErr.Error())
		flag.PrintDefaults()
		os.Exit(1)
	}

	if len(baseFlags.Args()) == 0 {
		fmt.Println("Missing subcommand")
		printSubcommandSelectionHelp()
		os.Exit(1)
	}

	subCmd := baseFlags.Args()[0]
	var generator func(tfgen.ContentWrappingParams, bool, []string) (string, error)

	switch subCmd {
	case "secret":
		generator = tfgen.GenerateConfidentialSecretTerraformTemplate
	case "password":
		generator = tfgen.GenerateConfidentialPasswordTemplate
	case "key":
		generator = tfgen.GenerateConfidentialKeyTerraformTemplate
	case "certificate":
		generator = tfgen.GenerateConfidentialCertificateTerraformTemplate
	default:
		_, _ = fmt.Printf("Unknown subcommand: %s", subCmd)
		printSubcommandSelectionHelp()
		os.Exit(1)
	}

	if validationErr := baseParams.Validate(); validationErr != nil {
		_, _ = fmt.Printf("Incorrect basic arguments: %s", validationErr.Error())
		flag.PrintDefaults()
		os.Exit(1)
	}

	tfCode, err := generator(baseParams, !outputCiphertextOnly, baseFlags.Args()[1:])
	if err != nil {
		// Error message must be printed by the sub-command
		fmt.Println("Cannot produce template:")
		fmt.Println(err.Error())
		os.Exit(2)
	}

	fmt.Println(tfCode)

	// End of program
}

func printSubcommandSelectionHelp() {
	fmt.Println("Usage: tfgen [<standard options>] <subcommand> [<args>]")
	fmt.Println("Possible sub-commands are:")
	for _, cmd := range subcommands {
		fmt.Printf("- %s", cmd)
		fmt.Println()
	}
}
