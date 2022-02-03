package contracts

import "lsp/entities"

type PersegiRepository interface {
	Save(entity entities.Persegi) (result entities.Persegi, err error)
	FindAll() (result []entities.Persegi, err error)
}
