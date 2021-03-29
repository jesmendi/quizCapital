package quiz

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/jesmendi/quizCapital/config"
)

type CountryData struct {
	CountryID string
	Country   string
	Capital   string
}

func init() {

	Countries = LoadData()
}

// Load the data in csv file
func LoadData() []CountryData {

	csvfile, err := os.Open(config.C.Conf.File)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	ret := make([]CountryData, 0)
	// Parse the file
	r := csv.NewReader(csvfile)

	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		country := CountryData{
			CountryID: record[0],
			Country:   record[1],
			Capital:   record[2],
		}
		ret = append(ret, country)
	}

	fmt.Println("The file is read")
	return ret

}
