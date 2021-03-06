package common

import "github.com/betalixt/gorr"

func NewForcastConflictError() *gorr.Error {
  return gorr.NewError(
    gorr.ErrorCode{
      Code: 1100,
      Message: "ForcastConflictError",
    },
    409,
    "",
  )
}

func NewForcastMissingError() *gorr.Error {
  return gorr.NewError(
    gorr.ErrorCode{
      Code: 1101,
      Message: "ForcastMissingError",
    },
    404,
    "",
  )
}

func NewBodyReadingError(err error) *gorr.Error {
  return gorr.NewError(
    gorr.ErrorCode{
      Code: 10000,
      Message: "BodyReadingError",
    },
    400,
    err.Error(),
  )
}
