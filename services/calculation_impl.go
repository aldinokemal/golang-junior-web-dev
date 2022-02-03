package services

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"lsp/entities"
	"lsp/repositories/contracts"
	"lsp/structs"
	"time"
)

type calculationImpl struct {
	PersegiRepo   contracts.PersegiRepository
	SegitigaRepo  contracts.SegitigaRepository
	LingkaranRepo contracts.LingkaranRepository
}

func NewCalculationService(persegiRepo *contracts.PersegiRepository, segitigaRepo *contracts.SegitigaRepository, lingkaranRepo *contracts.LingkaranRepository) CalculationService {
	return &calculationImpl{
		PersegiRepo:   *persegiRepo,
		SegitigaRepo:  *segitigaRepo,
		LingkaranRepo: *lingkaranRepo,
	}
}

func (s calculationImpl) CalculatePersegi(c *fiber.Ctx, request structs.CalculatePersegiRequest) (response structs.CalculatePersegiResponse, err error) {
	save, err := s.PersegiRepo.Save(entities.Persegi{
		Panjang:  request.Panjang,
		Lebar:    request.Lebar,
		Datetime: time.Now(),
		Hasil:    request.Panjang * request.Lebar,
	})
	if err != nil {
		return response, err
	}
	response.Panjang = save.Panjang
	response.Lebar = save.Lebar
	response.Hasil = save.Hasil
	response.Message = "Perhitungan persegi berhasil"
	return response, nil

}

func (s calculationImpl) CalculateLingkaran(c *fiber.Ctx, request structs.CalculateLingkaranRequest) (response structs.CalculateLingkaranResponse, err error) {
	save, err := s.LingkaranRepo.Save(entities.Lingkaran{
		Jari:     request.Jari,
		Datetime: time.Now(),
		Hasil:    3.14 * request.Jari * request.Jari,
	})
	if err != nil {
		return response, err
	}
	response.Jari = save.Jari
	response.Hasil = save.Hasil
	response.Message = "Perhitungan lingkaran berhasil"
	return response, nil
}

func (s calculationImpl) CalculateSegitiga(c *fiber.Ctx, request structs.CalculateSegitigaRequest) (response structs.CalculateSegitigaResponse, err error) {
	save, err := s.SegitigaRepo.Save(entities.Segitiga{
		Alas:     request.Alas,
		Tinggi:   request.Tinggi,
		Datetime: time.Now(),
		Hasil:    request.Alas * request.Tinggi / 2,
	})
	if err != nil {
		return response, err
	}
	response.Alas = save.Alas
	response.Tinggi = save.Tinggi
	response.Hasil = save.Hasil
	response.Message = "Perhitungan segitiga berhasil"
	return response, nil
}

func (s calculationImpl) HistoryPersegi(c *fiber.Ctx) (response structs.HistoryPersegi, err error) {
	data, err := s.PersegiRepo.FindAll()
	if err != nil {
		return response, err
	}
	response.Data = data
	response.Total = len(data)
	fmt.Println(data)
	return response, nil
}

func (s calculationImpl) HistoryLingkaran(c *fiber.Ctx) (response structs.HistoryLingkaran, err error) {
	data, err := s.LingkaranRepo.FindAll()
	if err != nil {
		return response, err
	}
	response.Data = data
	response.Total = len(data)
	return response, nil
}

func (s calculationImpl) HistorySegitiga(c *fiber.Ctx) (response structs.HistorySegitiga, err error) {
	data, err := s.SegitigaRepo.FindAll()
	if err != nil {
		return response, err
	}
	response.Data = data
	response.Total = len(data)
	return response, nil
}
