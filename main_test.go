package main

import (
	"testing"
)

func TestCreateTable(t *testing.T) {

	db, err := getDbConnection()
	defer db.Close()

	if err != nil {
		t.Errorf("Error while opening mysql : %s", err.Error())
	} else {
		p := &DummyUser{}
		exists := db.HasTable(p)
		if exists {
			t.Skipf("Error table already created - Skipping the test")
		} else {
			t.Log("Creating table Product ")
			db.CreateTable(p)
			tableCreated := db.HasTable(p)
			if !tableCreated {
				t.Errorf("Error table not created : %t ", tableCreated)
			}
		}
	}
}
