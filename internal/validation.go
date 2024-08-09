package internal

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"time"
)

func validateTitle(value any) error {
	title, ok := value.(string)
	if !ok {
		return fmt.Errorf("must be a string")
	}

	return validation.Validate(title,
		validation.Length(3, 250).Error("must be between 3 and 250 characters"),
	)
}

func validateDueDate(value any) error {
	dueDate, ok := value.(*time.Time)
	if !ok {
		return fmt.Errorf("must be a *time.Time")
	}

	return validation.Validate(dueDate,
		validation.Min(time.Now().AddDate(0, 0, -1)).Error("must be at least tomorrow"),
	)
}

func validateDescription(value any) error {
	description, ok := value.(*string)
	if !ok {
		return fmt.Errorf("must be a *string")
	}

	return validation.Validate(description,
		validation.Length(0, 1000).Error("must be less than 1000 characters"),
	)
}

func validateTags(value any) error {
	tags, ok := value.([]string)
	if !ok {
		return fmt.Errorf("must be a []string")
	}

	return validation.Validate(tags,
		validation.Length(0, 15).Error("must be less than 15 tags"),
		validation.Each(validation.Length(3, 50).Error("must be between 3 and 50 characters")),
	)
}
