package barrier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang-backend/internal/data/recognition"
	"golang-backend/internal/middleware/auth"
	"golang-backend/pkg/common"
	"net/http"
)

type BarrierService interface {
	ConfigureBarrier(deviceId string, pulse string, pulseWidth int) error
	ControlBarrier(sessionInfo *auth.AuthData, result *recognition.RecognitionResponse) (*recognition.RecognitionDisplay, error)
}

type barrierService struct {
}

func NewRecognitionService() BarrierService {
	return &barrierService{}
}

func (s *barrierService) ConfigureBarrier(deviceId string, pulse string, pulseWidth int) error {
	return nil
}

func (s *barrierService) ControlBarrier(sessionInfo *auth.AuthData, result *recognition.RecognitionResponse) (*recognition.RecognitionDisplay, error) {

	// Validating Recognition
	validatedFace := s.validateSimilarity(result)
	if validatedFace == nil {
		return nil, nil
	}

	// Building request struct
	controlBody := BarrierControl{
		DeviceID: sessionInfo.DeviceID,
		Data: BarrierControlData{
			Switch: common.ON,
		},
	}

	// Building request Body
	jsonData, err := json.Marshal(controlBody)
	if err != nil {
		fmt.Println("Error on creating barrierControl body - JSON Marshal", err)
		return nil, err
	}

	body := bytes.NewBuffer(jsonData)

	// Building request
	controlRequest, err := http.NewRequest(http.MethodPost, common.BarrierControlUrl, body)
	if err != nil {
		fmt.Println("Error on creating barrierControl request", err)
		return nil, err
	}

	controlRequest.Header.Set("Content-Type", common.JSON)

	// Executing request
	if common.TestBarrierPresent {
		response, err := http.DefaultClient.Do(controlRequest)
		if err != nil {
			fmt.Println("Error on executing barrierControl request", err)
			return nil, err
		}
		defer response.Body.Close()
	}

	return validatedFace, nil
}

func (h *barrierService) validateSimilarity(result *recognition.RecognitionResponse) *recognition.RecognitionDisplay {
	for _, res := range result.Result {
		for _, subject := range res.Subjects {
			if subject.Similarity >= common.MinimumSimilarity {
				return &recognition.RecognitionDisplay{
					Box:     res.Box,
					Subject: subject,
				}
			}
		}
	}
	return nil
}
