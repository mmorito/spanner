package helloworld

import (
	"github.com/mmorito/spanner/utilities/logger"
)

// UseCase struct
type UseCase func()

func (uc *UseCase) helloworld() (string, error) {
	log := logger.GetLogger()

	log.Infof("Hello")
	log.Warnf("World")
	return "hello world!", nil
}
