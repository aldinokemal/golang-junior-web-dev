package contracts

import "lsp/entities"

type LingkaranRepository interface {
	Save(entity entities.Lingkaran) (result entities.Lingkaran, err error)
	FindAll() (result []entities.Lingkaran, err error)
}
