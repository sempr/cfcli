/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"
)

// downCmd represents the down command
var downCmd = &cobra.Command{
	Use:   "down",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		contest, _ := cmd.Flags().GetString("contest")
		problem, _ := cmd.Flags().GetString("problem")
		download(contest, problem)
	},
}

func download(contest, code string) {
	// Request the HTML page.
	client := &http.Client{}
	url := fmt.Sprintf("https://codeforces.com/contest/%s/problem/%s", contest, code)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Cookie", "RCPC=4cdf00515de2d551305448315600819f")

	res, err := client.Do(req)
	// res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	in_idx := 0
	doc.Find(".sample-tests .input pre").Each(func(i int, s *goquery.Selection) {
		fmt.Println("CaseIn: ", in_idx)
		in_idx++
		f, err := os.Create(fmt.Sprintf("%d.in", in_idx))
		if err != nil {
			log.Panic(err)
		}
		defer f.Close()
		f.WriteString(s.Text())
	})
	out_idx := 0
	doc.Find(".sample-tests .output pre").Each(func(i int, s *goquery.Selection) {
		fmt.Println("CaseOut: ", out_idx)
		out_idx++
		f, err := os.Create(fmt.Sprintf("%d.out", out_idx))
		if err != nil {
			log.Panic(err)
		}
		defer f.Close()
		f.WriteString(s.Text())
	})
}

func init() {
	rootCmd.AddCommand(downCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	downCmd.Flags().StringP("contest", "c", "1000", "A help for foo")
	downCmd.Flags().StringP("problem", "p", "A", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// downCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
