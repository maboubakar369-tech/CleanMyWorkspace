package CleanMyWorkspace

func GenererateWorkSpace() *[][]*string {
	souk := [][]*string{
		{nil, ptr("Emballage de fast-food"), ptr("Mouchoir en papier"), nil},
		{ptr("Canette de soda"), ptr("Goblet"), nil, ptr("Paquet de confiserie")},
		{nil, ptr("Reste de nuggets"), ptr("Papier de brouillon"), nil},
	}
	return &souk
}

func ptr(s string) *string {
	return &s
}
