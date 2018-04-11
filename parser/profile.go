package parser

import (
	"github.com/fatrbaby/cobweb/engine"
	"github.com/fatrbaby/cobweb/entity"
	"regexp"
	"strconv"
)

var (
	AgeMatcher = regexp.MustCompile(`<td><span[^>]*>年龄：</span>([\d]+)岁</td>`)
	GenderMatcher = regexp.MustCompile(`<td><span[^>]*>性别：</span><span field="">([^<]+)</span></td>`)
	MarriageMatcher = regexp.MustCompile(`<td><span[^>]*>婚况：</span>([^<]+)</td>`)
	HeightMatcher = regexp.MustCompile(`<td><span[^>]*>身高：</span><span field="">([\d]+)CM</span></td>`)
	WeightMatcher = regexp.MustCompile(`<td><span[^>]*>体重：</span><span field="">([\d]+)</span></td>`)
	IncomeMatcher = regexp.MustCompile(`<td><span[^>]*>月收入：</span>([^<]+)</td>`)
	EducationMatcher = regexp.MustCompile(`<td><span[^>]*>学历：</span>([^<]+)</td>`)
	OccupationMatcher = regexp.MustCompile(`<td><span[^>]*>职业： </span>([^<]+)</td>`)
	ConstellationMatcher = regexp.MustCompile(`<td><span[^>]*>星座：</span><span field="">([^<]+)</span></td>`)
	HouseMatcher = regexp.MustCompile(`<td><span[^>]*>住房条件：</span><span field="">([^<]+)</span></td>`)
	CarMatcher = regexp.MustCompile(`<td><span[^>]*>是否购车：</span><span field="">([^<]+)</span></td>`)
)

func ProfileParser(contents []byte, name string) engine.ParsedResult {
	profile := entity.Profile{}

	profile.Name = name
	profile.Age = extractInt(contents, AgeMatcher)
	profile.Gender = extractString(contents, GenderMatcher)
	profile.Marriage = extractString(contents, MarriageMatcher)
	profile.Height = extractInt(contents, HeightMatcher)
	profile.Weight = extractInt(contents, WeightMatcher)
	profile.Income = extractString(contents, IncomeMatcher)
	profile.Education = extractString(contents, EducationMatcher)
	profile.Occupation = extractString(contents, OccupationMatcher)
	profile.Constellation = extractString(contents, ConstellationMatcher)
	profile.House = extractString(contents, HouseMatcher)
	profile.Car = extractString(contents, CarMatcher)

	result := engine.ParsedResult{
		Items: []interface{}{profile},
	}

	return result
}

func extractString(contents []byte, matcher *regexp.Regexp) string {
	matched := matcher.FindSubmatch(contents)

	if matched == nil {
		return ""
	}

	return string(matched[1])
}

func extractInt(contents []byte, matcher *regexp.Regexp) int {
	r := extractString(contents, matcher)

	num, err := strconv.Atoi(r)

	if err != nil {
		return 0
	}

	return num
}