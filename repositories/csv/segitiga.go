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

const csvSegitiga = "./storages/segitiga.csv"

type segitigaImpl struct {
}

func NewSegitigaRepository() contracts.SegitigaRepository {
	return &segitigaImpl{}
}

func (p segitigaImpl) Save(entity entities.Segitiga) (result entities.Segitiga, err error) {
	csvFile, err := os.OpenFile(csvSegitiga, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	defer csvFile.Close()
	if err != nil {
		return entities.Segitiga{}, err
	}

	csvwriter := csv.NewWriter(csvFile)
	dataSegitiga := []string{
		fmt.Sprintf("%f", entity.Alas),
		fmt.Sprintf("%f", entity.Tinggi),
		fmt.Sprintf("%f", entity.Hasil),
		fmt.Sprintf(entity.Datetime.Format("2006-01-02 15:04:05")),
	}

	err = csvwriter.Write(dataSegitiga)
	csvwriter.Flush()
	return entity, err
}

func (p segitigaImpl) FindAll() (result []entities.Segitiga, err error) {
	csvFile, err := os.OpenFile(csvSegitiga, os.O_RDONLY, 0644)
	if err != nil {
		return result, err
	}
	reader := csv.NewReader(csvFile)
	records, _ := reader.ReadAll()

	for _, record := range records {
		var segitiga entities.Segitiga
		segitiga.Alas, err = strconv.ParseFloat(record[0], 8)
		if err != nil {
			return nil, err
		}
		segitiga.Tinggi, err = strconv.ParseFloat(record[1], 8)
		if err != nil {
			return nil, err
		}
		segitiga.Hasil, err = strconv.ParseFloat(record[2], 8)
		if err != nil {
			return nil, err
		}
		segitiga.Datetime, err = time.Parse("2006-01-02 15:04:05", record[3])
		if err != nil {
			return nil, err
		}
		result = append(result, segitiga)
	}

	return result, nil
}
