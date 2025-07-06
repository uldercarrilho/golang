package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewTemperatureService(t *testing.T) {
	service := NewTemperatureService()
	assert.NotNil(t, service)
}

func TestConvertTemperature_CelsiusToFahrenheit(t *testing.T) {
	service := NewTemperatureService()

	req := &TemperatureConversionRequest{
		Value:    25.0,
		FromUnit: "celsius",
		ToUnit:   "fahrenheit",
	}

	resp, err := service.ConvertTemperature(req)
	require.NoError(t, err)

	assert.Equal(t, 25.0, resp.OriginalValue)
	assert.Equal(t, "celsius", resp.OriginalUnit)
	assert.Equal(t, 77.0, resp.ConvertedValue)
	assert.Equal(t, "fahrenheit", resp.ConvertedUnit)
	assert.Contains(t, resp.Formula, "°F = (°C × 9/5) + 32")
}

func TestConvertTemperature_CelsiusToKelvin(t *testing.T) {
	service := NewTemperatureService()

	req := &TemperatureConversionRequest{
		Value:    0.0,
		FromUnit: "celsius",
		ToUnit:   "kelvin",
	}

	resp, err := service.ConvertTemperature(req)
	require.NoError(t, err)

	assert.Equal(t, 0.0, resp.OriginalValue)
	assert.Equal(t, "celsius", resp.OriginalUnit)
	assert.Equal(t, 273.15, resp.ConvertedValue)
	assert.Equal(t, "kelvin", resp.ConvertedUnit)
	assert.Contains(t, resp.Formula, "K = °C + 273.15")
}

func TestConvertTemperature_FahrenheitToCelsius(t *testing.T) {
	service := NewTemperatureService()

	req := &TemperatureConversionRequest{
		Value:    212.0,
		FromUnit: "fahrenheit",
		ToUnit:   "celsius",
	}

	resp, err := service.ConvertTemperature(req)
	require.NoError(t, err)

	assert.Equal(t, 212.0, resp.OriginalValue)
	assert.Equal(t, "fahrenheit", resp.OriginalUnit)
	assert.Equal(t, 100.0, resp.ConvertedValue)
	assert.Equal(t, "celsius", resp.ConvertedUnit)
	assert.Contains(t, resp.Formula, "°C = (°F - 32) × 5/9")
}

func TestConvertTemperature_FahrenheitToKelvin(t *testing.T) {
	service := NewTemperatureService()

	req := &TemperatureConversionRequest{
		Value:    32.0,
		FromUnit: "fahrenheit",
		ToUnit:   "kelvin",
	}

	resp, err := service.ConvertTemperature(req)
	require.NoError(t, err)

	assert.Equal(t, 32.0, resp.OriginalValue)
	assert.Equal(t, "fahrenheit", resp.OriginalUnit)
	assert.Equal(t, 273.15, resp.ConvertedValue)
	assert.Equal(t, "kelvin", resp.ConvertedUnit)
	assert.Contains(t, resp.Formula, "K = (°F - 32) × 5/9 + 273.15")
}

func TestConvertTemperature_KelvinToCelsius(t *testing.T) {
	service := NewTemperatureService()

	req := &TemperatureConversionRequest{
		Value:    373.15,
		FromUnit: "kelvin",
		ToUnit:   "celsius",
	}

	resp, err := service.ConvertTemperature(req)
	require.NoError(t, err)

	assert.Equal(t, 373.15, resp.OriginalValue)
	assert.Equal(t, "kelvin", resp.OriginalUnit)
	assert.Equal(t, 100.0, resp.ConvertedValue)
	assert.Equal(t, "celsius", resp.ConvertedUnit)
	assert.Contains(t, resp.Formula, "°C = K - 273.15")
}

func TestConvertTemperature_KelvinToFahrenheit(t *testing.T) {
	service := NewTemperatureService()

	req := &TemperatureConversionRequest{
		Value:    273.15,
		FromUnit: "kelvin",
		ToUnit:   "fahrenheit",
	}

	resp, err := service.ConvertTemperature(req)
	require.NoError(t, err)

	assert.Equal(t, 273.15, resp.OriginalValue)
	assert.Equal(t, "kelvin", resp.OriginalUnit)
	assert.Equal(t, 32.0, resp.ConvertedValue)
	assert.Equal(t, "fahrenheit", resp.ConvertedUnit)
	assert.Contains(t, resp.Formula, "°F = (K - 273.15) × 9/5 + 32")
}

func TestConvertTemperature_SameUnit(t *testing.T) {
	service := NewTemperatureService()

	req := &TemperatureConversionRequest{
		Value:    25.0,
		FromUnit: "celsius",
		ToUnit:   "celsius",
	}

	resp, err := service.ConvertTemperature(req)
	require.NoError(t, err)

	assert.Equal(t, 25.0, resp.OriginalValue)
	assert.Equal(t, "celsius", resp.OriginalUnit)
	assert.Equal(t, 25.0, resp.ConvertedValue)
	assert.Equal(t, "celsius", resp.ConvertedUnit)
	assert.Equal(t, "Mesma unidade, sem conversão necessária", resp.Formula)
}

func TestConvertTemperature_DecimalPrecision(t *testing.T) {
	service := NewTemperatureService()

	req := &TemperatureConversionRequest{
		Value:    37.777777,
		FromUnit: "celsius",
		ToUnit:   "fahrenheit",
	}

	resp, err := service.ConvertTemperature(req)
	require.NoError(t, err)

	// Deve arredondar para 2 casas decimais
	assert.Equal(t, 100.0, resp.ConvertedValue)
}

func TestGetAllConversions_Celsius(t *testing.T) {
	service := NewTemperatureService()

	resp, err := service.GetAllConversions(25.0, "celsius")
	require.NoError(t, err)

	assert.Equal(t, 25.0, resp.OriginalValue)
	assert.Equal(t, "celsius", resp.OriginalUnit)

	// Verificar conversões
	assert.Equal(t, 25.0, resp.Conversions["celsius"])
	assert.Equal(t, 77.0, resp.Conversions["fahrenheit"])
	assert.Equal(t, 298.15, resp.Conversions["kelvin"])

	// Verificar fórmulas
	assert.Equal(t, "Mesma unidade, sem conversão necessária", resp.Formulas["celsius"])
	assert.Contains(t, resp.Formulas["fahrenheit"], "°F = (°C × 9/5) + 32")
	assert.Contains(t, resp.Formulas["kelvin"], "K = °C + 273.15")
}

func TestGetAllConversions_Fahrenheit(t *testing.T) {
	service := NewTemperatureService()

	resp, err := service.GetAllConversions(212.0, "fahrenheit")
	require.NoError(t, err)

	assert.Equal(t, 212.0, resp.OriginalValue)
	assert.Equal(t, "fahrenheit", resp.OriginalUnit)

	// Verificar conversões
	assert.Equal(t, 100.0, resp.Conversions["celsius"])
	assert.Equal(t, 212.0, resp.Conversions["fahrenheit"])
	assert.Equal(t, 373.15, resp.Conversions["kelvin"])
}

func TestGetAllConversions_Kelvin(t *testing.T) {
	service := NewTemperatureService()

	resp, err := service.GetAllConversions(273.15, "kelvin")
	require.NoError(t, err)

	assert.Equal(t, 273.15, resp.OriginalValue)
	assert.Equal(t, "kelvin", resp.OriginalUnit)

	// Verificar conversões
	assert.Equal(t, 0.0, resp.Conversions["celsius"])
	assert.Equal(t, 32.0, resp.Conversions["fahrenheit"])
	assert.Equal(t, 273.15, resp.Conversions["kelvin"])
}

// Testes de casos extremos
func TestConvertTemperature_ExtremeValues(t *testing.T) {
	service := NewTemperatureService()

	testCases := []struct {
		name     string
		value    float64
		fromUnit string
		toUnit   string
		expected float64
	}{
		{"Absolute Zero Celsius to Kelvin", -273.15, "celsius", "kelvin", 0.0},
		{"Absolute Zero Kelvin to Celsius", 0.0, "kelvin", "celsius", -273.15},
		{"Absolute Zero Celsius to Fahrenheit", -273.15, "celsius", "fahrenheit", -459.67},
		{"Absolute Zero Fahrenheit to Celsius", -459.67, "fahrenheit", "celsius", -273.15},
		{"Boiling Point Celsius to Fahrenheit", 100.0, "celsius", "fahrenheit", 212.0},
		{"Boiling Point Fahrenheit to Celsius", 212.0, "fahrenheit", "celsius", 100.0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := &TemperatureConversionRequest{
				Value:    tc.value,
				FromUnit: tc.fromUnit,
				ToUnit:   tc.toUnit,
			}

			resp, err := service.ConvertTemperature(req)
			require.NoError(t, err)

			assert.Equal(t, tc.expected, resp.ConvertedValue)
		})
	}
}
