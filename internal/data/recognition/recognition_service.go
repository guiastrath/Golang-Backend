package recognition

type RecognitionService interface {
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
