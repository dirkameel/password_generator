package main

import (
	"crypto/rand"
	"errors"
	"math/big"
)

// PasswordConfig holds the configuration for password generation
type PasswordConfig struct {
	Length      int
	UseUpper    bool
	UseLower    bool
	UseDigits   bool
	UseSpecial  bool
	NoSimilar   bool
	NoAmbiguous bool
}

// PasswordGenerator handles password generation
type PasswordGenerator struct {
	charPool string
}

// NewPasswordGenerator creates a new password generator instance
func NewPasswordGenerator() *PasswordGenerator {
	return &PasswordGenerator{}
}

// Generate creates a secure random password based on the provided configuration
func (pg *PasswordGenerator) Generate(config PasswordConfig) (string, error) {
	// Build character pool based on configuration
	pg.buildCharPool(config)
	
	if pg.charPool == "" {
		return "", errors.New("no character sets selected for password generation")
	}

	if config.Length < 1 {
		return "", errors.New("password length must be at least 1")
	}

	// Generate password
	password := make([]byte, config.Length)
	for i := 0; i < config.Length; i++ {
		char, err := pg.getRandomChar()
		if err != nil {
			return "", err
		}
		password[i] = char
	}

	return string(password), nil
}

// buildCharPool constructs the character pool based on configuration
func (pg *PasswordGenerator) buildCharPool(config PasswordConfig) {
	var charPool string

	// Define character sets
	upperCase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowerCase := "abcdefghijklmnopqrstuvwxyz"
	digits := "0123456789"
	specialChars := "!@#$%^&*()_+-=[]{}|;:,.<>?"

	// Remove similar characters if requested
	if config.NoSimilar {
		upperCase = removeChars(upperCase, "ILO")
		lowerCase = removeChars(lowerCase, "ilo")
		digits = removeChars(digits, "01")
	}

	// Remove ambiguous characters if requested
	if config.NoAmbiguous {
		specialChars = removeChars(specialChars, "{}[]()/\\'\"`~,;:.<>")
	}

	// Add character sets to pool based on configuration
	if config.UseUpper {
		charPool += upperCase
	}
	if config.UseLower {
		charPool += lowerCase
	}
	if config.UseDigits {
		charPool += digits
	}
	if config.UseSpecial {
		charPool += specialChars
	}

	pg.charPool = charPool
}

// getRandomChar returns a random character from the character pool
func (pg *PasswordGenerator) getRandomChar() (byte, error) {
	if pg.charPool == "" {
		return 0, errors.New("character pool is empty")
	}

	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(pg.charPool))))
	if err != nil {
		return 0, err
	}

	return pg.charPool[n.Int64()], nil
}

// removeChars removes specified characters from a string
func removeChars(source, charsToRemove string) string {
	result := source
	for _, char := range charsToRemove {
		result = removeChar(result, byte(char))
	}
	return result
}

// removeChar removes all occurrences of a character from a string
func removeChar(source string, charToRemove byte) string {
	var result []byte
	for i := 0; i < len(source); i++ {
		if source[i] != charToRemove {
			result = append(result, source[i])
		}
	}
	return string(result)
}