package models

import "errors"

type RegistrationError error
type LoginError error
type CommonDatabaseError error

var (
	NoUsernameError = CommonDatabaseError(errors.New("The username cannot be an empty field."))
	NoPasswordError = CommonDatabaseError(errors.New("The password cannot be an empty field."))
	NoEmailError = CommonDatabaseError(errors.New("Email cannot be empty."))

	NotFoundAccountError = LoginError(errors.New("We did not found your account. Is your email written correctly?"))
	AccessDeniedError = LoginError(errors.New("Access Denied. Incorrect email or password."))

	ShortUsernameError = RegistrationError(errors.New("The username needs to have 4 characters or more."))
	LongUsernameError = RegistrationError(errors.New("The username needs to have 20 characters or less."))
	DuplicatedUsernameError = RegistrationError(errors.New("This username is already token."))

	WeakPasswordError = RegistrationError(errors.New("The password is too weak. We recommend 8 characters or more."))
	LongPasswordError = RegistrationError(errors.New("The password needs to have 31 characters or less."))
	IsPasswordEncryptedError = RegistrationError(errors.New("Is the password encrypted?"))

	ShortEmailError = RegistrationError(errors.New("Email needs to be longer. It needs to be 12 characters or longer."))
	LongEmailError = RegistrationError(errors.New("Email needs to be shorter. Using 40 characters as max value."))
	WrongEmailFormatError = RegistrationError(errors.New("Email format is not correct."))
	DuplicatedEmailError = RegistrationError(errors.New("This email is already registered."))
)
