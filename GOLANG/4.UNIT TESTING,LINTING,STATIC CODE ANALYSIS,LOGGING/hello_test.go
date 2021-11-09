package main

import (
	"testing"
)

func TestHelloEmptyArg(t *testing.T) {

	EmptyResult := hello("")
	if EmptyResult != "Hello Dude!" {
		t.Errorf("hello(\"\") failed,expected %v,got %v", "Hello Dude!", EmptyResult)
	} else {
		t.Logf("hello(\"\") success,expected %v,got %v", "Dude!", EmptyResult)
	}

}

func TestHelloValidArg(t *testing.T) {

	Result := hello("Mike")
	if Result != "Hello Mike!" {
		t.Errorf("hello(\"Mike\") failed,expected %v,got %v", "Hello Dude!", Result)
	} else {
		t.Logf("hello(\"Mike\") success,expected %v,got %v", "Dude!", Result)
	}
}
