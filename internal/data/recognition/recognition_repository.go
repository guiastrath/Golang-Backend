package recognition

type RecognitionRepository interface {
}

type recognitionRepository struct {
}

func NewRecognitionRepository() RecognitionRepository {
	return &recognitionRepository{}
}
