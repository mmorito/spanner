package helloworld

import (
	"github.com/mmorito/spanner/src/utilities/logger"
)

// UseCase struct
type UseCase func()

func (uc *UseCase) helloworld() (string, error) {

	logger.Log.Infof("Hello")
	logger.Log.Warnf("World")
	return "hello world!", nil
}
