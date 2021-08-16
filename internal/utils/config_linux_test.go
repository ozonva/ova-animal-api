// +build linux

package utils_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"os/exec"
	. "ova-animal-api/internal/utils"
	"strings"
	"testing"
)

const MissingFile = "/some/missing/file"

func TestLoadConfig(t *testing.T) {
	openFilesBefore := countOpenFiles()
	LoadConfig(ConfigFileName)
	openFilesAfter := countOpenFiles()
	assert.Equal(t, openFilesBefore, openFilesAfter)
}

func TestLoadConfigsTenTimes(t *testing.T) {
	openFilesBefore := countOpenFiles()
	LoadConfigTenTimes()
	openFilesAfter := countOpenFiles()
	assert.Equal(t, openFilesBefore, openFilesAfter)
}

// No way to test error on io.ReadAll :(
//func TestLoadConfigForbidden(t *testing.T) {
//	openFilesBefore := countOpenFiles()
//	assert.Panics(t, func() {
//		LoadConfig(ForbiddenReadFile)
//	})
//	openFilesAfter := countOpenFiles()
//	assert.Equal(t, openFilesBefore, openFilesAfter)
//}

func TestLoadConfigMissing(t *testing.T) {
	openFilesBefore := countOpenFiles()
	assert.Panics(t, func() {
		LoadConfig(MissingFile)
	})
	openFilesAfter := countOpenFiles()
	assert.Equal(t, openFilesBefore, openFilesAfter)
}

func countOpenFiles() int {
	out, err := exec.Command("/bin/sh", "-c", fmt.Sprintf("lsof -p %v", os.Getpid())).Output()
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(out), "\n")
	return len(lines) - 1
}
