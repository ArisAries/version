package version

import "testing"

func TestVersionParser_GetVersion(t *testing.T){

	v, err := GetGlobalVersionParser().GetVersion("1.1.1")
	if err != nil{
		t.Error("Err when parse version 1.1.1" )
	}

	t.Log(v)
}

func TestCompareVersionF(t *testing.T){
	v, err := CompareVersionF(
		VFormat{"",",",VersionData{"1","1","1","1",}},
		VFormat{"",",",VersionData{"1","2","3","4"}})

	if err != nil{
		t.Error("Err when parse version 1.1.1" )
	}

	t.Log(v)
}

func TestCompareVersion(t *testing.T){
	v, err := CompareVersion("1.1.1.1","2.1.1.1")
	if err != nil{
		t.Error("Err when parse version 1.1.1" )
	}

	t.Log(v)
}

