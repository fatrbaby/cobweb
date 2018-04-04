package parser

import (
	"regexp"
)

type City struct {
	Name []byte
	Link []byte
}


const (
	CityListPattern = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`
)

func FindCities(content []byte) []City {
	re := regexp.MustCompile(CityListPattern)

	matches := re.FindAllSubmatch(content, -1)
	var cities = make([]City, len(matches))

	for _, match := range matches {
		cities = append(cities, City{match[2], match[1]})
	}

	return cities
}
