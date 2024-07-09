package valueobjects

import (
	"errors"
	re "regexp"
)

type Email struct {
	Value string
	Valid bool
}

func NewEmail(address string) (*Email, error) {
	emailAddress := Email{
		Value: address,
		Valid: false,
	}
	err := emailAddress.ValidateEmail()
	if err != nil {
		return nil, err
	}
	return &emailAddress, nil
}

func (e *Email) ValidateEmail() error {
	pattern := re.MustCompile(`.+@.+\..+`)
	if matches := pattern.Match([]byte(e.Value)); !matches {
		e.Valid = false
		return errors.New("Invalid email address")
	}
	e.Valid = true
	return nil
}

func (e *Email) MarshalJSON() ([]byte, error) {
	return []byte(e.Value), nil
}
