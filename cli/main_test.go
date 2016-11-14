package main

import (
	"log"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	version := executeChecked("go run main.go -v")
	assert.Equal(t, "oi version 0.0.1\n", version)
}

func executeChecked(command string) (output string) {
	words := strings.Split(command, " ")
	log.Println("Executing:", command)
	stdout, err := exec.Command(words[0], words[1:]...).Output()
	output = string(stdout)

	if err != nil {
		log.Fatalln("Error:", err)
	}
	return
}
