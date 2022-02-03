package csv

import (
	"encoding/csv"
	"fmt"
	"lsp/entities"
	"lsp/repositories/contracts"
	"os"
	"strconv"
	"time"
)

const csvLingkaran = "./storages/lingkaran.csv"

type lingkaranImpl struct {
}

func NewLingkaranRepository() contracts.LingkaranRepository {
	return &lingkaranImpl{}
}

func (p lingkaranImpl) Save(entity entities.Lingkaran) (result entities.Lingkaran, err error) {
	csvFile, err := os.OpenFile(csvLingkaran, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	defer csvFile.Close()
	if err != nil {
		return entities.Lingkaran{}, err
	}

	csvwriter := csv.NewWriter(csvFile)
	defer csvwriter.Flush()
	dataLingkaran := []string{
		fmt.Sprintf("%f", entity.Jari),
		fmt.Sprintf("%f", entity.Hasil),
		fmt.Sprintf(entity.Datetime.Format("2006-01-02 15:04:05")),
	}
	err = csvwriter.Write(dataLingkaran)
	return entity, err
}

func (p lingkaranImpl) FindAll() (result []entities.Lingkaran, err error) {
	csvFile, err := os.OpenFile(csvLingkaran, os.O_RDONLY, 0644)
	if err != nil {
		return result, err
	}
	reader := csv.NewReader(csvFile)
	records, _ := reader.ReadAll()

	for _, record := range records {
		var lingkaran entities.Lingkaran
		lingkaran.Jari, err = strconv.ParseFloat(record[0], 8)
		if err != nil {
			return nil, err
		}
		lingkaran.Hasil, err = strconv.ParseFloat(record[1], 8)
		if err != nil {
			return nil, err
		}
		lingkaran.Datetime, err = time.Parse("2006-01-02 15:04:05", record[2])
		if err != nil {
			return nil, err
		}
		result = append(result, lingkaran)
	}

	return result, nil
}
