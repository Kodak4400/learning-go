package main

import (
	"fmt"
	"log"
	"spreadsheet/pkg"

	"google.golang.org/api/sheets/v4"
)

const (
	spreadsheetId      = "1rjjZGqIQUb-IVZbY6ngFCSQVOVZVR5CiEm_Ex7EH_TI"
	credentialFilePath = "./credential.json"
)

func main() {
	client, err := pkg.HttpClient(credentialFilePath)
	if err != nil {
		log.Fatal(err)
	}
	service, _ := sheets.New(client)

	rangeStr := "シート1"
	resp, _ := service.Spreadsheets.Values.Get(spreadsheetId, rangeStr).Do()
	if err != nil {
		log.Fatalf("Unable to get Spreadsheets. %v", err)
	}

	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
	} else {
		fmt.Println("Name, Major:")
		for _, row := range resp.Values {
			fmt.Printf("%s\n", row[0])
		}
	}

}
