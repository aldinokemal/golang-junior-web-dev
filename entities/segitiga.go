package entities

import "time"

type Segitiga struct {
	Alas     float64   `json:"alas" csv:"alas"`
	Tinggi   float64   `json:"tinggi" csv:"tinggi"`
	Datetime time.Time `json:"datetime" csv:"datetime"`
	Hasil    float64   `json:"hasil" csv:"hasil"`
}
