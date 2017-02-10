package main

import "testing"

func TestAppName(t *testing.T) {
	expect := "Zoo Application"
	actual := AppName()

	if expect != actual {
		t.Error("%s != %s", expect, actual)
	}
}
