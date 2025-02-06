package commonErrors

import (
	"errors"
	"fmt"
)

var (
	ErrorEndDateBeforeStartDate = errors.New("end date is before start date")
	ErrorEmptyName              = errors.New("name is empty")
	ErrorNameTooLong            = errors.New("name is too long, max 50 characters") // Deprecated: use StringTooLongError
	ErrorDuplicate              = errors.New("duplicate")
	ErrorNegativeValue          = errors.New("negative value")
	//ErrorNotFound               = errors.New("not found %w %w") // Deprecated: use NotFoundError
	ErrorPasswordMismatch = errors.New("password mismatch")
)

func StringTooLongError(err error, ln int) error {
	return fmt.Errorf("string too long, max %d characters error [%w]", ln, err)
}

func NotFoundError(err error) error {
	return fmt.Errorf("not found error [%w]", err)
}
func ReadError(err error) error {
	return fmt.Errorf("read error [%w]", err)
}
func WriteError(err error) error {
	return fmt.Errorf("write error [%w]", err)
}
func EmptyError(err error) error {
	return fmt.Errorf("empty error [%w]", err)
}
func ClearError(err error) error {
	return fmt.Errorf("clear error [%w]", err)
}
func UpdateError(err error) error {
	return fmt.Errorf("update error [%w]", err)
}
func CreateError(err error) error {
	return fmt.Errorf("create error [%w]", err)
}
func DeleteError(err error) error {
	return fmt.Errorf("delete error [%w]", err)
}
func DropError(err error) error {
	return fmt.Errorf("drop error [%w]", err)
}
func ValidateError(err error) error {
	return fmt.Errorf("validate error [%w]", err)
}
func DisconnectError(err error) error {
	return fmt.Errorf("disconnect error [%w]", err)
}
func ConnectError(err error) error {
	return fmt.Errorf("connect error [%w]", err)
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
func SendEmailError(err error) error {
	return fmt.Errorf("send email error [%w]", err)
}
func IDGenerationError(err error) error {
	return fmt.Errorf("ID generation error [%w]", err)
}

func OSError(err error) error {
	return fmt.Errorf("OS error [%w]", err)
}

func MockingError(err error) error {
	return fmt.Errorf("mocking error [%w]", err)
}

func NotificationError(err error) error {
	return fmt.Errorf("notification error [%w]", err)
}

func FunctionalError(err error, f string) error {
	return fmt.Errorf("functional error - %v [%w]", f, err)
}

func GeneralError(err error) error {
	return fmt.Errorf("general error [%w]", err)
}
