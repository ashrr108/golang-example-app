package jwt

import (
    "github.com/golang-jwt/jwt"
    "github.com/spf13/cast"
    "github.com/spf13/cobra"
)

var Cmd = &cobra.Command{}

func Init() {
    Cmd.AddCommand(tokenCmd)
}

var tokenCmd = &cobra.Command{
    Use:   "token",
    Short: "Generate a token",
    Long:  `This subcommand generates a token`,
    Run: func(cmd *cobra.Command, args []string) {
        // TODO: Implement token generation logic
    },
}
