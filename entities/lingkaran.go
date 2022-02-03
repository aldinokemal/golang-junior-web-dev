package entities

import "time"

type Lingkaran struct {
	Jari     float64   `json:"jari" csv:"jari"`
	Datetime time.Time `json:"datetime" csv:"datetime"`
	Hasil    float64   `json:"hasil" csv:"hasil"`
}
