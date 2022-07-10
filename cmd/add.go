package cmd

import (
	"log"

	"github.com/robsongomes/tri/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var priority int

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Priority:1,2,3")
}

func addRun(cmd *cobra.Command, args []string) {
	items, _ := todo.ReadItems(viper.GetString("datafile"))
	for _, arg := range args {
		item := todo.Item{Text: arg}
		item.SetPriority(priority)
		items = append(items, item)
	}
	err := todo.SaveItems(viper.GetString("datafile"), items)
	if err != nil {
		log.Printf("%v", err)
	}
}
