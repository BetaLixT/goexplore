package standard

import "github.com/betalixt/gorr"

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
