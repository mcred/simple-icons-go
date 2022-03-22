package tests_test

import (
	"github.com/mcred/simple-icons-go"
	"reflect"
	"testing"
)

func TestLoadPackage(t *testing.T) {
	si, err := simple_icons_go.Load()
	if err != nil {
		t.Error("unable to load package")
	}
	if reflect.TypeOf(si) != reflect.TypeOf(simple_icons_go.SimpleIcon{}) {
		t.Errorf("unable to load package")
	}
}
