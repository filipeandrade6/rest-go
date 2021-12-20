package infile_test

import (
	"testing"

	"github.com/filipeandrade6/rest-go/app/infile"
)

func TestDB(t *testing.T) {
	db := infile.New()

	if err := db.Create("a", "1"); err != nil {
		t.Errorf("expected nil, got '%s'", err)
	}

	if err := db.Create("b", "2"); err != nil {
		t.Errorf("expected nil, got '%s'", err)
	}

	if err := db.Create("a", "3"); err == nil {
		t.Error("expected already exist, got nil")
	}

	value, err := db.Read("a")
	if err != nil || value != "1" {
		t.Errorf("expected 1, got '%s' and error '%q'", value, err)
	}

	value, err = db.Read("c")
	if err == nil {
		t.Errorf("expected doesn't exist, got '%s' and error '%q'", value, err)
	}

	err = db.Update("a", "57")
	if err != nil {
		t.Errorf("expected nil, got '%q'", err)
	}

	value, err = db.Read("a")
	if err != nil {
		t.Errorf("expected 57, got '%s' and error '%q'", value, err)
	}

	err = db.Update("d", "1")
	if err == nil {
		t.Errorf("expected doesn't exist, got '%q'", err)
	}

	err = db.Delete("a")
	if err != nil {
		t.Errorf("expected nil, got '%q'", err)
	}

	value, err = db.Read("a")
	if err == nil {
		t.Errorf("expected doesn't exist, got '%s' and error '%q'", value, err)
	}

	err = db.Delete("n")
	if err == nil {
		t.Errorf("expected doesn't exist, got '%q'", err)
	}

}
