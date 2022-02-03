package structs

import "lsp/entities"

// CalculatePersegiRequest : parameter yang dikirim untuk menghitung persegi
type CalculatePersegiRequest struct {
	Panjang float64 `json:"panjang"`
	Lebar   float64 `json:"lebar"`
}

// CalculatePersegiResponse : parameter yang dikirim untuk response persegi
type CalculatePersegiResponse struct {
	Panjang float64 `json:"panjang"`
	Lebar   float64 `json:"lebar"`
	Hasil   float64 `json:"hasil"`
	Message string  `json:"message"`
}

// CalculateLingkaranRequest : parameter yang dikirim untuk menghitung lingkaran
type CalculateLingkaranRequest struct {
	Jari float64 `json:"jari"`
}

// CalculateLingkaranResponse : parameter yang dikirim untuk response lingkaran
type CalculateLingkaranResponse struct {
	Jari    float64 `json:"jari"`
	Hasil   float64 `json:"hasil"`
	Message string  `json:"message"`
}

// CalculateSegitigaRequest : parameter yang dikirim untuk menghitung segitiga
type CalculateSegitigaRequest struct {
	Alas   float64 `json:"alas"`
	Tinggi float64 `json:"tinggi"`
}

// CalculateSegitigaResponse : parameter yang dikirim untuk response segitiga
type CalculateSegitigaResponse struct {
	Alas    float64 `json:"alas"`
	Tinggi  float64 `json:"tinggi"`
	Hasil   float64 `json:"hasil"`
	Message string  `json:"message"`
}

//================================================================

type HistoryPersegi struct {
	Data  []entities.Persegi `json:"data"`
	Total int                `json:"total"`
}
type HistorySegitiga struct {
	Data  []entities.Segitiga `json:"data"`
	Total int                 `json:"total"`
}
type HistoryLingkaran struct {
	Data  []entities.Lingkaran `json:"data"`
	Total int                  `json:"total"`
}
