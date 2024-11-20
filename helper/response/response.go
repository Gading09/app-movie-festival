package response

import (
	"fmt"
	"strconv"
	"strings"

	"movie-festival/helper/constant"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data"`
}

func ResponseOK(c *fiber.Ctx, statusCode int, msg string, data interface{}) error {
	response := Response{
		Status:  constant.Success,
		Message: msg,
		Data:    data,
	}

	return c.Status(statusCode).JSON(response)
}

func TrimMesssage(err error) (statusCode int, customError, originalError string) {
	fmt.Println(err.Error())
	errs := strings.Split(err.Error(), "|")
	if len(errs) == 1 {
		return constant.StatusInternalServerError, constant.ErrDefault, "Something went wrong"
	}
	statusCode, _ = strconv.Atoi(strings.TrimSpace(errs[0]))
	customError = strings.TrimSpace(errs[1])
	originalError = strings.TrimSpace(errs[2])
	return
}

func ResponseError(c *fiber.Ctx, err error) error {
	statusCode, errType, originalError := TrimMesssage(err)

	if errType == constant.ErrValidator {
		errType = originalError
	}

	response := Response{
		Status:  constant.Error,
		Message: errType,
		Data:    nil,
	}

	return c.Status(statusCode).JSON(response)
}
