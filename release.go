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

type Release struct {
	Version    string
	ModuleRoot string
}

func LoadRelease(v string) Release {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		panic(errors.New("unable to find module root"))
	}
	r := Release{
		Version:    v,
		ModuleRoot: path.Dir(filename),
	}
	return r
}

func (r Release) GetSlugs() []Slug {
	var s []Slug
	file, err := os.Open(r.ModuleRoot + "/assets/simple-icons/simple-icons-" + r.Version + "/slugs.md")
	if err != nil {
		panic(err)
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
	return s
}

func (r Release) GetIcons() Icons {
	var i Icons
	rawData, err := ioutil.ReadFile(r.ModuleRoot + "/assets/simple-icons/simple-icons-" + r.Version + "/_data/simple-icons.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(rawData, &i)
	if err != nil {
		panic(err)
	}
	return i
}

func (r Release) GetSvg(s string) ([]byte, error) {
	return ioutil.ReadFile(r.ModuleRoot + "/assets/simple-icons/simple-icons-" + r.Version + "/icons/" + s + ".svg")
}
