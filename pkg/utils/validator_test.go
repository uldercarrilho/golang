package utils

import (
	"testing"
)

func TestValidator_IsValidEmail(t *testing.T) {
	v := NewValidator()

	tests := []struct {
		name     string
		email    string
		expected bool
	}{
		{"valid email", "test@example.com", true},
		{"valid email with subdomain", "test@sub.example.com", true},
		{"invalid email - no @", "testexample.com", false},
		{"invalid email - no domain", "test@", false},
		{"invalid email - no local part", "@example.com", false},
		{"empty email", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := v.IsValidEmail(tt.email)
			if result != tt.expected {
				t.Errorf("IsValidEmail(%s) = %v, want %v", tt.email, result, tt.expected)
			}
		})
	}
}

func TestValidator_IsValidPassword(t *testing.T) {
	v := NewValidator()

	tests := []struct {
		name     string
		password string
		expected bool
	}{
		{"valid password", "Password123", true},
		{"valid password with special chars", "P@ssw0rd", true},
		{"too short", "Pass1", false},
		{"no uppercase", "password123", false},
		{"no lowercase", "PASSWORD123", false},
		{"no number", "Password", false},
		{"empty password", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := v.IsValidPassword(tt.password)
			if result != tt.expected {
				t.Errorf("IsValidPassword(%s) = %v, want %v", tt.password, result, tt.expected)
			}
		})
	}
}

func TestValidator_IsValidUUID(t *testing.T) {
	v := NewValidator()

	tests := []struct {
		name     string
		uuid     string
		expected bool
	}{
		{"valid UUID", "550e8400-e29b-41d4-a716-446655440000", true},
		{"valid UUID uppercase", "550E8400-E29B-41D4-A716-446655440000", true},
		{"invalid UUID - wrong format", "550e8400-e29b-41d4-a716-44665544000", false},
		{"invalid UUID - wrong characters", "550e8400-e29b-41d4-a716-44665544000g", false},
		{"empty UUID", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := v.IsValidUUID(tt.uuid)
			if result != tt.expected {
				t.Errorf("IsValidUUID(%s) = %v, want %v", tt.uuid, result, tt.expected)
			}
		})
	}
}

func TestValidator_IsValidPhone(t *testing.T) {
	v := NewValidator()

	tests := []struct {
		name     string
		phone    string
		expected bool
	}{
		{"valid phone with DDD", "11987654321", true},
		{"valid phone with formatting", "(11) 98765-4321", true},
		{"valid phone with spaces", "11 98765 4321", true},
		{"valid phone 10 digits", "1198765432", true},
		{"too short", "119876543", false},
		{"too long", "119876543210", false},
		{"empty phone", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := v.IsValidPhone(tt.phone)
			if result != tt.expected {
				t.Errorf("IsValidPhone(%s) = %v, want %v", tt.phone, result, tt.expected)
			}
		})
	}
}

func TestValidator_IsValidCPF(t *testing.T) {
	v := NewValidator()

	tests := []struct {
		name     string
		cpf      string
		expected bool
	}{
		{"valid CPF", "12345678909", true},
		{"valid CPF with formatting", "123.456.789-09", true},
		{"invalid CPF - all same digits", "11111111111", false},
		{"invalid CPF - wrong length", "1234567890", false},
		{"invalid CPF - wrong digits", "12345678900", false},
		{"empty CPF", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := v.IsValidCPF(tt.cpf)
			if result != tt.expected {
				t.Errorf("IsValidCPF(%s) = %v, want %v", tt.cpf, result, tt.expected)
			}
		})
	}
}
