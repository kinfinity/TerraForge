package templates

import "github.com/spf13/cobra"

type CommandGroup struct {
	Message  string
	Commands []*cobra.Command
}

type CommandGroups []CommandGroup

//Adds command to a
func (cg CommandGroups) Add(command *cobra.Command) {
	for _, g := range cg { // adding to every command group
		command.AddCommand(g.Commands...) //
	}
}

// Checks if CommandGroups has a command Nested
func (cg CommandGroups) Has(command *cobra.Command) bool {
	for _, g := range cg {
		for _, c := range g.Commands {
			if c == command {
				return true
			}
		}
	}
	return false
}

// Create CommandGroup and fills with commands into a CommandGroups
func ExtendCommands(cg CommandGroups, message string, commands []*cobra.Command) CommandGroups {
	g := CommandGroup{Message: message}

	// Iterate through commands and Check if they are Nested in CommandGroups
	for _, c := range commands {
		if !cg.Has(c) && len(c.Short) != 0 {
			//
			g.Commands = append(g.Commands, c)
		}
	}
	if len(g.Commands) == 0 {
		return cg // We did not fill any new commands
	}
	// Add CommandGroup to CommandGroups
	return append(cg, g)
}
