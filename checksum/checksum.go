package checksum

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
)

func GetAllTables(db *sql.DB) ([]string, error) {
	log.Println("Fetching all tables from the database")
	rows, err := db.Query("SHOW TABLES")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tables []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		tables = append(tables, name)
	}
	log.Println("Tables found :", len(tables))
	return tables, nil
}

func ChecksumTable(db *sql.DB, table string) (string, error) {
	query := fmt.Sprintf("SELECT * FROM `%s`", table)
	log.Println("Executing checksum for ", table, " : ", query)
	rows, err := db.Query(query)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	columns, _ := rows.Columns()
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	hasher := md5.New()

	for rows.Next() {
		if err := rows.Scan(scanArgs...); err != nil {
			return "", err
		}
		for _, val := range values {
			hasher.Write(val)
		}
	}
	sum := hex.EncodeToString(hasher.Sum(nil))
	log.Println("Checksum for table", table, "is", sum)
	return sum, nil
}

func ChecksumAllTables(db *sql.DB) map[string]string {
	results := make(map[string]string)

	tables, err := GetAllTables(db)
	if err != nil {
		log.Println("Error getting tables : ", err)
		return results
	}

	for _, table := range tables {
		sum, err := ChecksumTable(db, table)
		if err != nil {
			log.Printf("Error checksumming table %s: %v\n", table, err)
			continue
		}
		results[table] = sum
	}

	return results
}
