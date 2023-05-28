package main

import (
	"database/sql"
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

type Column struct {
	ColumnName string `sql:"column_name"`
	IsNullable string `sql:"is_nullable"`
	DataType   string `sql:"data_type"`
}

func Connect(creds *Credential) (*sql.DB, error) {
	// this is only an example, please modify it to your need
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", creds.Host, creds.Username, creds.Password, creds.DatabaseName, creds.Port)

	// connect using database/sql + pq
	dbConn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return dbConn, nil
}

func CreateTable(db *sql.DB) error {
	_, err := db.Exec(`DROP TABLE IF EXISTS students`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students (
		id INTEGER PRIMARY KEY,
		fullname VARCHAR(255) NOT NULL,
		address TEXT,
		gender VARCHAR(50) NOT NULL,
		day_of_birth INTEGER,
		month_of_birth INTEGER,
		year_of_birth INTEGER,
		grade INTEGER
	)`)

	if err != nil {
		return err
	}

	fmt.Println("Table created successfully")
	return nil
}

//go:embed add_column_table_students.sql
var alterAdd string

//go:embed drop_column_table_students.sql
var alterDrop string

//go:embed modify_column_table_students.sql
var alterModify string


func AlterAdd(db *sql.DB) error {

	alterAdd = `
	ALTER TABLE students ADD street VARCHAR(255);
	ALTER TABLE students ADD city VARCHAR(100);
	ALTER TABLE students ADD province VARCHAR(100);
	ALTER TABLE students ADD country VARCHAR(100);
	ALTER TABLE students ADD postal_code VARCHAR(50);
	ALTER TABLE students ADD date_of_birth DATE NOT NULL
	
	`
	_, err := db.Exec(alterAdd)
	if err != nil {
		fmt.Println("Failed to alter table add column")
		return err
	}

	err = ioutil.WriteFile("add_column_table_students.sql", []byte(alterAdd), 0644)
	if err != nil {
		fmt.Println("Failed to write file")
		return err
	}

	fmt.Println("Alter add table success")
	return nil

}

func AlterDrop(db *sql.DB) error {
	alterDrop = `
	ALTER TABLE students DROP COLUMN address;
	ALTER TABLE students DROP COLUMN day_of_birth;
	ALTER TABLE students DROP COLUMN month_of_birth;
	ALTER TABLE students DROP COLUMN year_of_birth`
	_, err := db.Exec(alterDrop)
	if err != nil {
		fmt.Println("Failed to alter table drop column")
		return err
	}

	err = ioutil.WriteFile("drop_column_table_students.sql", []byte(alterDrop), 0644)
	if err != nil {
		fmt.Println("Failed to write file")
		return err
	}

	fmt.Println("Alter drop table success")
	return nil

}

func AlterModify(db *sql.DB) error {

	alterModify = `ALTER TABLE students ALTER COLUMN grade TYPE FLOAT`
	_, err := db.Exec(alterModify)
	if err != nil {
		fmt.Println("Failed to alter table modify column")
		return err
	}

	err = ioutil.WriteFile("modify_column_table_students.sql", []byte(alterModify), 0644)
	if err != nil {
		fmt.Println("Failed to write file")
		return err
	}

	fmt.Println("Alter modify table success")
	return nil
}

func main() {
	//Change this with your database credential
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

	err = CreateTable(dbConn)
	if err != nil {
		log.Fatal(err)
	}

	err = AlterAdd(dbConn)
	if err != nil {
		log.Fatal(err)
	}

	err = AlterDrop(dbConn)
	if err != nil {
		log.Fatal(err)
	}

	err = AlterModify(dbConn)
	if err != nil {
		log.Fatal(err)
	}
}
