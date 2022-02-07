package message

import (
	"github.com/romberli/go-util/config"
)

const (
	InfoGenerateGetter = 200001
)

func initInfoMessage() {
	Messages[InfoGenerateGetter] = config.NewErrMessage(DefaultMessageHeader, InfoGenerateGetter, "generate getter completed successfully. output file: %s")
}
