package goseo

type Report struct {
	Passed   bool
	Feedback []string
}

const H1ShortError = "The H1 is too short, aim for 20 characters minimum."
const H1LongError = "The H1 is too long, aim for 70 characters maximum."

func CheckH1Length(h1 string) Report {
	report := Report{Passed: true}

	if len(h1) < 20 {
		report.Passed = false
		report.Feedback = append(report.Feedback, H1ShortError)
	}

	if len(h1) > 70 {
		report.Passed = false
		report.Feedback = append(report.Feedback, H1LongError)
	}

	return report
}
