package simple_icons_go

import (
	"bufio"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type Release struct {
	Version string
	File    string
}

func LoadRelease(v string) Release {
	r := Release{
		Version: v,
		File:    "vendor/simple-icons/" + v + ".zip",
	}
	return r
}

func (r Release) GetSlugs() []Slug {
	var s []Slug
	file, _ := os.Open("vendor/simple-icons/simple-icons-" + version + "/slugs.md")
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
	rawData, err := ioutil.ReadFile("vendor/simple-icons/simple-icons-" + version + "/_data/simple-icons.json")
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
	return ioutil.ReadFile("vendor/simple-icons/simple-icons-" + r.Version + "/icons/" + s + ".svg")
}
