package validation

import (
	"errors"
	"gofs/internal/types"
	"net/mail"
	"regexp"
	"unicode/utf8"
)

const (
	allowedLength = 22
)

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil // false if ParseAddress Throws error
}

func isPasswordValid(password string) bool {
	// Some Patterns
	hasMinLen := regexp.MustCompile(`.{8,}`)
	hasUpper := regexp.MustCompile(`[A-Z]`)
	hasLower := regexp.MustCompile(`[a-z]`)
	hasNumber := regexp.MustCompile(`[0-9]`)
	hasSpecial := regexp.MustCompile(`[ !@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]`)

	if  hasMinLen.MatchString(password) &&
		hasUpper.MatchString(password) &&
		hasLower.MatchString(password) &&
		hasNumber.MatchString(password) &&
		hasSpecial.MatchString(password) {
		return true
	} else {
		return false
	}

}

// Get User Req--->Extract and validate ---> put into struct ---> return err or validated struct
func UserCreateRequestValidator(req types.UserCreateRequest) error {

	// Validating username
	if utf8.RuneCountInString(req.Username) > allowedLength {
		return errors.New("[Validation Error] Username exceeds allowed Length 22")
	}
	// Validating email
	if !isValidEmail(req.Email) {
		return errors.New("[Validation Error] Email is Invalid")
	}
	// Checking Passwword and Confirm Password
	if req.Password != req.ConfirmPassword {
		return errors.New("[Validation Error] Passsword and Confirm Password must be Same")
	}
	// Validating Password 
	if !isPasswordValid(req.Password){
		return errors.New("[Validation Error] Strong Password Is Required !")
	}

	return nil
}
