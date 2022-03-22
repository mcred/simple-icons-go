package simple_icons_go

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"runtime"
	"testing"
)

func TestRelease_Load(t *testing.T) {
	_, err := LoadRelease("0.0.0")
	if err != nil {
		t.Error("release failed to load")
	}
}

func TestRelease_LoadFails(t *testing.T) {
	runtimeCaller = func(skip int) (pc uintptr, file string, line int, ok bool) {
		return 2, "file", 0, false
	}
	_, err := LoadRelease("0.0.0")
	if err == nil {
		t.Error("module root expected failure")
	}
	runtimeCaller = runtime.Caller
}

func TestRelease_GetIconSourceFails(t *testing.T) {
	ioutilReadFile = func(filename string) ([]byte, error) {
		return nil, errors.New("unable to open file")
	}
	_, err := testRelease.GetIcons()
	if err == nil {
		t.Error("get icon expected file failure")
	}
	ioutilReadFile = ioutil.ReadFile
}

func TestRelease_GetIconUnmarshallFails(t *testing.T) {
	jsonUnmarshall = func(data []byte, v interface{}) error {
		return errors.New("unable to unmarshall")
	}
	_, err := testRelease.GetIcons()
	if err == nil {
		t.Error("get icon expected json unmarshall failure")
	}
	jsonUnmarshall = json.Unmarshal
}

func TestRelease_GetSlugsFails(t *testing.T) {
	osOpen = func(name string) (*os.File, error) {
		return nil, errors.New("unable to open file")
	}
	_, err := testRelease.GetSlugs()
	if err == nil {
		t.Error("get slugs expected file failure")
	}
	osOpen = os.Open
}
