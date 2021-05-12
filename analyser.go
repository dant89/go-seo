package goseo

type Analyse struct {
	Report []Report
}

type Analyser interface {
	CheckH1Length()
}

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

func (a Analyse) CheckH1Length(h1 string) Report {
	report := Report{passed: true}

	if len(h1) < 20 {
		report.passed = false
		report = report.setFeedback(H1ShortError)
	}

	if len(h1) > 70 {
		report.passed = false
		report = report.setFeedback(H1LongError)
	}

	a.Report = append(a.Report, report)

	return report
}

func (r Report) Passed() bool {
	return r.passed
}

func (r Report) setFeedback(feedback string) Report {
	r.feedback = append(r.feedback, feedback)
	return r
}

func (r Report) GetFeedback() []string {
	return r.feedback
}
