package speech

type Speech struct {
	date   string
	origin string
}

func (s *Speech) saveSpeechToDB() bool {
	return true
}
