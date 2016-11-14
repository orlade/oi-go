package oi

import (
	"strings"
	"testing"

	"log"
	"os/exec"

	"github.com/stretchr/testify/assert"
)

var searcher = new(HomePluginSearcher)

func TestHomePluginSearch_single(t *testing.T) {
	path := createHomeCommand("foo")
	defer executeChecked("rm -f " + path)

	result := searcher.Search()
	assert.Equal(t, path, result["foo"])
}

func createHomeCommand(name string) (path string) {
	dir := usr.HomeDir + "/.oi/modules/"
	executeChecked("mkdir -p " + dir)
	path = dir + name
	executeChecked("touch " + path)
	return
}

func executeChecked(command string) (err error) {
	words := strings.Split(command, " ")
	err = exec.Command(words[0], words[1:]...).Run()
	if err != nil {
		log.Fatalln(err)
	}
	return
}
