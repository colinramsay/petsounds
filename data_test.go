package main

import (
	"os"
	"testing"
)

func TestWriteSettings(t *testing.T) {
	tmpCfg := "./tmpcfg.json"

	os.Remove(tmpCfg)

	Settings{
		TorrentConfiguration{
			"blackholecfg",
			"proxyurlcfg",
			"completedircfg",
		},
		"ppcfg",
	}.Write(tmpCfg)

	f, err := os.Open(tmpCfg)
	fi, err := f.Stat()

	if err != nil || fi.Size() == 0 {
		t.Fatal("config file was empty")
	}

	os.Remove(tmpCfg)
}

func TestReadSettings(t *testing.T) {
	settings := ReadSettings("./testdata/sample.config.json")

	if settings.PostProcessingScript != "pp" {
		t.Fatal("PostProcessingScript was not pp")
	}
}

// import (
// 	"database/sql"
// 	_ "github.com/mattn/go-sqlite3"
// 	"testing"
// 	"os"
// )

// func TestCreateTables(t *testing.T) {
// 	os.Remove(databaseFilename)
// 	checkDatabase()

// 	db, err := sql.Open("sqlite3", databaseFilename)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer db.Close()

// 	tables := []string{"artists", "albums"}

// 	for _, table := range tables {
// 		stmt, err := db.Prepare("SELECT count(*) FROM sqlite_master WHERE type='table' AND name=?")

// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 		defer stmt.Close()

// 		rows, err := stmt.Query(table)

// 		defer rows.Close()

// 		if err != nil {
// 			t.Error(err)
// 		}

// 		for rows.Next() {
// 			var count int
// 			rows.Scan(&count)
// 			if count != 1 {
// 				t.Errorf("Table count was actually: %d for %s", count, table)
// 			}
// 		}
// 	}

// 	// This should not throw a duplicate table error
// 	checkDatabase()
// }
