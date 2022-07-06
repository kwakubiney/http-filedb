package engine

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"log"
	"strings"
)

func NewDecoder(db *DB) *gob.Decoder {
	return gob.NewDecoder(db.IndexFile)
}

//TODO: Make this a transaction
func DeserializeAllRows(db *DB) ([]string, error) {
	var arrayOfRows []string
	for k := range db.Bucket {
		row, err := DeserializeSpecificRow(db, k)
		arrayOfRows = append(arrayOfRows, row)
		if err != nil {
			return nil, err
		}
	}
	return arrayOfRows, nil
}

func ParseDecodedRow(row string) string {
	sliceOfDecodedRowValues := strings.Split(row, ":")
	rowToConsole := fmt.Sprintf("ID: %s, Name: %s, Email: %s", sliceOfDecodedRowValues[0], sliceOfDecodedRowValues[1], sliceOfDecodedRowValues[2])
	return rowToConsole
}

func DeserializeSpecificRow(db *DB, id string) (string, error) {
	db.Mutex.Lock()
	defer db.Mutex.Unlock()
	offsetOfRow, ok := db.Bucket[id]
	if !ok {
		return "", fmt.Errorf("no such row id: %s in the database", id)
	}
	var rowLength [4]byte
	_, err := db.File.ReadAt(rowLength[:], offsetOfRow)
	if err != nil {
		log.Println(err)
		return "", err
	}
	rowLengthBuffer := bytes.NewReader(rowLength[:])
	var rowLengthAsUInt32 uint32
	err = binary.Read(rowLengthBuffer, binary.BigEndian, &rowLengthAsUInt32)
	if err != nil {
		log.Println(err)
		return "", err
	}

	rowBuffer := make([]byte, uint8(rowLengthAsUInt32))
	offsetOfRowData := offsetOfRow + 4
	_, err = db.File.ReadAt(rowBuffer[:], offsetOfRowData)
	if err != nil {
		log.Println(err)
		return "", err
	}
	decodedRow := ParseDecodedRow(string(rowBuffer))
	if db.Test{
		return decodedRow, err
	}
	fmt.Println(decodedRow)
	return decodedRow, err
}
