package goseo

type Report struct {
	passed   bool
	feedback []string
}

type Reporter interface {
	AddFeedback()
	GetFeedback()
}

const H1ShortError = "The H1 is too short, aim for 20 characters minimum."
const H1LongError = "The H1 is too long, aim for 70 characters maximum."

// TODO find a better method for analyser functions
// Possibly a struct per validation type with
// interface grouping functionality?
func CheckH1Length(h1 string) Report {
	report := Report{passed: true}

	if len(h1) < 20 {
		report.passed = false
		report.AddFeedback(H1ShortError)
	}

	if len(h1) > 70 {
		report.passed = false
		report.AddFeedback(H1LongError)
	}

	return report
}

func (r Report) Passed() bool {
	return r.passed
}

func (r Report) AddFeedback(feedback string) {
	r.feedback = append(r.feedback, feedback)
}

func (r Report) GetFeedback() []string {
	return r.feedback
}
