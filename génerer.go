package main

func GénérerWorkSpace() *[][]*string {
	souk := [][]*string{
		{ptr("stylo"), ptr("crayon"), nil},
		{ptr("papier"), nil, ptr("livre")},
		{nil, ptr("ordinateur"), ptr("tasse")},
	}
	return &souk
}

func ptr(s string) *string {
	return &s
}
