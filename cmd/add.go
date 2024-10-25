// cmd/add.go
package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "depview/pkg/models"
    "depview/pkg/storage"
)

var addCmd = &cobra.Command{
    Use:   "add [entity]",
    Short: "Add a new dependency or element)",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        entityType := args[0]

        switch entityType {
        case "credentials":
            name, _ := cmd.Flags().GetString("name")

            credentials := models.Credentials{
                ID:          "0",
                Type:        "credentials",
                Name:        name,
                UsedIn:      []string{},
            }
            type Credentials struct {
                ID           string   `json:"id"`
                Type         string   `json:"type"`
                Name         string   `json:"name"`
                Descr        string   `json:"desc`
                Tags         []string `json:"tags"`
                TmpTags      []string `json:"tmp_tags"`
                UsedIn       []string `json:"used_in"`
            }
            // Save the server to JSON
            storage.SaveData(fmt.Sprintf("data/creds/%s.json", name), credentials)

            fmt.Printf("Added server: %s\n", name)

        case "server":
            name, _ := cmd.Flags().GetString("name")
            ip, _ := cmd.Flags().GetString("ip")
            os, _ := cmd.Flags().GetString("os")

            server := models.Server{
                ID:          name,
                Type:        "server",
                Name:        name,
                IPAddress:   ip,
                OS:          os,
                Dependencies: []string{},
            }

            // Save the server to JSON
            storage.SaveData(fmt.Sprintf("data/servers/%s.json", name), server)

            fmt.Printf("Added server: %s\n", name)

        case "app":
            name, _ := cmd.Flags().GetString("name")
            version, _ := cmd.Flags().GetString("version")
            deps, _ := cmd.Flags().GetStringSlice("depends-on")

            app := models.Application{
                ID:          name,
                Type:        "application",
                Name:        name,
                Version:     version,
                Dependencies: deps,
            }

            // Save the application to JSON
            storage.SaveData(fmt.Sprintf("data/apps/%s.json", name), app)

            fmt.Printf("Added application: %s\n", name)
        }
    },
}

func init() {
    rootCmd.AddCommand(addCmd)

    // Server flags
    addCmd.Flags().String("name", "", "Name of the entity")
    addCmd.Flags().String("ip", "", "IP address of the server (for server entity)")
    addCmd.Flags().String("os", "", "Operating System (for server entity)")

    // App flags
    addCmd.Flags().String("version", "", "Version of the application (for app entity)")
    addCmd.Flags().StringSlice("depends-on", []string{}, "Dependencies for the application")
}
