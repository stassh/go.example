package main

import (
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"net/http"
	"net/http/httptest"
	"swapid/entity"
	"swapid/global"
	"testing"
)

func TestReadCharacter(t *testing.T) {

	rawDB, mock, _ := sqlmock.New()
	global.DB, _ = gorm.Open(global.DBDialect, rawDB)
	global.DB.LogMode(true)

	rows := mock.NewRows([]string{"id", "name"}).AddRow(1, "Luke")

	mock.ExpectQuery(`^SELECT \* FROM "characters".*$`).WillReturnRows(rows)

	req, err := http.NewRequest("GET", "/characters/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(readCharacter)
	handler.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("Wrong status code %v", status)
	}

	var character entity.Character
	decoder := json.NewDecoder(responseRecorder.Body)

	if err := decoder.Decode(&character); err != nil {
		t.Fatal(err)
	}

	if character.Name != "Luke" {
		t.Fatal("Wrong name")
	}
}

func TestListCharacter(t *testing.T) {

	rawDB, mock, _ := sqlmock.New()
	global.DB, _ = gorm.Open(global.DBDialect, rawDB)
	global.DB.LogMode(true)

	rows := mock.NewRows([]string{"id", "name"}).AddRow(1, "Luke").AddRow(2, "Puck").AddRow(3, "Duke")

	mock.ExpectQuery(`^SELECT \* FROM "characters"`).WillReturnRows(rows)

	req, err := http.NewRequest("GET", "/characters", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(listCharacters)
	handler.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("Wrong status code %v", status)
	}

	var characters []entity.Character
	decoder := json.NewDecoder(responseRecorder.Body)

	if err := decoder.Decode(&characters); err != nil {
		t.Fatal(err)
	}

	if characters[0].Name != "Luke" {
		t.Fatal("Wrong name")
	}

}
