package todo_test

import (
	"khidr/todo/api/todo"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildTitleAndDescriptionSqlFilterString(t *testing.T) {
	title := "Todo"
	description := "First todo"
	actualFilter := todo.BuildTitleAndDescriptionSqlFilterString(&title, &description)
	expectedFilter := "title LIKE '%Todo%' AND description LIKE '%First todo%'"
	assert.Equal(t, expectedFilter, actualFilter)
}

func TestBuildTitleAndNoDescriptionSqlFilterString(t *testing.T) {
	title := "Todo"
	actualFilter := todo.BuildTitleAndDescriptionSqlFilterString(&title, nil)
	expectedFilter := "title LIKE '%Todo%'"
	assert.Equal(t, expectedFilter, actualFilter)
}

func TestBuildDescriptionAndNoTitleSqlFilterString(t *testing.T) {
	description := "First todo"
	actualFilter := todo.BuildTitleAndDescriptionSqlFilterString(nil, &description)
	expectedFilter := "description LIKE '%First todo%'"
	assert.Equal(t, expectedFilter, actualFilter)
}
