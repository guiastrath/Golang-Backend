package recognition

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"golang-backend/internal/middleware/auth"
	"golang-backend/pkg/common"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
)

type RecognitionService interface {
	Recognize(ctx context.Context, file []byte, params url.Values) (*RecognitionResponse, error)
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

func (s *recognitionService) Recognize(ctx context.Context, file []byte, params url.Values) (*RecognitionResponse, error) {

	// Building Recognition Body with file
	body, contentType, err := s.buildRecognitionBody(file)
	if err != nil {
		return nil, err
	}

	// Building URL with parameters
	baseURL, err := url.Parse(common.RecognitionURL)
	if err != nil {
		fmt.Println("Error on parsing base URL", err)
		return nil, err
	}
	parsedURL := baseURL.Query()

	for key, value := range params {
		parsedURL.Add(key, value[0])
	}

	baseURL.RawQuery = parsedURL.Encode()

	// Building Recognition Request
	recognitionRequest, err := http.NewRequest(http.MethodPost, baseURL.String(), body)
	if err != nil {
		fmt.Println("Error on creating recognition request", err)
		return nil, err
	}

	recognitionRequest.Header.Set("Content-Type", *contentType)
	recognitionRequest.Header.Set("x-api-key", ctx.Value(auth.API_KEY).(string))

	response, err := http.DefaultClient.Do(recognitionRequest)
	if err != nil {
		fmt.Println("Error on executing recognition request", err)
		return nil, err
	}
	defer response.Body.Close()

	// Getting response body
	decodedResponse, err := s.getResponseBody(response)
	if err != nil {
		return nil, err
	}

	return decodedResponse, nil
}

func (s *recognitionService) buildRecognitionBody(file []byte) (*bytes.Buffer, *string, error) {

	// Creating buffer for file
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", "temp.jpg")
	if err != nil {
		fmt.Println("Error on creating form file", err)
		return nil, nil, err
	}

	// Writing file to form field
	_, err = part.Write(file)
	if err != nil {
		fmt.Println("Error on writing file", err)
		return nil, nil, err
	}

	contentType := writer.FormDataContentType()

	// Closing writer
	err = writer.Close()
	if err != nil {
		fmt.Println("Error on closing writer", err)
		return nil, nil, err
	}

	return body, &contentType, nil
}

func (s *recognitionService) getResponseBody(response *http.Response) (*RecognitionResponse, error) {

	// Getting response body
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error on getting response body", err)
		return nil, err
	}

	// // Decoding response
	decodedResponse := &RecognitionResponse{}
	err = json.Unmarshal(responseBody, decodedResponse)
	if err != nil {
		fmt.Println("Error on unmarshaling JSON response", err)
		return nil, err
	}

	return decodedResponse, nil
}
