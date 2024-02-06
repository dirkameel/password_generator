## Secure Password Generator

A simple Go-based command-line tool for generating secure random passwords with customizable options.

### Installation

1. Ensure you have Go installed (version 1.21 or later)
2. Clone or download the source files
3. Build the application:
   ```bash
   go build -o password-generator
   ```

### Usage

```bash
./password-generator [options]
```

### Options

- `-length int`: Length of the password (default: 16)
- `-count int`: Number of passwords to generate (default: 1)
- `-upper bool`: Include uppercase letters (default: true)
- `-lower bool`: Include lowercase letters (default: true)
- `-digits bool`: Include digits (default: true)
- `-special bool`: Include special characters (default: true)
- `-no-similar bool`: Exclude similar characters (i, l, 1, L, o, 0, O) (default: false)
- `-no-ambiguous bool`: Exclude ambiguous characters ({ } [ ] ( ) / \ ' " ` ~ , ; : . < >) (default: false)

### Examples

1. Generate a single 16-character password with default settings:
   ```bash
   ./password-generator
   ```

2. Generate 5 passwords with 20 characters each:
   ```bash
   ./password-generator -length 20 -count 5
   ```

3. Generate a password with only letters (no digits or special characters):
   ```bash
   ./password-generator -digits=false -special=false
   ```

4. Generate a password excluding similar and ambiguous characters:
   ```bash
   ./password-generator -no-similar -no-ambiguous
   ```

5. Generate a numeric PIN code:
   ```bash
   ./password-generator -upper=false -lower=false -special=false -length 6
   ```

### Features

- **Cryptographically Secure**: Uses Go's `crypto/rand` package for secure random number generation
- **Customizable Character Sets**: Choose which character types to include
- **Exclusion Options**: Optionally exclude similar or ambiguous characters
- **Multiple Passwords**: Generate multiple passwords at once
- **Input Validation**: Validates configuration to ensure valid passwords can be generated
- **Character Set Enforcement**: Ensures at least one character from each selected character set is included

### Security Notes

- The generator uses cryptographically secure random number generation
- Passwords are generated with uniform distribution across the selected character sets
- The tool ensures that if you select multiple character sets, each generated password will contain at least one character from each selected set

### Error Handling

The tool provides helpful error messages for common issues:
- Invalid password length
- No character sets selected
- Password length exceeding available character combinations
- Random number generation failures