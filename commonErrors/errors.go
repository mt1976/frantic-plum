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
	return fmt.Errorf("string too long, max %d characters Error [%w]", ln, err)
}

func NotFoundError(err error) error {
	return fmt.Errorf("not found Error [%w]", err)
}
func ReadError(err error) error {
	return fmt.Errorf("read Error [%w]", err)
}
func WriteError(err error) error {
	return fmt.Errorf("write Error [%w]", err)
}
func EmptyError(err error) error {
	return fmt.Errorf("empty Error [%w]", err)
}
func ClearError(err error) error {
	return fmt.Errorf("clear Error [%w]", err)
}
func UpdateError(err error) error {
	return fmt.Errorf("update Error [%w]", err)
}
func CreateError(err error) error {
	return fmt.Errorf("create Error [%w]", err)
}
func DeleteError(err error) error {
	return fmt.Errorf("delete Error [%w]", err)
}
func DropError(err error) error {
	return fmt.Errorf("drop Error [%w]", err)
}
func ValidateError(err error) error {
	return fmt.Errorf("validate Error [%w]", err)
}
func DisconnectError(err error) error {
	return fmt.Errorf("disconnect Error [%w]", err)
}
func ConnectError(err error) error {
	return fmt.Errorf("connect Error [%w]", err)
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
	return fmt.Errorf("send email Error [%w]", err)
}
func IDGenerationError(err error) error {
	return fmt.Errorf("ID generation Error [%w]", err)
}

func OSError(err error) error {
	return fmt.Errorf("OS Error [%w]", err)
}

func MockingError(err error) error {
	return fmt.Errorf("Mocking Error [%w]", err)
}

func NotificationError(err error) error {
	return fmt.Errorf("Notification Error [%w]", err)
}
