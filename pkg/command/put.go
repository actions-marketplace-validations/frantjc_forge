package command

import "github.com/spf13/cobra"

func NewPut() *cobra.Command {
	var (
		params  = map[string]string{}
		version = map[string]string{}
		cmd     = &cobra.Command{
			Use:  "put",
			Args: cobra.ExactArgs(1),
			Run: func(cmd *cobra.Command, args []string) {
				if err := processResource(cmd.Context(), cmd.Use, args[0], params, version); err != nil {
					cmd.PrintErrln(err)
					return
				}
			},
		}
	)

	cmd.Flags().StringToStringVarP(&params, "param", "p", make(map[string]string), "params")
	cmd.Flags().StringToStringVarP(&version, "version", "w", make(map[string]string), "version")

	return cmd
}
