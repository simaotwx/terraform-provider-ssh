package main

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/loafoe/terraform-provider-ssh/ssh"
)

func main() {
	var debugMode bool

	flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := &plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return ssh.Provider()
		}}
	if debugMode {
		err := plugin.Debug(context.Background(), "registry.terraform.io/loafoe/ssh", opts)
		if err != nil {
			log.Fatal(err.Error())
		}
		return
	}
	plugin.Serve(opts)
}
