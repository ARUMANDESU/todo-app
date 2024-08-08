package internal

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"time"
)

func validateTitle(value any) error {
	title, ok := value.(string)
	if !ok {
		return fmt.Errorf("title must be a string")
	}

	return validation.Validate(title,
		validation.Length(3, 250).Error("title must be between 3 and 250 characters"),
	)
}

func validateDueDate(value any) error {
	dueDate, ok := value.(*time.Time)
	if !ok {
		return fmt.Errorf("due date must be a time.Time")
	}

	return validation.Validate(dueDate,
		validation.Min(time.Now().AddDate(0, 0, 1)).Error("due date must be at least tomorrow"),
	)
}
