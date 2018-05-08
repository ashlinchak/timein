package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ashlinchak/timein/lib"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var selectedTime string
var selectedTag string

var rootCmd = &cobra.Command{
	Use:     "timein [tag]",
	Short:   "Show your data by time zones",
	Example: "timein berlin -t=21:00",
	Run:     runRootCommand,
}

func runRootCommand(cmd *cobra.Command, args []string) {
	if len(args) > 0 {
		selectedTag = args[0]
	}

	localTime := time.Now()
	if selectedTime != "" {
		timeArr := strings.Split(selectedTime, ":")
		hour, err := strconv.Atoi(timeArr[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		minute, err := strconv.Atoi(timeArr[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		localTime = time.Date(localTime.Year(), localTime.Month(), localTime.Day(), hour, minute, 0, 0, localTime.Location())
	}

	printTime(localTime)
	fmt.Println()

	renderTable(selectedTag, localTime)
}

func renderTable(tag string, t time.Time) {
	table := lib.Table{}
	table.Render(tag, t)
}

func printTime(t time.Time) {
	fmt.Fprintln(color.Output, fmt.Sprintf("Time: %s", color.CyanString(t.Format("2006-01-02 15:04:05"))))
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&selectedTime, "time", "t", "", "Time value e.g. \"10:00\"")
}
