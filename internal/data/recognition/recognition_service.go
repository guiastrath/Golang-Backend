package recognition

import (
	"context"
	"fmt"
	"golang-backend/internal/middleware/auth"
	"golang-backend/pkg/common"
	"io"
	"net/http"
)

type RecognitionService interface {
	Recognize(ctx context.Context, requestParams io.ReadCloser) (io.ReadCloser, error)
}

type recognitionService struct {
	recognitionRepository RecognitionRepository
}

func NewRecognitionService() RecognitionService {
	recognitionRepository := NewRecognitionRepository()

	return &recognitionService{
		recognitionRepository: recognitionRepository,
	}
}

func (s *recognitionService) Recognize(ctx context.Context, requestParams io.ReadCloser) (io.ReadCloser, error) {

	// Building Recognition Request
	recognitionRequest, err := http.NewRequest(http.MethodPost, common.RecognitionURL, requestParams)
	if err != nil {
		fmt.Println("Error on creating recognition request", err)
		return nil, err
	}

	recognitionRequest.Header.Add("Content-Type", common.RecognitionContentType)
	recognitionRequest.Header.Add("x-api-key", ctx.Value(auth.API_KEY).(string))

	response, err := http.DefaultClient.Do(recognitionRequest)
	if err != nil {
		fmt.Println("Error on executing recognition request", err)
		return nil, err
	}
	defer response.Body.Close()

	// Treating Response Data
	// body, err := io.ReadAll(response.Body)
	// if err != nil {
	// 	fmt.Println("Error on executing recognition request", err)
	// 	return nil, err
	// }

	// var result *string

	// err = json.Unmarshal(body, result)
	// if err != nil {
	// 	fmt.Println("Error unmarshalling recognition response", err)
	// 	return nil, err
	// }

	return response.Body, nil
}
