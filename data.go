package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type TorrentConfiguration struct {
	PirateBayProxy       string
	BlackHoleDirectory   string
	CompleteDirectory    string
}

type Settings struct {
	TorrentConfiguration TorrentConfiguration
	PostProcessingScript string
}

const databaseFilename = "petsounds.sqlite"

func checkDatabase() {
	var rowCount int
	db, err := sql.Open("sqlite3", databaseFilename)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	row := db.QueryRow("SELECT COUNT(name) FROM sqlite_master WHERE type='table' AND name='artists';")

	err = row.Scan(&rowCount)

	if rowCount == 0 {
		_, err = db.Exec(`
	    	CREATE TABLE artists(
			   	id INT PRIMARY KEY     		NOT NULL,
			   	name           CHAR(200)    NOT NULL
			);
	    `)

		if err != nil {
			log.Fatal(err)
		}
	}

	row = db.QueryRow("SELECT COUNT(name) FROM sqlite_master WHERE type='table' AND name='albums';")

	err = row.Scan(&rowCount)

	if rowCount == 0 {
		_, err = db.Exec(`
	    	CREATE TABLE albums(
			   	id INT PRIMARY KEY     		NOT NULL,
			   	name           CHAR(200)    NOT NULL,
			   	artist_id		INT 		NOT NULL
			);
	    `)

		if err != nil {
			log.Fatal(err)
		}
	}
}
