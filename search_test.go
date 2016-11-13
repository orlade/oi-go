package oi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var searcher = new(HomePluginSearcher)

func TestHomePluginSearch(t *testing.T) {
	assert.Equal(t, searcher.Search(), map[string]string{"git": OIPATH + "/oi-git"})
}
