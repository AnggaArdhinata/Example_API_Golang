package command

import (
	"restapiexample/src/configs/serve"
	"restapiexample/src/databases"

	"github.com/spf13/cobra"
)

var initCommand = &cobra.Command{
	Short: "example API using golang",
}


func init() {
	initCommand.AddCommand(serve.ServeComand)
	initCommand.AddCommand(databases.MigrateCmd)
}

func Run(args []string) error {
	initCommand.SetArgs(args)
	return initCommand.Execute()
}

	