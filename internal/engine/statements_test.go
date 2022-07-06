package engine_test

import (
	"fmt"
	"log"
	"os"
	"testing"
	"math/rand"
	"time"
	"github.com/sqlite-go"
	"github.com/sqlite-go/internal/engine"
	"github.com/stretchr/testify/assert"
)

func TestPrepareStatement(t *testing.T) {
	type PrepareStatementTestSuite struct {
		Buffer    sqlitego.InputBuffer
		Statement engine.Statement
		Response  engine.PrepareResult
	}
	for _, buffer := range []PrepareStatementTestSuite{
		{
			sqlitego.InputBuffer{
				Buffer: "insert 1 adam kwaku@mail.com",
			},

			engine.Statement{},

			engine.PrepareSuccess,
		},
		{
			sqlitego.InputBuffer{
				Buffer: "insert adam kwaku kwaku@mail.com",
			},

			engine.Statement{},
			engine.PrepareSyntaxError,
		},
		{
			sqlitego.InputBuffer{
				Buffer: "some random unrecognizable statement",
			},
			engine.Statement{},
			engine.PrepareUnrecognizedStatement,
		}} {

		actualResponse := engine.PrepareStatement(buffer.Buffer, &buffer.Statement)
		assert.Equal(t, buffer.Response, actualResponse)
	}
}

func BenchmarkSerialization(b *testing.B) {
	row := engine.Row{
		ID:       "5",
		Username: "kwakubiney",
		Email:    "k@mail.com",
	}

	DB, err := engine.DbOpen("", "", 0644, "test")
	if err != nil {
		log.Println(err)
	}
	
	for n := 0; n < b.N; n++ {
		err := engine.SerializeRow(row, DB)
		if err != nil {
			log.Fatalln(err)
		}
	}
	defer os.Remove(DB.IndexFile.Name())
	defer DB.IndexFile.Close()
	defer os.Remove(DB.File.Name())
	defer DB.File.Close()
}


func BenchmarkDeserialization(b *testing.B) {
	DB, err := engine.DbOpen("", "", 0644, "test")
	if err != nil {
		log.Println(err)
	}

	for n := 0; n < 20; n++ {
		row := engine.Row{
			ID:       fmt.Sprint(n),
			Username: "kwakubiney",
			Email:    "k@mail.com",
		}
		err := engine.SerializeRow(row, DB)
		if err != nil {
			log.Fatalln(err)
		}
	}

	rand.Seed(time.Now().UnixNano())
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		n := rand.Intn(20 - 0) + 0
		b.StartTimer()
		_, err := engine.DeserializeSpecificRow(DB , fmt.Sprint(n))
		if err != nil {
			log.Fatalln(err)
		}
	}
	defer os.Remove(DB.IndexFile.Name())
	defer DB.IndexFile.Close()
	defer os.Remove(DB.File.Name())
	defer DB.File.Close()
}
