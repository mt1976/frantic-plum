package commonErrors

import (
	"errors"
	"fmt"

	"github.com/mt1976/frantic-core/logger"
)

var (
	ErrorEndDateBeforeStartDate = errors.New("end date is before start date")
	ErrorEmptyName              = errors.New("name is empty")
	ErrorNameTooLong            = errors.New("name is too long, max 50 characters") // Deprecated: use StringTooLongError
	ErrorDuplicate              = errors.New("duplicate")
	ErrorNegativeValue          = errors.New("negative value")
	//ErrorNotFound               = errors.New("not found %w %w") // Deprecated: use NotFoundError
	ErrorPasswordMismatch   = errors.New("password mismatch")
	ErrorUserNotFound       = errors.New("user not found")
	ErrorUserNotActive      = errors.New("user not active")
	ErrNoTranslation        = errors.New("no translation available")
	ErrNoMessageToTranslate = errors.New("no message to translate")
	ErrProtocolIsRequired   = errors.New("protocol is required")
	ErrInvalidProtocol      = errors.New("invalid protocol")
	ErrHostIsRequired       = errors.New("host is required")
	ErrInvalidHost          = errors.New("invalid host")
	ErrPortIsRequired       = errors.New("port is required")
	ErrInvalidPort          = errors.New("invalid port")
	ErrUsernameIsRequired   = errors.New("username is required")
	ErrInvalidUsername      = errors.New("invalid username")
	ErrPasswordIsRequired   = errors.New("password is required")
	ErrInvalidPassword      = errors.New("invalid password")
	ErrOriginIsRequired     = errors.New("no origin defined, and origin identifier is required")
	ErrInvalidOrigin        = errors.New("invalid origin")
)

func WrapStringTooLongErr(err error, ln int) error {
	return fmt.Errorf("string too long, max %d characters error (%w)", ln, err)
}

func WrapNotFoundError(err error) error {
	return fmt.Errorf("not found error (%w)", err)
}
func WrapReadError(err error) error {
	return fmt.Errorf("read error (%w)", err)
}
func WrapWriteError(err error) error {
	return fmt.Errorf("write error (%w)", err)
}
func WrapEmptyError(err error) error {
	return fmt.Errorf("empty error (%w)", err)
}
func WrapClearError(err error) error {
	return fmt.Errorf("clear error (%w)", err)
}
func WrapUpdateError(err error) error {
	return fmt.Errorf("update error (%w)", err)
}
func WrapCreateError(err error) error {
	return fmt.Errorf("create error (%w)", err)
}
func WrapDeleteError(err error) error {
	return fmt.Errorf("delete error (%w)", err)
}
func WrapDropError(err error) error {
	return fmt.Errorf("drop error (%w)", err)
}
func WrapValidationError(err error) error {
	return fmt.Errorf("validate error (%w)", err)
}
func WrapDisconnectError(err error) error {
	return fmt.Errorf("disconnect error (%w)", err)
}
func WrapConnectError(err error) error {
	return fmt.Errorf("connect error (%w)", err)
}
func HandleGoValidatorError(err error) error {
	return nil
	// if err != nil {

	// 	if _, ok := err.(*validator.InvalidValidationError); ok {
	// 		logger.InfoLogger.Println(err)
	// 		return err
	// 	}

	// 	for _, err := range err.(validator.ValidationErrors) {

	// 		op := fmt.Sprintf("VALIDATION: Field[%s] Tag[%s] Kind[%s] Param[%s] Value[%s]", err.Field(), err.Tag(), err.Kind(), err.Param(), err.Value())
	// 		logger.InfoLogger.Println(op)

	// 	}

	// 	return err
	// }
	// return nil
}
func WrapEmailError(err error) error {
	return fmt.Errorf("send email error (%w)", err)
}
func WrapIDGenerationError(err error) error {
	return fmt.Errorf("ID generation error (%w)", err)
}

func WrapOSError(err error) error {
	return fmt.Errorf("OS error (%w)", err)
}

func WrapErrorForMocking(err error) error {
	return fmt.Errorf("mocking error (%w)", err)
}

func WrapNotificationError(err error) error {
	return fmt.Errorf("notification error (%w)", err)
}

func WrapFunctionalError(err error, f string) error {
	return fmt.Errorf("functional error - %v (%w)", f, err)
}

func WrapError(err error) error {
	logger.WarningLogger.Println("It is not advised to wrap errors without a specific error message")
	return fmt.Errorf("error (%w)", err)
}

func WrapInvalidFilterError(err error, f string) error {
	return fmt.Errorf("invalid filter [%v] (%w)", f, err)
}

func WrapInvalidHttpReturnStatusError(s string) error {
	return fmt.Errorf("inavalid/unsupported http return status [%v]", s)
}

func WrapInvalidHttpReturnStatusWithMessageError(status, message string) error {
	return fmt.Errorf("inavalid/unsupported http return status [%v] (%v)", status, message)
}
