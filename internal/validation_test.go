package internal

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

func TestTitleValidation(t *testing.T) {
	tests := []struct {
		name      string
		title     string
		expectErr bool
		errMsg    string
	}{
		{"ValidLength", "Valid Title", false, ""},
		{"TooShort", "No", true, "title must be between 3 and 250 characters"},
		{"TooLong", string(make([]byte, 251)), true, "title must be between 3 and 250 characters"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateTitle(tt.title)
			if tt.expectErr {
				assert.Error(t, err)
				assert.Equal(t, tt.errMsg, err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDueDateValidation(t *testing.T) {
	tests := []struct {
		name      string
		dueDate   any
		expectErr bool
		errMsg    string
	}{
		{"ValidDate", func() *time.Time { t := time.Now().AddDate(0, 0, 1); return &t }(), false, ""},
		{"PastDate", func() *time.Time { t := time.Now().AddDate(0, 0, -2); return &t }(), true, "due date must be at least tomorrow"},
		{"InvalidType", "invalid type", true, "due date must be a time.Time"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateDueDate(tt.dueDate)
			if tt.expectErr {
				assert.Error(t, err)
				assert.Equal(t, tt.errMsg, err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDescriptionValidation(t *testing.T) {
	tests := []struct {
		name        string
		description *string
		expectErr   bool
		errMsg      string
	}{
		{"ValidDescription", func() *string { s := "This is a valid description."; return &s }(), false, ""},
		{"EmptyDescription", func() *string { s := ""; return &s }(), false, ""},
		{"TooLongDescription", func() *string { s := string(make([]byte, 1001)); return &s }(), true, "must be less than 1000 characters"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateDescription(tt.description)
			if tt.expectErr {
				assert.Error(t, err)
				assert.Equal(t, tt.errMsg, err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestTagsValidation(t *testing.T) {
	tests := []struct {
		name      string
		tags      []string
		expectErr bool
		errMsg    string
	}{
		{"ValidTags", []string{"tag1", "tag2", "tag3"}, false, ""},
		{"TooManyTags", make([]string, 16), true, "must be less than 15 tags"},
		{"TagTooShort", []string{"ta"}, true, "must be between 3 and 50 characters"},
		{"TagTooLong", []string{string(make([]byte, 51))}, true, "must be between 3 and 50 characters"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateTags(tt.tags)
			if tt.expectErr {
				assert.Error(t, err)
				assert.True(t, strings.Contains(err.Error(), tt.errMsg))
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
