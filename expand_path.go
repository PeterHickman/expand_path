package expand_path

import (
	"os"
	"os/user"
	"strings"
)

func getUser(name string) (string, int) {
	i := strings.Index(name, "/")

	if i == -1 {
		return name[1:], i
	} else {
		return name[1:i], i
	}
}

func normalisePath(path string) string {
	parts := strings.Split(path, "/")
	normalised := []string{}

	for _, part := range parts {
		if part == "." {
			// Do nothing
		} else if part == ".." {
			// Drop current head
			normalised = normalised[:len(normalised)-1]
		} else {
			normalised = append(normalised, part)
		}
	}

	return strings.Join(normalised, "/")
}

func ExpandPath(path string) (string, error) {
	newPath := path

	home, _ := os.UserHomeDir()

	// First look to expand any ~ paths
	if strings.HasPrefix(newPath, "~") {
		if newPath == "~" {
			newPath = home
		} else if strings.HasPrefix(newPath, "~/") {
			newPath = strings.Replace(newPath, "~", home, 1)
		} else if strings.HasPrefix(newPath, "~") {
			// A different users home directory
			username, continued := getUser(newPath)
			user, err := user.Lookup(username)
			if err != nil {
				return "", err
			}

			if continued == -1 {
				newPath = user.HomeDir
			} else {
				newPath = user.HomeDir + newPath[continued:]
			}
		}
	}

	// Patch up relative paths
	if !strings.HasPrefix(newPath, "/") {
		pwd, _ := os.Getwd()
		newPath = pwd + "/" + newPath
	}

	newPath = normalisePath(newPath)

	return newPath, nil
}
