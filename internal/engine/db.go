package engine

import (
	"fmt"
	"os"
	"sync"
)

type DB struct {
	File          *os.File
	Path          string
	Opened        bool
	IndexFilePath string
	IndexFile     *os.File
	Bucket        map[string]int64
	Mutex         sync.Mutex
	Test          bool
}

func DbOpen(path string, indexFilePath string, mode os.FileMode, test string) (*DB, error) {
	var db = &DB{Opened: true}
	db.Bucket = make(map[string]int64)
	db.Path = path
	// for test purposes
	if test == "test"{
		db.Test = true
		var err error
		if db.File , err = os.CreateTemp("", "tmp_test_file-*"); err != nil{
			return nil, fmt.Errorf("error whilst opening test file, %s", err)
		}

		if db.IndexFile , err = os.CreateTemp("", "tmp_test_index_file-*"); err != nil{
			return nil, fmt.Errorf("error whilst opening test index file, %s", err)
		}
		ReadMapFromIndexFile(db)
	}else{
		db.Test = false
		db.IndexFilePath = indexFilePath
		flag := os.O_RDWR
		var err error
		if db.File, err = os.OpenFile(db.Path, os.O_APPEND|flag|os.O_CREATE, mode); err != nil {
			return nil, fmt.Errorf("error whilst opening file, %s", err)
		}
		if os.IsNotExist(err) {
			fmt.Println("Creating new db file...")
			createdFile, err := os.Create(path)
			if err != nil {
				return nil, fmt.Errorf("error whilst creating file, %s", err)
			}
			db.File = createdFile
		}
	
		if db.IndexFile, err = os.OpenFile(db.IndexFilePath, os.O_APPEND|flag|os.O_CREATE, mode); err != nil {
			return nil, fmt.Errorf("error whilst opening index file, %s", err)
		}
		if os.IsNotExist(err) {
			fmt.Println("Creating new index file...")
			createdIndexFile, err := os.Create(path)
			db.IndexFile = createdIndexFile
			if err != nil {
				return nil, fmt.Errorf("error whilst creating index file, %s", err)
			}
		}
		ReadMapFromIndexFile(db)
		RemoveIndexFile(db)
	}


	return db, nil

}

func (db *DB) Close() error {
	if !db.Opened {
		return nil
	}

	db.Opened = false

	WriteToIndexFile(db)

	if err := db.File.Close(); err != nil {
		return fmt.Errorf("failed to close db: %s", err)
	}

	if err := db.IndexFile.Close(); err != nil {
		return fmt.Errorf("failed to close index file: %s", err)
	}

	db.Path = ""
	db.IndexFilePath = ""
	return nil
}
