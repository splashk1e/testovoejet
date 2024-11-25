package services

import (
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	licensev1 "github.com/splashk1e/jet/gen"
	mock_services "github.com/splashk1e/jet/internal/services/mocks"
)

func TestServerService_GetStatus(t *testing.T) {
	specificTime1 := time.Date(2023, 11, 19, 10, 0, 0, 0, time.UTC)
	license1 := &licensev1.License{
		Uid:           "98765",
		CreatedAt:     specificTime1.Unix(),
		UpdatedAt:     specificTime1.Unix(),
		CheckDate:     specificTime1.Unix(),
		RecheckDate:   specificTime1.Unix(),
		Worktime:      1800,
		Modules:       []licensev1.Module{licensev1.Module_T, licensev1.Module_B},
		Version:       "3.0.0",
		ReadOnly:      true,
		RecheckNeeded: false,
		WarningNotice: []*licensev1.WarningNotice{
			{Notice: "Warning 3"},
		},
		CriticalNotice: []*licensev1.CriticalNotice{
			{Notice: "Critical 3"},
		},
		Problems: []*licensev1.Problem{
			{Error: "Error 3", Date: specificTime1.Unix()},
		},
		MaxBasicConn:              150,
		MaxComplianceConn:         75,
		ConnSoftLimit:             true,
		ConnLimitExcess:           []int64{5, 10, 15},
		ComplianceConnLimitExcess: []int64{20, 25},
		PublicKey:                 "specific_public_key",
	}
	type mockBehavior func(s *mock_services.MockIService)
	testTable := []struct {
		name             string
		mockbehavior     mockBehavior
		expectedJsonData []byte
		expectedError    error
	}{
		{
			name: "OK",
			mockbehavior: func(s *mock_services.MockIService) {
				s.EXPECT().FileRead().Return(license1, nil)
			},
			expectedJsonData: []byte(`{"uid":"98765","check_date":1700388000,"modules":[2,0],"version":"3.0.0","read_only":true,"recheck_needed":false,"warning_notice":[{"notice":"Warning 3"}],"critical_notice":[{"notice":"Critical 3"}],"problems":[{"error":"Error 3","date":1700388000}],"max_basic_conn":150,"max_compliance_conn":75,"conn_soft_limit":true,"conn_limit_excess":[5,10,15],"compliance_conn_limit_excess":[20,25]}`),
			expectedError:    nil,
		},
		{
			name: "wrong protoclass",
			mockbehavior: func(s *mock_services.MockIService) {
				s.EXPECT().FileRead().Return(nil, nil)
			},
			expectedJsonData: nil,
			expectedError:    errors.New("wrong protoclass type"),
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			service := mock_services.NewMockIService(c)
			testCase.mockbehavior(service)
			serverService := ServerService{IService: service}

			jsonData, err := serverService.GetStatus()

			assert.Equal(t, jsonData, testCase.expectedJsonData)
			assert.Equal(t, err, testCase.expectedError)
		})
	}
}
