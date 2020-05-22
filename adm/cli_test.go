package main

import (
	"github.com/HongJaison/go-admin/modules/system"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetLatestVersion(t *testing.T) {
	assert.Equal(t, getLatestVersion(), system.Version())
}
