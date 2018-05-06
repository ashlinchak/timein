package lib

import (
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
)

var config Config

type Table struct{}

func (table *Table) Render(tag string, t time.Time) {
	initConfig()

	tw := tablewriter.NewWriter(os.Stdout)
	tw.SetAutoWrapText(false)

	if config.ShowHeaders {
		if tag == "" {
			tw.SetHeader([]string{"Tag", "Time", "Data", "Time zone"})
		} else {
			tw.SetHeader([]string{"Time", "Data", "Time zone"})
		}
	}

	for _, v := range getData().Items {
		if len(tag) > 0 {
			if tag == v.Tag {
				row := tableRow(&v, tag, t)
				tw.Append(row)
			}
		} else {
			row := tableRow(&v, tag, t)
			tw.Append(row)
		}
	}
	tw.Render()
}

func tableRow(dataItem *DataItem, tag string, t time.Time) []string {
	data := strings.Join(dataItem.Payload, "\n")
	// time
	localTime := GetTimeFromTimezone(dataItem.Timezone, &t)
	// fmt.Println(localTime)
	timeCell := color.GreenString(localTime.Format("2006-01-02 15:04:05"))

	if tag == "" {
		return []string{dataItem.Tag, timeCell, data, dataItem.Timezone}
	} else {
		return []string{timeCell, data, dataItem.Timezone}
	}
}

func initConfig() {
	config = Config{}
	config.SetDefaults()
}

func getData() Data {
	data := Data{}
	data.Get()
	return data
}
