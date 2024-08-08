package internal

import (
	"github.com/stretchr/testify/assert"
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
