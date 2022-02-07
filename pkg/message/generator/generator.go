package generator

import (
	"github.com/romberli/go-generator/pkg/message"
	"github.com/romberli/go-util/config"
)

const (
	ErrGenerateGetter = 401001
)

func init() {
	initErrorMessage()
}

func initErrorMessage() {
	message.Messages[ErrGenerateGetter] = config.NewErrMessage(message.DefaultMessageHeader, ErrGenerateGetter, "generate getter failed")
}
