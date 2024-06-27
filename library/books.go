package library

import (
	"os"
	"time"

	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/viper"
)

type Book struct {
	Title  string
	Author string
	Status map[string][]time.Time
}

func NewBook(title, author, status string) *Book {
	statusMap := make(map[string][]time.Time)
	statusMap[status] = append(statusMap[status], time.Now())
	return &Book{
		Title:  title,
		Author: author,
		Status: statusMap,
	}
}

func (b *Book) SaveBookToTOML() {
	dataDir := viper.GetString("data.dir")
	bookFile := dataDir + "/" + b.Title + ".toml"
	// marshal book to TOML
	data, err := toml.Marshal(b)
	if err != nil {
		panic(err)
	}
	// write to file using os
	err = os.WriteFile(bookFile, data, 0644)
	if err != nil {
		panic(err)
	}
}

func (b *Book) SaveBookToMarkdown() {
	mdDir := viper.GetString("data.mdDir")

	if _, err := os.Stat(mdDir); os.IsNotExist(err) {
		os.MkdirAll(mdDir, os.ModePerm)
	}

	mdFile := mdDir + "/" + b.Title + ".md"

	frontmatter := "+++\n"
	data, err := toml.Marshal(b)
	if err != nil {
		panic(err)
	}
	frontmatter += string(data)
	frontmatter += "+++\n"

	err = os.WriteFile(mdFile, []byte(frontmatter), 0644)
	if err != nil {
		panic(err)
	}
}

func (b *Book) UpdateStatus(status string) {
	b.Status[status] = append(b.Status[status], time.Now())
	b.SaveBookToTOML()
}

func (b *Book) GetCurrentStatusAndDate() (string, time.Time) {
	// each field in the status map is an array of dates
	// get the key with the most recent date
	var recentDate time.Time
	var recentStatus string
	for status, dates := range b.Status {
		if dates[len(dates)-1].After(recentDate) {
			recentDate = dates[len(dates)-1]
			recentStatus = status
		}
	}
	return recentStatus, recentDate
}

func (b *Book) GetDaysAtCurrentStatus() int {
	// get the current status and date
	_, date := b.GetCurrentStatusAndDate()
	// get the difference between now and the date
	return int(time.Since(date).Hours() / 24)
}

// a slice of objects that contain the status and date, sorted in ascending order
func (b *Book) GetStatusHistory() []map[string]interface{} {
	// create a slice of maps
	statusHistory := make([]map[string]interface{}, 0)
	// iterate through the status map
	for status, dates := range b.Status {
		// iterate through the dates
		for _, date := range dates {
			// create a map with status and date
			statusMap := map[string]interface{}{
				"status": status,
				"date":   date,
			}
			// append to the slice
			statusHistory = append(statusHistory, statusMap)
		}
	}
	return statusHistory
}
