package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define command-line flags
	length := flag.Int("length", 16, "Length of the password")
	count := flag.Int("count", 1, "Number of passwords to generate")
	useUpper := flag.Bool("upper", true, "Include uppercase letters")
	useLower := flag.Bool("lower", true, "Include lowercase letters")
	useDigits := flag.Bool("digits", true, "Include digits")
	useSpecial := flag.Bool("special", true, "Include special characters")
	noSimilar := flag.Bool("no-similar", false, "Exclude similar characters (i, l, 1, L, o, 0, O)")
	noAmbiguous := flag.Bool("no-ambiguous", false, "Exclude ambiguous characters ({ } [ ] ( ) / \\ ' \" ` ~ , ; : . < >)")

	flag.Parse()

	// Validate input
	if *length < 1 {
		fmt.Println("Error: Password length must be at least 1")
		os.Exit(1)
	}

	if *count < 1 {
		fmt.Println("Error: Count must be at least 1")
		os.Exit(1)
	}

	// Check if at least one character set is selected
	if !*useUpper && !*useLower && !*useDigits && !*useSpecial {
		fmt.Println("Error: At least one character set must be selected")
		os.Exit(1)
	}

	// Generate passwords
	generator := NewPasswordGenerator()
	
	for i := 0; i < *count; i++ {
		password, err := generator.Generate(PasswordConfig{
			Length:      *length,
			UseUpper:    *useUpper,
			UseLower:    *useLower,
			UseDigits:   *useDigits,
			UseSpecial:  *useSpecial,
			NoSimilar:   *noSimilar,
			NoAmbiguous: *noAmbiguous,
		})
		
		if err != nil {
			fmt.Printf("Error generating password: %v\n", err)
			os.Exit(1)
		}
		
		fmt.Println(password)
	}
}