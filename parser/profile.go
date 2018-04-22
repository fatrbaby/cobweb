package parser

import (
	"github.com/fatrbaby/cobweb/engine"
	"github.com/fatrbaby/cobweb/entity"
	"regexp"
	"strconv"
)

var (
	UrlIdMatcher         = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)
	AgeMatcher           = regexp.MustCompile(`<td><span[^>]*>年龄：</span>([\d]+)岁</td>`)
	GenderMatcher        = regexp.MustCompile(`<td><span[^>]*>性别：</span><span field="">([^<]+)</span></td>`)
	MarriageMatcher      = regexp.MustCompile(`<td><span[^>]*>婚况：</span>([^<]+)</td>`)
	HeightMatcher        = regexp.MustCompile(`<td><span[^>]*>身高：</span><span field="">([\d]+)CM</span></td>`)
	WeightMatcher        = regexp.MustCompile(`<td><span[^>]*>体重：</span><span field="">([\d]+)</span></td>`)
	IncomeMatcher        = regexp.MustCompile(`<td><span[^>]*>月收入：</span>([^<]+)</td>`)
	EducationMatcher     = regexp.MustCompile(`<td><span[^>]*>学历：</span>([^<]+)</td>`)
	OccupationMatcher    = regexp.MustCompile(`<td><span[^>]*>职业： </span>([^<]+)</td>`)
	ConstellationMatcher = regexp.MustCompile(`<td><span[^>]*>星座：</span><span field="">([^<]+)</span></td>`)
	HouseMatcher         = regexp.MustCompile(`<td><span[^>]*>住房条件：</span><span field="">([^<]+)</span></td>`)
	CarMatcher           = regexp.MustCompile(`<td><span[^>]*>是否购车：</span><span field="">([^<]+)</span></td>`)
	RecommendMatcher     = regexp.MustCompile(`<a class="exp-user-name"[^>]*href="(http://album.zhenai.com/u/[\d]+)">([^<]+)</a>`)
)

type ProfileParser struct {
	Username string
}

func (p *ProfileParser) Parse(contents []byte, url string) engine.ParsedResult {
	profile := entity.Profile{}

	profile.Name = p.Username
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
		Items: []engine.Item{
			{
				Id:      extractString([]byte(url), UrlIdMatcher),
				Url:     url,
				Type:    "zhenai",
				Payload: profile,
			},
		},
	}

	matches := RecommendMatcher.FindAllSubmatch(contents, -1)

	for _, match := range matches {
		result.Spiders = append(
			result.Spiders,
			engine.Spider{
				Url:    string(match[1]),
				Parser: &ProfileParser{Username: string(match[2])},
			},
		)
	}

	return result
}

func (p *ProfileParser) Serialize() (name string, args interface{}) {
	return "ProfileParser", p.Username
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
