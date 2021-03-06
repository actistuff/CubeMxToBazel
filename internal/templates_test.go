package cubemxtobazelinternal

import (
	"testing"
)

func TestAttributeBString(t *testing.T) {
	op := attributeBString{Key: "name", Value: "hello_world"}
	expected := `name="hello_world",`
	got := op.Attribute().String()
	if expected != got {
		t.Errorf("Expected:\n%#v \nGot:\n%#v \n", expected, got)
	}
}

func TestBStringList(t *testing.T) {
	var list bStringList = []string{"a", "b", "c"}
	expected := `["a","b","c",]`
	got := list.String()
	if expected != got {
		t.Errorf("Expected:\n%#v \nGot:\n%#v \n", expected, got)
	}
}

func TestAttributeBBool(t *testing.T) {
	op := attributeBBool{Key: "linkstatic", Value: true}
	expected := `linkstatic=True,`
	got := op.Attribute().String()
	if expected != got {
		t.Errorf("Expected:\n%#v \nGot:\n%#v \n", expected, got)
	}
}
func TestBBoolString(t *testing.T) {
	var toggle bBool = true
	expected := "True"
	got := toggle.String()
	if expected != got {
		t.Errorf("Expected:\n%#v \nGot:\n%#v \n", expected, got)
	}
	toggle = false
	expected = "False"
	got = toggle.String()
	if expected != got {
		t.Errorf("Expected:\n%#v \nGot:\n%#v \n", expected, got)
	}
}

func TestAttributeStringList(t *testing.T) {
	var list bStringList = []string{"a", "b", "c"}
	srcs := attSrcs
	expected := `srcs=["a","b","c",],`
	strList := attributeBStringList{Key: srcs, Value: list}
	got := strList.Attribute().String()
	if expected != got {
		t.Errorf("Expected:\n%#v \nGot:\n%#v \n", expected, got)
	}
}

func TestAttributeListString(t *testing.T) {
	name := attributeBString{Key: "name", Value: "hello_world"}
	linkOpt := attributeBBool{Key: "linkstatic", Value: true}
	expected := name.Attribute().String() + linkOpt.Attribute().String()
	var list attributeList = []attribute{name, linkOpt}
	got := list.String()
	if expected != got {
		t.Errorf("Expected:\n%#v \nGot:\n%#v \n", expected, got)
	}
}

func TestCCLibraryString(t *testing.T) {
	name := attributeBString{Key: "name", Value: "hello_world"}
	linkOpt := attributeBBool{Key: "linkstatic", Value: true}
	var srclist bStringList = []string{"a", "b", "c"}
	srcs := attributeBStringList{Key: attSrcs, Value: srclist}
	var list attributeList = []attribute{name, srcs, linkOpt}
	ccLibrary := CcLibraryRule{rule{Keys: list, comment: comment{Comment: "#Hello test"}}}
	expected := "#Hello test\ncc_library(name=\"hello_world\",srcs=[\"a\",\"b\",\"c\",],linkstatic=True,)\n"
	got := ccLibrary.String()
	if expected != got {
		t.Errorf("Expected:\n%#v \nGot:\n%#v \n", expected, got)
	}
}

func TestCCBinaryString(t *testing.T) {
	name := attributeBString{Key: "name", Value: "hello_world"}
	linkOpt := attributeBBool{Key: "linkstatic", Value: true}
	var srclist bStringList = []string{"a", "b", "c"}
	srcs := attributeBStringList{Key: attSrcs, Value: srclist}
	var list attributeList = []attribute{name, srcs, linkOpt}
	ccBinary := CcBinaryRule{rule{Keys: list, comment: comment{Comment: "#Hello test"}}}
	expected := "#Hello test\ncc_binary(name=\"hello_world\",srcs=[\"a\",\"b\",\"c\",],linkstatic=True,)\n"
	got := ccBinary.String()
	if expected != got {
		t.Errorf("Expected:\n%#v \nGot:\n%#v \n", expected, got)
	}
}

func TestCCImportString(t *testing.T) {
	name := attributeBString{Key: "name", Value: "hello_world"}
	linkOpt := attributeBBool{Key: "linkstatic", Value: true}
	var srclist bStringList = []string{"a", "b", "c"}
	srcs := attributeBStringList{Key: attSrcs, Value: srclist}
	var list attributeList = []attribute{name, srcs, linkOpt}
	ccImport := ccImportRule{rule{Keys: list, comment: comment{Comment: "#Hello test"}}}
	expected := "#Hello test\ncc_import(name=\"hello_world\",srcs=[\"a\",\"b\",\"c\",],linkstatic=True,)\n"
	got := ccImport.String()
	if expected != got {
		t.Errorf("Expected:\n%#v \nGot:\n%#v \n", expected, got)
	}
}
