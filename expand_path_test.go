package expand_path

import (
	"os"
	"os/user"
	"strings"
	"testing"
)

var home, _ = os.UserHomeDir()
var pwd, _ = os.Getwd()
var currentUser, _ = user.Current()

func compare(t *testing.T, test_path, expected_path string) {
	actual_path, _ := ExpandPath(test_path)

	if expected_path != actual_path {
		t.Fatalf("%s did not expand to %s, got %s", test_path, expected_path, actual_path)
	}
}

func TestHomeOnly(t *testing.T) {
	compare(t, "~", home)
}

func TestHomeAndPath(t *testing.T) {
	compare(t, "~/fred", home+"/fred")
}

func TestNamedUserOnly(t *testing.T) {
	compare(t, "~"+currentUser.Username, home)
}

func TestNamedUserAndPath(t *testing.T) {
	compare(t, "~"+currentUser.Username+"/fred", home+"/fred")
}

func TestRelativePath(t *testing.T) {
	compare(t, "fred", pwd+"/fred")
}

func TestRelativePathAndFile(t *testing.T) {
	compare(t, "fred/another", pwd+"/fred/another")
}

func TestSingleDot(t *testing.T) {
	compare(t, ".", pwd)
}

func TestSingleDotAndFile(t *testing.T) {
	compare(t, "./fred", pwd+"/fred")
}

func TestSingleDotMidPath(t *testing.T) {
	compare(t, "/a/b/./c", "/a/b/c")
}

func TestUserAndSingleDot(t *testing.T) {
	compare(t, "~/.", home)
}

func TestUserAndSingleDotAndFile(t *testing.T) {
	compare(t, "~/./fred", home+"/fred")
}

func TestDoubleDot(t *testing.T) {
	i := strings.LastIndex(pwd, "/")

	compare(t, "..", pwd[0:i])
}

func TestDoubleDotAndFile(t *testing.T) {
	i := strings.LastIndex(pwd, "/")

	compare(t, "../fred", pwd[0:i]+"/fred")
}

func TestDoubleDotMidPath(t *testing.T) {
	compare(t, "/a/b/../c", "/a/c")
}

func TestNamedUserDoesNotExist(t *testing.T) {
	dummy_user := "~aslkdhfaslkdch"
	_, err := ExpandPath(dummy_user)

	if err == nil {
		t.Fatalf("%s should have triggered an error", dummy_user)
	}
}
