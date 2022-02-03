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

const csvPersegi = "./storages/persegi.csv"

type persegiImpl struct {
}

func NewPersegiRepository() contracts.PersegiRepository {
	return &persegiImpl{}
}

func (p persegiImpl) Save(entity entities.Persegi) (result entities.Persegi, err error) {
	csvFile, err := os.OpenFile(csvPersegi, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	defer csvFile.Close()
	if err != nil {
		return entities.Persegi{}, err
	}

	csvwriter := csv.NewWriter(csvFile)
	defer csvwriter.Flush()
	dataPersegi := []string{
		fmt.Sprintf("%f", entity.Lebar),
		fmt.Sprintf("%f", entity.Panjang),
		fmt.Sprintf("%f", entity.Hasil),
		entity.Datetime.Format("2006-01-02 15:04:05"),
	}
	err = csvwriter.Write(dataPersegi)
	return entity, err
}

func (p persegiImpl) FindAll() (result []entities.Persegi, err error) {
	csvFile, err := os.OpenFile(csvPersegi, os.O_RDONLY, 0777)
	if err != nil {
		return result, err
	}
	reader := csv.NewReader(csvFile)
	records, err := reader.ReadAll()
	if err != nil {
		return result, err
	}
	fmt.Println(records)

	for _, record := range records {
		var persegi entities.Persegi
		persegi.Lebar, err = strconv.ParseFloat(record[0], 8)
		if err != nil {
			return nil, err
		}
		persegi.Panjang, err = strconv.ParseFloat(record[1], 8)
		if err != nil {
			return nil, err
		}
		persegi.Hasil, err = strconv.ParseFloat(record[2], 8)
		if err != nil {
			return nil, err
		}
		persegi.Datetime, err = time.Parse("2006-01-02 15:04:05", record[3])
		if err != nil {
			return nil, err
		}
		result = append(result, persegi)
	}

	return result, nil
}
