package main

func GénérerWorkSpace() *[][]*string {
	souk := [][]*string{
		{nil, ptr("Emballage de fast-food"), nil, nil, nil, ptr("Mouchoir en papier"), nil, nil, nil, nil, nil, ptr("Canette de soda"), ptr("Goblet"), nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil},
		{nil, nil, ptr("Canette de soda"), nil, ptr("Paquet de confiserie"), nil, nil, ptr("Goblet"), nil, nil, ptr("Paquet de confiserie"), nil, nil, nil, nil, nil, ptr("Emballage de fast-food"), nil, ptr("Goblet"), nil, ptr("Goblet"), nil, nil, ptr("Goblet"), nil, ptr("Paquet de confiserie"), ptr("Emballage de fast-food"), nil, nil, nil, nil, nil, nil, nil, nil, nil},
		{nil, nil, nil, ptr("Goblet"), ptr("Mouchoir en papier"), nil, nil, nil, ptr("Papier de brouillon"), nil, ptr("Mouchoir en papier"), nil, nil, nil, nil, nil, ptr("Reste de nuggets"), nil, nil, nil, nil, nil, ptr("Canette de soda"), nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil},
		{ptr("Papier de brouillon"), nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, ptr("Paquet de confiserie"), nil, nil, ptr("Canette de soda"), ptr("Paquet de confiserie"), nil, nil, ptr("Goblet"), ptr("Goblet"), nil, ptr("Emballage de fast-food"), nil, nil, nil, nil, ptr("Emballage de sandwich"), ptr("Mouchoir en papier"), nil, nil, nil, nil, ptr("Mouchoir en papier"), nil, nil},
		{nil, nil, nil, nil, ptr("Emballage de sandwich"), nil, nil, nil, nil, nil, nil, ptr("Canette de soda"), nil, ptr("Paquet de confiserie"), nil, ptr("Goblet"), nil, nil, nil, nil, nil, nil, nil, ptr("Reste de nuggets"), nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, ptr("Emballage de fast-food")},
		// etc. 
	}
	return &souk
}

func ptr(s string) *string {
	return &s
}
