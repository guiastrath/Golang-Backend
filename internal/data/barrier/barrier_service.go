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
	ControlBarrier(sessionInfo *auth.AuthData, result *recognition.RecognitionResponse) (bool, error)
}

type barrierService struct {
}

func NewRecognitionService() BarrierService {
	return &barrierService{}
}

func (s *barrierService) ConfigureBarrier(deviceId string, pulse string, pulseWidth int) error {
	return nil
}

func (s *barrierService) ControlBarrier(sessionInfo *auth.AuthData, result *recognition.RecognitionResponse) (bool, error) {

	// Validating Recognition
	if !s.validateSimilarity(result) {
		return true, fmt.Errorf("no valid recognition")
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
		return false, err
	}

	body := bytes.NewBuffer(jsonData)

	// Building request
	controlRequest, err := http.NewRequest(http.MethodPost, common.BarrierControlUrl, body)
	if err != nil {
		fmt.Println("Error on creating barrierControl request", err)
		return false, err
	}

	controlRequest.Header.Set("Content-Type", common.JSON)

	// Executing request
	if common.TestBarrierPresent {
		response, err := http.DefaultClient.Do(controlRequest)
		if err != nil {
			fmt.Println("Error on executing barrierControl request", err)
			return false, err
		}
		defer response.Body.Close()
	}

	return true, nil
}

func (h *barrierService) validateSimilarity(result *recognition.RecognitionResponse) bool {
	for _, res := range result.Result {
		for _, subject := range res.Subjects {
			fmt.Printf("%s recognized with %.2f%% similarity\n", subject.Subject, subject.Similarity*100)
			if subject.Similarity >= common.MinimumSimilarity {
				return true
			}
		}
	}
	return false
}
