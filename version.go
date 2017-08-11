package version

import (
	"errors"
	"strings"
)

// version format as : prefix1.2.3.4, 4 can be empty
// prefix defined by user, 1 means major , 2 means minor, 3 means branch , 4 means build
// example: 1.2.3.4 , prefix = "", parse by "."
// example :go1.2.3 , prefix ="go" , 4 is empty

var Global = VersionParser{}

//prefix : if set , parse version by prefix, if not set, auto parse by first number;
//split : if set , parse versionnumbers by split, if not , auto parse by "." ;
type VersionData struct {
	Major string
	Minor string
	Branch string
	Build string
}

type VFormat struct{
	Prefix string
	SplitWith string
	Version VersionData
}

type Version interface{
	GetVersion(string)(VFormat, error)
}

type VersionParser struct{
	prefix string
	split string
}

func (v VersionParser)GetVersion(value string)( VFormat , error) {
	if value == ""{
		return VFormat{}, errors.New("Err, empty input value.")
	}

	prefix := v.prefix
	split := v.split

	if !strings.HasPrefix(value, prefix) && prefix != ""{
		return VFormat{}, errors.New("Err, invalid version format , e.g :v1.2.3.4, 4 can be empty.")
	}

	foundInt := false

	for i, d := range value{
		if d >= rune('0') && d <= rune('9'){
			foundInt = true
			prefix = value[0:i]
			break
		}
	}

	if !foundInt{
		return VFormat{}, errors.New("Err, invalid version format , e.g : v1.2.3.4, 4 can be empty.")
	}

	ver := strings.TrimPrefix(value, prefix)

	if split == ""{
		split = "."
	}

	data := strings.Split(ver, split)

	if len(data) > 4 || len(data) < 2 {
		return VFormat{}, errors.New("Err, invalid version format , e.g :v1.2.3.4, 4 can be empty.")
	}


	if len(data) == 3 {

		v3 := VersionData{
			data[0],
			data[1],
			data[2],
			"", }

		if !isValidVersionData(v3){
			return VFormat{}, errors.New("Err, invalid input version data." + value)
		}

		return VFormat{
			prefix,split,
			v3,
		}, nil
	}

	if len(data) == 4 {
		v4 := VersionData{
			data[0],
			data[1],
			data[2],
			data[3],
		}
		if !isValidVersionData(v4){
			return VFormat{}, errors.New("Err, invalid input version data." + value)
		}

		return VFormat{
			prefix,split,
			v4,
		}, nil
	}

	return VFormat{}, errors.New("Err, invalid version format , e.g :v1.2.3.4, 4 can be empty.")
}

func SetPrefix(value string){
	Global.prefix = value
}

func SetSplit(value string){
	Global.split = value
}

func SetGlobalVersionParser(prefix string , split string) {
	Global.prefix = prefix
	Global.split = split
}

func GetGlobalVersionParser() *VersionParser{
	return &Global
}

func CompareVersionF(v1 VFormat, v2 VFormat)(int, error){
	if v1.Prefix != v2.Prefix{
		return -2, errors.New("Err, prefix " + v1.Prefix + ", and prefix " + v2.Prefix + " not match, cannot compare.")
	}

	if v1.SplitWith!= v2.SplitWith{
		return -2, errors.New("Err, splitwith" + v1.SplitWith + ", and splitwith" + v2.SplitWith+ " not match, cannot compare.")
	}


	if v1.Version.Major > v2.Version.Major || v1.Version.Minor > v2.Version.Minor || v1.Version.Branch > v2.Version.Branch || v1.Version.Build > v2.Version.Build{
		return 1, nil
	}

	if v1.Version.Major == v2.Version.Major && v1.Version.Minor == v2.Version.Minor && v1.Version.Branch == v2.Version.Branch && v1.Version.Build == v2.Version.Build{
		return 0, nil
	}

	return -1, nil
}

func CompareVersion(v1 , v2 string)(int ,error){
	v11 , err := Global.GetVersion(v1)
	if err != nil{
		return -2, errors.New("Err, parse " + v1 +" err.")
	}

	v21 , err := Global.GetVersion(v2)
	if err != nil{
		return -2, errors.New("Err, parse " + v2 +" err.")
	}


	return CompareVersionF(v11, v21)
}

func isValidVersionData(data VersionData) bool{
	if !isValidNumber(data.Major) || !isValidNumber(data.Minor) || !isValidNumber(data.Branch){
		return false
	}

	if data.Build != "" {
		if !isValidNumber(data.Build){
			return false
		}
	}

	return true
}

func isValidNumber(number string)bool{
	for _, d := range number{
		if d < rune('0') || d > rune('9'){
			return false
		}
	}
	return true
}