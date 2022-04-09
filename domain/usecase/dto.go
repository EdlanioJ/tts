package usecase

// InputTextToSpeech is the input struct for the TextToSpeech use case.
type InputTextToSpeech struct {
	Language string
	Text     string
}

// OutputTextToSpeech is the output value for the TextToSpeech use case.
type OutputTextToSpeech []byte
