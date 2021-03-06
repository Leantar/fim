package cmd

import (
	"fmt"
	"git.leantar.de/torge/fimclient/client"
	"git.leantar.de/torge/fimclient/modules/config"
	"github.com/spf13/cobra"
)

func approveBaselineUpdate(_ *cobra.Command, _ []string) error {
	var conf client.Config
	err := config.FromYamlFile(confPath, &conf)
	if err != nil {
		return fmt.Errorf("failed to read config: %w", err)
	}

	c, err := client.NewConnectedClient(conf)
	if err != nil {
		return err
	}

	return c.CreateBaselineUpdateApproval(agentName)
}
