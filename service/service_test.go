package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPagination(t *testing.T) {
	page := 2
	limit := 3
	actualOffset := getOffsetFromPageAndLimit(page, limit)
	expectedOffset := 3
	assert.Equal(t, expectedOffset, actualOffset)
}
