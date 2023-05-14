/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("test called")
		test()
	},
}

func init() {
	rootCmd.AddCommand(testCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type LaunchTemplateData struct {
	*types.ResponseLaunchTemplateData        // Embed the original struct
	CustomField                       string // Add additional fields
}

func NewLaunchTemplateData(data *types.ResponseLaunchTemplateData, customField string) *LaunchTemplateData {
	return &LaunchTemplateData{
		ResponseLaunchTemplateData: data,
		CustomField:                customField,
	}
}

// Add any additional methods you need here
func (d *LaunchTemplateData) PrettyPrint() {
	pp.Print(d)
}

func (d *LaunchTemplateData) GetCustomField() string {
	return d.CustomField
}

func test() {
	imageId := "ami-123"
	originalData := &types.ResponseLaunchTemplateData{
		ImageId: &imageId,
	}
	data := NewLaunchTemplateData(originalData, "my custom field value")

	// Access embedded struct fields using dot notation
	// fmt.Println(data.VersionNumber) // Output: 1
	// fmt.Println(data.ImageId)       // Output: ami-12345678

	// Access custom fields
	fmt.Println(data.CustomField) // Output: "my custom field value"

	// Call custom method
	fmt.Print(data.GetCustomField())
	data.PrettyPrint()
}
