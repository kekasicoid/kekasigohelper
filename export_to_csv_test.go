package kekasigohelper

import "testing"

func TestExportToCSV(t *testing.T) {
	header := []string{"firt_name", "last_name"}
	var body [][]string
	record := []string{"Arditya", "kekasi"}
	body = append(body, record)
	if err := ExportToCSV("./kekasi.co.id.csv", header, body); err != nil {
		t.Error(err)
	}
}
