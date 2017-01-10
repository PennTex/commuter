package utils_test

import "github.com/PennTex/commuter/cmd/utils"

func ExampleLogger_Log_withLogging() {
	var logger utils.Logger
	logger.Logging = true

	logger.Log("hi")
	// Output: hi
}

func ExampleLogger_Log_withoutLogging() {
	var logger utils.Logger
	logger.Logging = false

	logger.Log("hi")
	// Output:
}
