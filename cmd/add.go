/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
    // import type from todo package
    "github.com/hrutvikyadav/go-cli/todo"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new todo",
	Long: `Add a task to list.
You can specify title, priority, due date and project`,
	Run: addRun,
}

func addRun(cmd *cobra.Command, args []string) {
    // initialize list with empty todoitem
    list := []todo.Item{}
    for _, av := range args {
        // add to list
        list = append(list, todo.Item{Title: av})
        fmt.Println("Added", av)
    }
    //fmt.Println(list).todods.json
    err := todo.SaveItems("/home/hrutvik_/.todods.json", list)
    if err != nil {
        fmt.Errorf("%v", err)
    }

}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
