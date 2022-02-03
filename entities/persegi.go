package entities

import "time"

type Persegi struct {
	Panjang  float64   `json:"panjang" csv:"panjang"`
	Lebar    float64   `json:"lebar" csv:"lebar"`
	Datetime time.Time `json:"datetime" csv:"datetime"`
	Hasil    float64   `json:"hasil" csv:"hasil"`
}
