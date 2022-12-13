package errorunwrap

import (
	"fmt"

	"github.com/betalixt/gorr"
)

func main() {
	err := gorr.NewError(
		gorr.ErrorCode{
			Code:    123,
			Message: "AnError",
		},
		"",
	)
	werr := fmt.Errorf("wrapping error: %w")
}
