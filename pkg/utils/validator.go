package utils

import (
	"regexp"
	"strings"
)

// Validator contém funções de validação comuns.
type Validator struct{}

// NewValidator cria uma nova instância do validator.
func NewValidator() *Validator {
	return &Validator{}
}

// IsValidEmail valida se um email é válido.
func (v *Validator) IsValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

// IsValidPassword valida se uma senha é válida.
func (v *Validator) IsValidPassword(password string) bool {
	// Mínimo 8 caracteres, pelo menos uma letra maiúscula, uma minúscula e um número
	if len(password) < 8 {
		return false
	}

	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)

	return hasUpper && hasLower && hasNumber
}

// IsValidUUID valida se uma string é um UUID válido.
func (v *Validator) IsValidUUID(uuid string) bool {
	uuidRegex := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)
	return uuidRegex.MatchString(strings.ToLower(uuid))
}

// IsValidPhone valida se um telefone é válido (formato brasileiro).
func (v *Validator) IsValidPhone(phone string) bool {
	// Remove caracteres não numéricos
	phone = regexp.MustCompile(`[^0-9]`).ReplaceAllString(phone, "")

	// Verifica se tem 10 ou 11 dígitos (com DDD)
	return len(phone) >= 10 && len(phone) <= 11
}

// IsValidCPF valida se um CPF é válido.
func (v *Validator) IsValidCPF(cpf string) bool {
	// Remove caracteres não numéricos
	cpf = regexp.MustCompile(`[^0-9]`).ReplaceAllString(cpf, "")

	// Verifica se tem 11 dígitos
	if len(cpf) != 11 {
		return false
	}

	// Verifica se todos os dígitos são iguais
	if cpf == "00000000000" || cpf == "11111111111" || cpf == "22222222222" ||
		cpf == "33333333333" || cpf == "44444444444" || cpf == "55555555555" ||
		cpf == "66666666666" || cpf == "77777777777" || cpf == "88888888888" ||
		cpf == "99999999999" {
		return false
	}

	// Validação dos dígitos verificadores
	return v.validateCPFDigits(cpf)
}

// validateCPFDigits valida os dígitos verificadores do CPF.
func (v *Validator) validateCPFDigits(cpf string) bool {
	// Primeiro dígito verificador
	sum := 0
	for i := 0; i < 9; i++ {
		sum += int(cpf[i]-'0') * (10 - i)
	}
	remainder := sum % 11
	digit1 := 0
	if remainder >= 2 {
		digit1 = 11 - remainder
	}

	// Segundo dígito verificador
	sum = 0
	for i := 0; i < 10; i++ {
		sum += int(cpf[i]-'0') * (11 - i)
	}
	remainder = sum % 11
	digit2 := 0
	if remainder >= 2 {
		digit2 = 11 - remainder
	}

	// Verifica se os dígitos calculados são iguais aos do CPF
	return int(cpf[9]-'0') == digit1 && int(cpf[10]-'0') == digit2
}
