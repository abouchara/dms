// cmd/query.go
package cmd

import (
   	"encoding/json"
    "fmt"
    "github.com/spf13/cobra"
    "depview/pkg/storage"
    "depview/pkg/models"
)

var queryCmd = &cobra.Command{
    Use:   "query [entity]",
    Short: "Query dependencies for a givent element of infrastructure",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        entityType := args[0]
        name, _ := cmd.Flags().GetString("name")

        switch entityType {
        case "credentials":
            var creds models.Credentials
            storage.LoadData(fmt.Sprintf("data/creds/creds.json"), &creds)
            jsonData, err := json.MarshalIndent(creds, "", "  ")
            if err != nil {
                fmt.Println("Error marshaling JSON:", err)
                return
            }

            fmt.Printf("Json data: %s\n", jsonData)

            // fmt.Printf("Credentials for [%s] : [%s]\n", name, creds)
            fmt.Printf("Dependencies for credentials [%s]: %v\n", name, creds.UsedIn)

        case "app":
            var app models.Application
            storage.LoadData(fmt.Sprintf("data/apps/%s.json", name), &app)
            fmt.Printf("Dependencies for app %s: %v\n", name, app.Dependencies)

        case "server":
            var server models.Server
            storage.LoadData(fmt.Sprintf("data/servers/%s.json", name), &server)
            fmt.Printf("Dependencies for server %s: %v\n", name, server.Dependencies)
        }
    },
}

func init() {
    rootCmd.AddCommand(queryCmd)
    queryCmd.Flags().String("name", "", "Name of the element to query")
}
