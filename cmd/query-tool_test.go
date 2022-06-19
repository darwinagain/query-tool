package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	filePath := "../query_params.csv"
	workers := 2

	stats := Run(&filePath, &workers)

	assert.NotNil(t, stats)
}
