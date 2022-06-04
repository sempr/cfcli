/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		lang, _ := cmd.Flags().GetString("lang")
		path_ := path.Join(os.Getenv("HOME"), "/.codeforces/templates", "t."+lang)
		source_file, _ := filepath.Abs(path_)
		fmt.Println(source_file)
		if _, err := os.Stat(source_file); err != nil {
			log.Fatal(err)
		}
		target_file := fmt.Sprintf("code.%s", lang)
		if _, err := os.Stat(target_file); err == nil {
			log.Fatal("File exsit")
		}
		sf, err := os.Open(source_file)
		fmt.Print(source_file)
		if err != nil {
			log.Fatal(err)
		}
		defer sf.Close()
		tf, err := os.Create(target_file)
		if err != nil {
			log.Fatal(err)
		}
		defer tf.Close()
		if _, err := io.Copy(tf, sf); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	newCmd.PersistentFlags().StringP("lang", "l", "cc", "language")
	newCmd.PersistentFlags().BoolP("force", "f", false, "force replace")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
