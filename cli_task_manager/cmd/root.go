package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/volodimyr/vikinGO/cli_task_manager/persistent"
	"os"
	"strings"
)

func init() {
	rootCmd.AddCommand(do, add, list, rm, completed)
}

var rootCmd = &cobra.Command{
	Use:   "task_manager",
	Short: "To do list. Make your life easier.",
	Long:  `Be cool to organize your life with this application. Add, delete, list and remove your daily routine tasks.`,
	Run: func(cmd *cobra.Command, args []string) {
		// some code
	},
}

var do = &cobra.Command{
	Use:   "do",
	Short: "Mark a task on your TODO list as complete",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Doing your task...")
	},
}

var add = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to your TODO list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		if task == "" {
			fmt.Println("Task cannot be empty")
			return
		}
		done := persistent.AddTask(strings.Join(args, " "))
		if done {
			fmt.Println("Task has been added.")
		}
	},
}

var list = &cobra.Command{
	Use:   "list",
	Short: "List all of your incomplete tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks := persistent.ViewTasks()
		if len(tasks) == 0 {
			fmt.Println("Not found")
		}
		index := 1
		for _, v := range tasks {
			fmt.Printf("%d. %s\n", index, v.Name)
			index++
		}
	},
}

var completed = &cobra.Command{
	Use:   "completed",
	Short: "List all completed tasks today",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("List of completed tasks...")
	},
}

var rm = &cobra.Command{
	Use:   "remove",
	Short: "Remove your daily routine task",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Removing your task...")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
