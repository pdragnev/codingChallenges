/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ccwc",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var reader io.Reader
		var file *os.File
		var filename string
		var err error
		var bytes int64

		if len(args) == 1 {
			filename = args[0]
			file, err = os.Open(filename)
			if err != nil {
				fmt.Println("Error opening file:", err)
				return
			}
			defer file.Close()
			reader = file

			// Always get file size, as it's used in the default output
			if fileInfo, err := file.Stat(); err == nil {
				bytes = fileInfo.Size()
			} else {
				fmt.Println("Error getting file size:", err)
				return
			}
		} else {
			reader = os.Stdin
		}

		lines, words, characters, err := countMetrics(reader)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		if lineFlag {
			fmt.Printf("%d lines in %s\n", lines, filename)
		}
		if wordFlag {
			fmt.Printf("%d words in %s\n", words, filename)
		}
		if byteFlag {
			fmt.Printf("%d bytes in %s\n", bytes, filename)
		}
		if charactersFlag {
			fmt.Printf("%d characters in %s\n", characters, filename)
		}

		// Default case when no specific flag is provided
		if !lineFlag && !wordFlag && !byteFlag && !charactersFlag {
			fmt.Printf("%d lines, %d words, %d bytes in %s\n", lines, words, bytes, filename)
		}
	},
}

func countMetrics(reader io.Reader) (int, int, int, error) {
	var lines, words, characters int
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()
		lines++
		words += len(strings.Fields(line))
		characters += len(line)
	}

	if err := scanner.Err(); err != nil {
		return 0, 0, 0, err
	}

	return lines, words, characters, nil
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var lineFlag, wordFlag, byteFlag, charactersFlag bool

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.Flags().BoolVarP(&lineFlag, "lines", "l", false, "Count lines in a file")
	rootCmd.Flags().BoolVarP(&wordFlag, "words", "w", false, "Count words in a file")
	rootCmd.Flags().BoolVarP(&byteFlag, "bytes", "c", false, "Count bytes in a file")
	rootCmd.Flags().BoolVarP(&charactersFlag, "characters", "m", false, "Count characters in a file")

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ccwc.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
