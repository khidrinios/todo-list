package persistence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildTitleAndDescriptionSqlFilterString(t *testing.T) {
	title := "Todo"
	description := "First todo"
	actualFilter := buildTitleAndDescriptionSqlFilterString(&title, &description)
	expectedFilter := "title LIKE '%Todo%' AND description LIKE '%First todo%'"
	assert.Equal(t, expectedFilter, *&actualFilter)
}

func TestBuildTitleAndNoDescriptionSqlFilterString(t *testing.T) {
	title := "Todo"
	actualFilter := buildTitleAndDescriptionSqlFilterString(&title, nil)
	expectedFilter := "title LIKE '%Todo%'"
	assert.Equal(t, expectedFilter, *&actualFilter)
}

func TestBuildDescriptionAndNoTitleSqlFilterString(t *testing.T) {
	description := "First todo"
	actualFilter := buildTitleAndDescriptionSqlFilterString(nil, &description)
	expectedFilter := "description LIKE '%First todo%'"
	assert.Equal(t, expectedFilter, *&actualFilter)
}
