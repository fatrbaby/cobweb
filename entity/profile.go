package entity

import "encoding/json"

type Profile struct {
	Name          string
	Gender        string
	Age           int
	Height        int
	Weight        int
	Income        string
	Marriage      string
	Education     string
	Occupation    string // 职业
	Constellation string // 星座
	House         string
	Car           string
}

func FromJsonObject(obj interface{}) (Profile, error)  {
	var profile Profile
	j, err := json.Marshal(obj)

	if err != nil {
		return profile, err
	}

	err = json.Unmarshal(j, &profile)

	return profile, err
}
