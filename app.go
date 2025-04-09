package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"runtime"

	_ "github.com/mattn/go-sqlite3"
)

type App struct {
	ctx     context.Context
	db      *sql.DB
	version string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	home, _ := os.UserHomeDir()

	content, err := ioutil.ReadFile("./wails.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	var payload map[string]interface{}
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	infos := payload["info"].(map[string]interface{})

	a.version = "v" + infos["productVersion"].(string) + " - " + runtime.GOOS + " - " + runtime.GOARCH

	a.ctx = ctx
	if _, err := os.Stat(home + string(os.PathSeparator) + ".cheat_sheets.db"); os.IsNotExist(err) {
		a.db = nil
	} else {
		db, err := sql.Open("sqlite3", home+string(os.PathSeparator)+".cheat_sheets.db")
		check(err)
		a.db = db
	}
}

func (a *App) GetPrograms() string {
	tables := []string{}

	if a.db != nil {
		dbtables, err := a.db.Query("select name from sqlite_master where type = 'table'")
		check(err)

		for dbtables.Next() {
			var name string
			_ = dbtables.Scan(&name)
			tables = append(tables, name)
		}
	}

	jsonData, err := json.Marshal(tables)
	check(err)

	return string(jsonData)
}

func (a *App) GetVersion() string {
	return a.version
}

func (a *App) GetCheatSheet(program string) string {
	if a.db == nil {
		return "{}"
	}

	rows, err := a.db.Query("SELECT * FROM " + program)
	check(err)

	data := make(map[string][]map[string]string)

	for rows.Next() {
		var command string
		var about string
		var session string
		_ = rows.Scan(&command, &about, &session)
		commands := map[string]string{
			"command": command,
			"about":   about,
		}
		data[session] = append(data[session], commands)
	}

	jsonData, err := json.Marshal(data)
	check(err)

	return string(jsonData)
}
