package engine_test

import (
	"fmt"
	"log"
	"testing"
	"github.com/sqlite-go/internal/engine"
	"github.com/stretchr/testify/assert"
	"os"
)

var DB *engine.DB

func TestSerializationAndDeserialization(t *testing.T) {
	DB, err := engine.DbOpen("", "", 0644, "test")
	if err != nil {
		log.Println(err)
	}
	
	type DeserializeCases struct {
		Row              engine.Row
		ExpectedResponse string
	}

	for _, row := range []DeserializeCases{
		{
			engine.Row{ID: "5",
				Username: "Kwaku Egyir Biney",
				Email:    "kwakz@yahoo.com",
			},
			"ID: 5, Name: Kwaku Egyir Biney, Email: kwakz@yahoo.com"},
		{
			engine.Row{ID: "10",
				Username: "Sally Akua",
				Email:    "sally@yahoo.com",
			},
			"ID: 10, Name: Sally Akua, Email: sally@yahoo.com"},
	} {
		err := engine.SerializeRow(row.Row, DB)
		assert.NoError(t, err)

		deserializedRow, err := engine.DeserializeSpecificRow(DB, row.Row.ID)

		assert.NoError(t, err)

		assert.Equal(t, fmt.Sprintf("ID: %s, Name: %s, Email: %s", row.Row.ID, row.Row.Username, row.Row.Email), deserializedRow)
	}
	
	defer os.Remove(DB.IndexFile.Name())
	defer DB.IndexFile.Close()
	defer os.Remove(DB.File.Name())
	defer DB.File.Close()
}


