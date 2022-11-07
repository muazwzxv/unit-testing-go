package error_formats

import (
	"fmt"
	"strings"
	"unit_test/utils/error_utils"

	"github.com/jackc/pgx/v5/pgconn"
)

func ParseError(err error) error_utils.MessageErr {
	sqlErr, ok := err.(*pgconn.PgError)
	if !ok {
		if strings.Contains(err.Error(), "no rows in result set") {
			return error_utils.NewNotFoundError("no record matching given id")
		}
		return error_utils.NewInternalServerError(fmt.Sprintf("error when trying to save message: %s", err.Error()))
	}

	return error_utils.NewInternalServerError(fmt.Sprintf("error when processing request: %s", sqlErr.Message))
}
