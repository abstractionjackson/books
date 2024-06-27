package library

import (
	"os"

	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/viper"
)

func FindBookByTitle(title string) *Book {
	// get the data directory
	dataDir := viper.GetString("data.dir")
	// the file is the title + .toml
	bookFile := dataDir + "/" + title + ".toml"
	// check for the presence of the file
	if _, err := os.Stat(bookFile); os.IsNotExist(err) {
		return nil
	}
	// unmarshall the file to a book
	data, err := os.ReadFile(bookFile)
	if err != nil {
		panic(err)
	}
	var book Book
	err = toml.Unmarshal(data, &book)
	if err != nil {
		panic(err)
	}
	return &book
}
