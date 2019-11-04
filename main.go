package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/terraform-providers/terraform-provider-aws/aws"
	"log"
)

func main() {
	log.SetFlags(log.Llongfile)
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: aws.Provider})
}
