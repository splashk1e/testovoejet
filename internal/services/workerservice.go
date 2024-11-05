package services

import (
	"errors"

	licensev1 "github.com/splashk1e/jet/gen"
	"github.com/splashk1e/jet/internal/config"
)

type WorkerService struct {
	*Service
}

func (s *WorkerService) FileUpdate() error {
	protoclass, err := s.FileRead()
	if err != nil {
		return err
	}
	license, ok := protoclass.(*licensev1.License)
	if !ok {
		return errors.New("wrong protoclass type")
	}
	license.WarningNotice = nil
	license.CriticalNotice = nil
	license.ReadOnly = false
	s.FileWrite(license)
	return nil
}

func NewWorkerService(cfg config.Config) *WorkerService {
	return &WorkerService{
		Service: NewService(cfg),
	}
}
