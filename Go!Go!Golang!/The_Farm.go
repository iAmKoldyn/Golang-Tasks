package thefarm

import (
	"errors"
	"fmt"
)
type SillyNephewError struct {
	message string
	details string
}
func (e *SillyNephewError) Error() string {
	return fmt.Sprintf("%s, %s", e.message, e.details)
  }
// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	availableFodder, err := weightFodder.FodderAmount()
	if cows == 0 {
		return 0.0, errors.New("division by zero")
	} else if cows < 0 {
		return 00, &SillyNephewError {
			message: "silly nephew",
			details: fmt.Sprintf("there cannot be %v cows", cows),
		}
	}
	if err == ErrScaleMalfunction {
		if availableFodder < 0 {
			return 0.0, errors.New("negative fodder")
		}
		availableFodder *= 2
	} else if err != nil {
		return 0.0, err
	}
	if availableFodder < 0 {
		return 0.0, errors.New("negative fodder")
	}
	return availableFodder / float64(cows), nil
}