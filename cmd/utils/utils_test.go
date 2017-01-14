package utils_test

import (
	"errors"
	"testing"

	"github.com/PennTex/commuter/cmd/utils"
)

func setup() {

}

func TestProcessError(t *testing.T) {
	err := errors.New("Test")
	es := "Error String"

	utils.ProcessError(err, es)
	// Output: Error String: Test
}
