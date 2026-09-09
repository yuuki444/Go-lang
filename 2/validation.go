package main

import "errors"

func validateMovie(input MovieInput) error {
	if input.Title == "" {
		return errors.New("title cannot be empty")
	}
	if input.Director == "" {
		return errors.New("director cannot be empty")
	}
	if input.Year <= 1888 || input.Year > 2100 {
		return errors.New("year must be between 1889 and 2100")
	}
	if input.Rating < 0 || input.Rating > 10 {
		return errors.New("rating must be between 0 and 10")
	}
	return nil
}
