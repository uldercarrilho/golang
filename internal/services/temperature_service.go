package services

import (
	"fmt"
	"math"
)

// TemperatureService fornece funcionalidades para conversão de temperatura
type TemperatureService struct{}

// NewTemperatureService cria uma nova instância do serviço de temperatura
func NewTemperatureService() *TemperatureService {
	return &TemperatureService{}
}

// TemperatureConversionRequest representa a requisição de conversão
type TemperatureConversionRequest struct {
	Value    float64 `json:"value" binding:"required"`
	FromUnit string  `json:"from_unit" binding:"required,oneof=kelvin celsius fahrenheit"`
	ToUnit   string  `json:"to_unit" binding:"required,oneof=kelvin celsius fahrenheit"`
}

// TemperatureConversionResponse representa a resposta de conversão
type TemperatureConversionResponse struct {
	OriginalValue  float64 `json:"original_value"`
	OriginalUnit   string  `json:"original_unit"`
	ConvertedValue float64 `json:"converted_value"`
	ConvertedUnit  string  `json:"converted_unit"`
	Formula        string  `json:"formula"`
}

// ConvertTemperature converte uma temperatura de uma unidade para outra
func (s *TemperatureService) ConvertTemperature(req *TemperatureConversionRequest) (*TemperatureConversionResponse, error) {
	var convertedValue float64
	var formula string

	switch req.FromUnit {
	case "celsius":
		switch req.ToUnit {
		case "fahrenheit":
			convertedValue = (req.Value * 9 / 5) + 32
			formula = fmt.Sprintf("°F = (°C × 9/5) + 32 = (%.2f × 9/5) + 32 = %.2f", req.Value, convertedValue)
		case "kelvin":
			convertedValue = req.Value + 273.15
			formula = fmt.Sprintf("K = °C + 273.15 = %.2f + 273.15 = %.2f", req.Value, convertedValue)
		case "celsius":
			convertedValue = req.Value
			formula = "Mesma unidade, sem conversão necessária"
		}
	case "fahrenheit":
		switch req.ToUnit {
		case "celsius":
			convertedValue = (req.Value - 32) * 5 / 9
			formula = fmt.Sprintf("°C = (°F - 32) × 5/9 = (%.2f - 32) × 5/9 = %.2f", req.Value, convertedValue)
		case "kelvin":
			convertedValue = (req.Value-32)*5/9 + 273.15
			formula = fmt.Sprintf("K = (°F - 32) × 5/9 + 273.15 = (%.2f - 32) × 5/9 + 273.15 = %.2f", req.Value, convertedValue)
		case "fahrenheit":
			convertedValue = req.Value
			formula = "Mesma unidade, sem conversão necessária"
		}
	case "kelvin":
		switch req.ToUnit {
		case "celsius":
			convertedValue = req.Value - 273.15
			formula = fmt.Sprintf("°C = K - 273.15 = %.2f - 273.15 = %.2f", req.Value, convertedValue)
		case "fahrenheit":
			convertedValue = (req.Value-273.15)*9/5 + 32
			formula = fmt.Sprintf("°F = (K - 273.15) × 9/5 + 32 = (%.2f - 273.15) × 9/5 + 32 = %.2f", req.Value, convertedValue)
		case "kelvin":
			convertedValue = req.Value
			formula = "Mesma unidade, sem conversão necessária"
		}
	}

	// Arredondar para 2 casas decimais
	convertedValue = math.Round(convertedValue*100) / 100

	return &TemperatureConversionResponse{
		OriginalValue:  req.Value,
		OriginalUnit:   req.FromUnit,
		ConvertedValue: convertedValue,
		ConvertedUnit:  req.ToUnit,
		Formula:        formula,
	}, nil
}

// GetAllConversions retorna todas as conversões possíveis para um valor dado
type AllConversionsResponse struct {
	OriginalValue float64            `json:"original_value"`
	OriginalUnit  string             `json:"original_unit"`
	Conversions   map[string]float64 `json:"conversions"`
	Formulas      map[string]string  `json:"formulas"`
}

// GetAllConversions converte um valor para todas as unidades disponíveis
func (s *TemperatureService) GetAllConversions(value float64, fromUnit string) (*AllConversionsResponse, error) {
	conversions := make(map[string]float64)
	formulas := make(map[string]string)

	units := []string{"celsius", "fahrenheit", "kelvin"}

	for _, unit := range units {
		if unit == fromUnit {
			conversions[unit] = value
			formulas[unit] = "Mesma unidade, sem conversão necessária"
			continue
		}

		req := &TemperatureConversionRequest{
			Value:    value,
			FromUnit: fromUnit,
			ToUnit:   unit,
		}

		resp, err := s.ConvertTemperature(req)
		if err != nil {
			return nil, err
		}

		conversions[unit] = resp.ConvertedValue
		formulas[unit] = resp.Formula
	}

	return &AllConversionsResponse{
		OriginalValue: value,
		OriginalUnit:  fromUnit,
		Conversions:   conversions,
		Formulas:      formulas,
	}, nil
}
