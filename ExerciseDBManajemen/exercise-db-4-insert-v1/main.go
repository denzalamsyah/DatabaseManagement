package main

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"

	_ "embed"

	_ "github.com/lib/pq"
)

type Credential struct {
	Host         string
	Username     string
	Password     string
	DatabaseName string
	Port         int
	Schema       string
}

func Connect(creds *Credential) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", creds.Host, creds.Username, creds.Password, creds.DatabaseName, creds.Port)

	// connect using database/sql + pq
	dbConn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return dbConn, nil
}

//go:embed insert.sql
var insertStr string

func InsertSQL(dbConn *sql.DB) error {
	_, err := dbConn.Exec(insertStr)
	if err != nil {
		return err
	}

	insertStr = `INSERT INTO students VALUES (1, 'Abdi', 'Doe','2003-12-01', 'Jakarta', '1A', 'active'),
	(2, 'Jane', 'Doe', '2004-02-01', 'Jakarta', '1A', 'active'),
	(3, 'Bernard', 'Smith', '2004-02-01', 'Jakarta', '1A', 'active'),
	(4, 'Jane', 'Smith', '2003-12-02', 'Jakarta', '1B', 'active'),
	(5, 'Andrew', 'Doe','2004-07-04', 'Jakarta', '1B', 'inactive'),
	(6, 'Rendy', 'Doe', '2004-06-10', 'Jakarta', '1B', 'inactive'),
	(7, 'John', 'Smith', '2004-05-11', 'Jakarta', '1B', 'inactive'),
	(8, 'Herry', 'Smith', '2004-04-12', NULL, '1B', 'active'),
	(9, 'John', 'William', '2004-03-20', NULL, '1B', 'active'),
	(10, 'Wendy', 'Doe', '2004-02-21', NULL, '1B', 'active')
	`
	_, err = dbConn.Exec(insertStr)
	if  err != nil {
		panic(err)
	}

	_ = ioutil.WriteFile("insert.sql", []byte(insertStr), 0644)
	fmt.Println("success insert data")
	return nil
}

var (
	sqlScript1 = "Q1JFQVRFIFRBQkxFIElGIE5PVCBFWElTVFMgc3R1ZGVudHMgKAoJCWlkIElOVCwgCgkJZmlyc3RfbmFtZSBWQVJDSEFSKDEwMCksCgkJbGFzdF9uYW1lIFZBUkNIQVIoMTAwKSwKCQlkYXRlX29mX2JpcnRoIERBVEUsCgkJYWRkcmVzcyBWQVJDSEFSKDI1NSksCgkJY2xhc3MgVkFSQ0hBUigxMDApLAoJCXN0YXR1cyBWQVJDSEFSKDEwMCkKCSk="
)

func CreateTable(dbConn *sql.DB) error {
	sqlScript1, _ := base64.StdEncoding.DecodeString(sqlScript1)
	_, err := dbConn.Exec(string(sqlScript1))

	if err != nil {
		return err
	}

	
	var create string
	create = `CREATE TABLE IF NOT EXISTS students(
		id INT,
		first_name VARCHAR(100),
		last_name VARCHAR(100),
		date_of_birth DATE,
		address VARCHAR(255),
		class VARCHAR(100),
		status VARCHAR(100)
	)`
	_, err = dbConn.Exec(create)
	if err != nil {
		panic(err)
	}
	fmt.Println("success create table")
	return nil
}

func main() {
	dbCredential := Credential{
		Host:         "localhost",
		Username:     "postgres",
		Password:     "rizwan123",
		DatabaseName: "Coba",
		Port:         5432,
	}
	dbConn, err := Connect(&dbCredential)
	if err != nil {
		log.Fatal(err)
	}

	_, err = dbConn.Exec("DROP TABLE IF EXISTS students CASCADE")
	if err != nil {
		log.Fatal(err)
	}

	err = CreateTable(dbConn)
	if err != nil {
		log.Fatal(err)
	}

	err = InsertSQL(dbConn)
	if err != nil {
		log.Fatal(err)
	}
}
