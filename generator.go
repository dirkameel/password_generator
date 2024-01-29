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
	charPools map[string]string
}

// NewPasswordGenerator creates a new password generator instance
func NewPasswordGenerator() *PasswordGenerator {
	return &PasswordGenerator{
		charPools: map[string]string{
			"upper":   "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			"lower":   "abcdefghijklmnopqrstuvwxyz",
			"digits":  "0123456789",
			"special": "!@#$%^&*()_+-=[]{}|;:,.<>?",
		},
	}
}

// Generate creates a secure random password based on the configuration
func (pg *PasswordGenerator) Generate(config PasswordConfig) (string, error) {
	// Build character pool based on configuration
	charPool := pg.buildCharacterPool(config)
	
	if len(charPool) == 0 {
		return "", errors.New("no character sets selected for password generation")
	}

	if config.Length > len(charPool) {
		return "", errors.New("password length exceeds available character combinations")
	}

	// Generate password
	password := make([]byte, config.Length)
	
	for i := 0; i < config.Length; i++ {
		char, err := pg.getRandomChar(charPool)
		if err != nil {
			return "", err
		}
		password[i] = char
	}

	// Ensure at least one character from each selected set is included
	err := pg.ensureCharacterSets(&password, config, charPool)
	if err != nil {
		return "", err
	}

	return string(password), nil
}

// buildCharacterPool constructs the character pool based on configuration
func (pg *PasswordGenerator) buildCharacterPool(config PasswordConfig) string {
	var pool string
	
	// Start with base character sets
	basePools := make(map[string]string)
	for key, chars := range pg.charPools {
		basePools[key] = chars
	}

	// Apply filters
	if config.NoSimilar {
		basePools["upper"] = removeCharacters(basePools["upper"], "ILO")
		basePools["lower"] = removeCharacters(basePools["lower"], "ilo")
		basePools["digits"] = removeCharacters(basePools["digits"], "01")
	}

	if config.NoAmbiguous {
		basePools["special"] = removeCharacters(basePools["special"], "{}[]()/\\'\"`~,;:.<>")
	}

	// Add selected character sets to pool
	if config.UseUpper {
		pool += basePools["upper"]
	}
	if config.UseLower {
		pool += basePools["lower"]
	}
	if config.UseDigits {
		pool += basePools["digits"]
	}
	if config.UseSpecial {
		pool += basePools["special"]
	}

	return pool
}

// getRandomChar returns a random character from the pool
func (pg *PasswordGenerator) getRandomChar(pool string) (byte, error) {
	if len(pool) == 0 {
		return 0, errors.New("character pool is empty")
	}
	
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(pool))))
	if err != nil {
		return 0, err
	}
	
	return pool[n.Int64()], nil
}

// ensureCharacterSets ensures at least one character from each selected set is included
func (pg *PasswordGenerator) ensureCharacterSets(password *[]byte, config PasswordConfig, fullPool string) error {
	requiredSets := make([]string, 0)
	
	if config.UseUpper {
		requiredSets = append(requiredSets, "upper")
	}
	if config.UseLower {
		requiredSets = append(requiredSets, "lower")
	}
	if config.UseDigits {
		requiredSets = append(requiredSets, "digits")
	}
	if config.UseSpecial {
		requiredSets = append(requiredSets, "special")
	}

	// Build filtered character sets
	charSets := make(map[string]string)
	
	upper := pg.charPools["upper"]
	lower := pg.charPools["lower"]
	digits := pg.charPools["digits"]
	special := pg.charPools["special"]
	
	if config.NoSimilar {
		upper = removeCharacters(upper, "ILO")
		lower = removeCharacters(lower, "ilo")
		digits = removeCharacters(digits, "01")
	}
	if config.NoAmbiguous {
		special = removeCharacters(special, "{}[]()/\\'\"`~,;:.<>")
	}
	
	charSets["upper"] = upper
	charSets["lower"] = lower
	charSets["digits"] = digits
	charSets["special"] = special

	// Check and fix missing character sets
	for _, set := range requiredSets {
		if !pg.containsCharFromSet(string(*password), charSets[set]) {
			// Replace a random position with a character from the missing set
			pos, err := pg.getRandomPosition(len(*password))
			if err != nil {
				return err
			}
			
			char, err := pg.getRandomChar(charSets[set])
			if err != nil {
				return err
			}
			
			(*password)[pos] = char
		}
	}
	
	return nil
}

// containsCharFromSet checks if the password contains at least one character from the given set
func (pg *PasswordGenerator) containsCharFromSet(password, charSet string) bool {
	for _, passChar := range password {
		for _, setChar := range charSet {
			if passChar == setChar {
				return true
			}
		}
	}
	return false
}

// getRandomPosition returns a random position in the password
func (pg *PasswordGenerator) getRandomPosition(length int) (int, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(length)))
	if err != nil {
		return 0, err
	}
	return int(n.Int64()), nil
}

// removeCharacters removes specified characters from a string
func removeCharacters(source, charsToRemove string) string {
	result := source
	for _, char := range charsToRemove {
		result = removeChar(result, byte(char))
	}
	return result
}

// removeChar removes all occurrences of a character from a string
func removeChar(source string, char byte) string {
	var result []byte
	for i := 0; i < len(source); i++ {
		if source[i] != char {
			result = append(result, source[i])
		}
	}
	return string(result)
}