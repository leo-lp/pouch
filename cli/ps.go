package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

// psDescription is used to describe ps command in detail and auto generate command doc.
// TODO: add description
var psDescription = ""

// PsCommand is used to implement 'ps' command.
type PsCommand struct {
	baseCommand
}

// Init initializes PsCommand command.
func (p *PsCommand) Init(c *Cli) {
	p.cli = c
	p.cmd = &cobra.Command{
		Use:   "ps",
		Short: "List all containers",
		Long:  psDescription,
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return p.runPs(args)
		},
		Example: psExample(),
	}
	p.addFlags()
}

// addFlags adds flags for specific command.
func (p *PsCommand) addFlags() {
	// TODO: add flags here
}

// runPs is the entry of PsCommand command.
func (p *PsCommand) runPs(args []string) error {
	apiClient := p.cli.Client()

	containers, err := apiClient.ContainerList()
	if err != nil {
		return fmt.Errorf("failed to get container list: %v", err)
	}

	display := p.cli.NewTableDisplay()
	display.AddRow([]string{"Name", "ID", "Status", "Image"})
	for _, c := range containers {
		display.AddRow([]string{c.Names[0], c.ID[:6], c.Status, c.Image})
	}
	display.Flush()
	return nil
}

// psExample shows examples in ps command, and is used in auto-generated cli docs.
// TODO: add example
func psExample() string {
	return ""
}
