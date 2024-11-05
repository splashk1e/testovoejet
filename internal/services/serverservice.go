package services

import (
	"encoding/json"
	"errors"

	licensev1 "github.com/splashk1e/jet/gen"
	"github.com/splashk1e/jet/internal/config"
)

type ServerService struct {
	Service
}

func (s *ServerService) GetStatus() ([]byte, error) {
	protoclass, err := s.FileRead()
	if err != nil {
		return nil, err
	}
	license, ok := protoclass.(*licensev1.License)
	if !ok {
		return nil, errors.New("wrong protoclass type")
	}
	jsonData, err := json.Marshal(struct {
		Uid                       string                      `json:"uid"`
		CheckDate                 int64                       `json:"chech_date"`
		Modules                   []licensev1.Module          `json:"modules"`
		Version                   string                      `json:"version"`
		ReadOnly                  bool                        `json:"read_only"`
		RecheckNeeded             bool                        `json:"recheck_needed"`
		WarningNotice             []*licensev1.WarningNotice  `json:"warning_notice"`
		CriticalNotice            []*licensev1.CriticalNotice ` json:"critical_notice"`
		Problems                  []*licensev1.Problem        `json:"problems"`
		MaxBasicConn              int32                       `json:"max_basic_conn"`
		MaxComplianceConn         int32                       `json:"max_compliance_conn"`
		ConnSoftLimit             bool                        `json:"conn_soft_limit"`
		ConnLimitExcess           []int64                     ` json:"conn_limit_excess"`
		ComplianceConnLimitExcess []int64                     `json:"compliance_conn_limit_excess"`
	}{license.Uid, license.CheckDate, license.Modules, license.Version, license.ReadOnly, license.RecheckNeeded, license.WarningNotice, license.CriticalNotice, license.Problems, license.MaxBasicConn, license.MaxComplianceConn, license.ConnSoftLimit, license.ConnLimitExcess, license.ComplianceConnLimitExcess})
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}
func NewServerService(cfg config.Config) *ServerService {
	return &ServerService{
		Service: *NewService(cfg),
	}
}
