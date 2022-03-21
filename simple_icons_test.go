package simple_icons_go

import (
	"errors"
	"io/ioutil"
	"runtime"
	"testing"
)

func TestLoad(t *testing.T) {
	tests := []struct {
		name     string
		want     SimpleIcon
		mockFunc func()
		wantErr  bool
	}{
		{
			name: "simple icon load fails",
			mockFunc: func() {
				runtimeCaller = func(skip int) (pc uintptr, file string, line int, ok bool) {
					return 2, "file", 0, false
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()
			_, err := Load()
			if (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
	runtimeCaller = runtime.Caller
}

func TestSimpleIcon_Get(t *testing.T) {
	ioutilReadFile = func(filename string) ([]byte, error) {
		return nil, errors.New("unable to load file")
	}
	r := Release{"0.0.0", "tests"}
	slugs, _ := r.GetSlugs()
	icons, _ := r.GetIcons()
	si := SimpleIcon{
		release: r,
		slugs:   slugs,
		icons:   icons,
	}
	_, err := si.Get("liquibase")
	if err == nil {
		t.Error("svg load error not caught")
	}
	ioutilReadFile = ioutil.ReadFile
}
