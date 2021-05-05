package goseo

func CheckH1Length(h1 string) string {
	if len(h1) < 20 {
		return "The H1 is too short, aim for 20 characters minimum."
	}

	if len(h1) > 70 {
		return "The H1 is too long, aim for 70 characters maximum."
	}

	return "The H1 is perfect."
}
