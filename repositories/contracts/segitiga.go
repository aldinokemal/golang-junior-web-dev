package contracts

import "lsp/entities"

type SegitigaRepository interface {
	Save(entity entities.Segitiga) (result entities.Segitiga, err error)
	FindAll() (result []entities.Segitiga, err error)
}
