package internalapi

import (
	"testing"
)

func TestNewPrivateAPI(t *testing.T) {
	api := NewPrivateAPI()
	if api == nil {
		t.Error("got nil on NewPrivateAPI call")
	}
}

var tablesQ = `SELECT name FROM sqlite_master WHERE type ='table'`

func TestSQLiteConnect(t *testing.T) {
	api := NewPrivateAPI()
	dbfile := "file::memory:?cache=shared"
	// This will tell SQLite to use a temporary database in system memory.
	db, err := api.LiteConnect(dbfile)
	if err != nil {
		t.Errorf("got error on LiteConnect method call: %v", err)
	}

	tableName := `temp_table1`
	createErr := db.Exec(`CREATE TABLE ` + tableName + `( name TEXT );`).Error
	if createErr != nil {
		t.Errorf("got error on sqlite create table: %v", createErr)
	}

	result := []map[string]interface{}{}
	selectErr := db.Raw(tablesQ).Scan(&result).Error
	if selectErr != nil {
		t.Errorf("got error on sqlite query: %v", selectErr)
	}
	// fmt.Printf("result: %#v\n", result)

	if len(result) == 0 {
		t.Error("not found any tables after create")
	}
	gotTableName, ok := result[0]["name"]
	if !ok {
		t.Error("not found row 'name' in the first row")
	}
	if gotTableName != tableName {
		t.Errorf("table name check failed. expected: %v, got: %v", tableName, gotTableName)
	}
}
