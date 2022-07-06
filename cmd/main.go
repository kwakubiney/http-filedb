package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/sqlite-go"
	"github.com/sqlite-go/internal/config"
	"github.com/sqlite-go/internal/engine"
	"github.com/sqlite-go/internal/handlers"
	"github.com/sqlite-go/internal/server"
)

func PrintPrompt() {
	fmt.Printf("db > ")
}

func main() {

	config.LoadMainConfig("../.env")
	cmd := flag.String("cmd", "", "")
	flag.Parse()
	argument := *cmd
	db, err := engine.DbOpen("../db", "../index", 0644, "")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if argument == "cli" {
		inputBuffer := sqlitego.NewInputBuffer()
		scanner := bufio.NewScanner(os.Stdin)
		for {
			PrintPrompt()
			scanner.Scan()
			command := scanner.Text()
			inputBuffer.Buffer = command
			var statement engine.Statement
			err := engine.ParseAndExecuteStatement(inputBuffer, db, statement)
			if err != nil {
				log.Println(err)
			}
		}

	} else if argument == "http" {
		handlers := handlers.New(db)
		server := server.New(handlers)
		server.Start()
	} else {
		fmt.Println("Unknown arguments. Refer to docs.")
	}
}
