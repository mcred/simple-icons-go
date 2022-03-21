package simple_icons_go

import (
	"bufio"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"
)

//functions to override during testing
var (
	runtimeCaller  = runtime.Caller
	osOpen         = os.Open
	ioutilReadFile = ioutil.ReadFile
	jsonUnmarshall = json.Unmarshal
)

type Release struct {
	Version    string
	ModuleRoot string
}

func LoadRelease(v string) (Release, error) {
	_, filename, _, ok := runtimeCaller(1)
	if !ok {
		return Release{}, errors.New("unable to find module root")
	}
	r := Release{
		Version:    v,
		ModuleRoot: path.Dir(filename),
	}
	return r, nil
}

func (r Release) GetSlugs() ([]Slug, error) {
	var s []Slug
	file, err := osOpen(r.ModuleRoot + "/assets/simple-icons/simple-icons-" + r.Version + "/slugs.md")
	if err != nil {
		return []Slug{}, errors.New("unable to open slugs source")
	}
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if strings.HasPrefix(line, "|") {
			split := strings.Split(line, "|")
			if split[1] != "Brand name" && split[1] != ":---" {
				slug := Slug{Name: split[1][2 : len(split[1])-2], Slug: split[2][2 : len(split[2])-2]}
				s = append(s, slug)
			}
		}
		if err == io.EOF {
			break
		}
	}
	return s, nil
}

func (r Release) GetIcons() (Icons, error) {
	var i Icons
	rawData, err := ioutilReadFile(r.ModuleRoot + "/assets/simple-icons/simple-icons-" + r.Version + "/_data/simple-icons.json")
	if err != nil {
		return i, errors.New("unable to open simple-icons source")
	}
	err = jsonUnmarshall(rawData, &i)
	if err != nil {
		return i, errors.New("unable to unmarshall simple-icons source")
	}
	return i, nil
}

func (r Release) GetSvg(s string) ([]byte, error) {
	return ioutilReadFile(r.ModuleRoot + "/assets/simple-icons/simple-icons-" + r.Version + "/icons/" + s + ".svg")
}
