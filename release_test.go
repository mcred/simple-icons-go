package simple_icons_go

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"runtime"
	"testing"
)

func TestLoadRelease(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name     string
		args     args
		want     Release
		mockFunc func()
		wantErr  bool
	}{
		{
			name: "module root fails to load",
			args: args{"0.0.0"},
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
			_, err := LoadRelease(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadRelease() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
	runtimeCaller = runtime.Caller
}

func TestRelease_GetIcons(t *testing.T) {
	type fields struct {
		Version    string
		ModuleRoot string
	}
	tests := []struct {
		name     string
		fields   fields
		want     Icons
		mockFunc func()
		wantErr  bool
	}{
		{
			name:   "icons file fails to open",
			fields: fields{"0.0.0", "tests"},
			mockFunc: func() {
				ioutilReadFile = func(filename string) ([]byte, error) {
					return nil, errors.New("unable to open file")
				}
			},
			wantErr: true,
		},
		{
			name:   "icons file fails to unmarshall",
			fields: fields{"0.0.0", "tests"},
			mockFunc: func() {
				jsonUnmarshall = func(data []byte, v interface{}) error {
					return errors.New("unable to unmarshall")
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()
			r := Release{
				Version:    tt.fields.Version,
				ModuleRoot: tt.fields.ModuleRoot,
			}
			_, err := r.GetIcons()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetIcons() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
	ioutilReadFile = ioutil.ReadFile
	jsonUnmarshall = json.Unmarshal
}

func TestRelease_GetSlugs(t *testing.T) {
	type fields struct {
		Version    string
		ModuleRoot string
	}
	tests := []struct {
		name     string
		fields   fields
		want     []Slug
		mockFunc func()
		wantErr  bool
	}{
		{
			name:   "slugs file fails to open",
			fields: fields{"0.0.0", "tests"},
			mockFunc: func() {
				osOpen = func(name string) (*os.File, error) {
					return nil, errors.New("unable to open file")
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()
			r := Release{
				Version:    tt.fields.Version,
				ModuleRoot: tt.fields.ModuleRoot,
			}
			_, err := r.GetSlugs()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSlugs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
	osOpen = os.Open
}
