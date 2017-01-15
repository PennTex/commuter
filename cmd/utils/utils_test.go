package utils_test

import (
	"errors"

	"github.com/PennTex/commuter/cmd/utils"
)

// cant test this because process error exists the process
func ExampleProcessError() {
	err := errors.New("Test")
	es := "Error String"

	utils.ProcessError(err, es)
}
