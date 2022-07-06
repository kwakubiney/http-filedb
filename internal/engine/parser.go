package engine

import (
	"fmt"
	"strings"
	"github.com/sqlite-go"
)

func ParseAndExecuteStatement(inputBuffer sqlitego.InputBuffer, db *DB, statement Statement) error {
	if strings.HasPrefix(inputBuffer.Buffer, ".") {
		switch DoMetaCommand(inputBuffer, db) {
		case MetaCommandSuccess:
			return nil
		case MetaCommandUnrecognizedCommand:
			fmt.Printf("Unrecognized command %q \n", inputBuffer.Buffer)
			return nil
		}
	}

	switch PrepareStatement(inputBuffer, &statement) {
	case PrepareSuccess:

	case PrepareUnrecognizedStatement:
		fmt.Printf("Unrecognized command at start of %q \n", inputBuffer.Buffer)
		return nil

	case PrepareSyntaxError:
		fmt.Println("Syntax error. Could not parse statement.")
		return nil
	}

	_, err := ExecuteStatement(statement, db)
	if err != nil {
		return err
	}
	return nil
}
