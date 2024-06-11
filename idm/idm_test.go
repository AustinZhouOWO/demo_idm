package idm

import "testing"

func TestTrim(t *testing.T) {
	name := "Austin Zhou"
	want := []string{"Austin", "Zhou"}
	testName := Trim(name)
	if len(testName) != len(want) || testName[0] != want[0] || testName[1] != want[1] {
		t.Fatalf(`Expect: %s :Get: %s`, want, testName)
	}
}
