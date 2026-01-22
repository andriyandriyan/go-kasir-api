package models

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var Categories = []Category{
	{
		ID:          1,
		Name:        "Analgisik",
		Description: "Obat pereda nyeri atau penghilang rasa sakit",
	},
	{
		ID:          2,
		Name:        "Antibiotik",
		Description: "Obat yang digunakan untuk mencegah dan mengobati infeksi bakteri",
	},
	{
		ID:          3,
		Name:        "Antihistamin",
		Description: "Obat yang digunakan untuk meredakan gejala alergi",
	},
}
