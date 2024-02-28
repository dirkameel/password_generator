## Secure Password Generator

A simple Go-based command-line tool for generating secure random passwords with customizable options.

### Installation

1. Ensure you have Go installed (version 1.21 or later)
2. Clone or download the source files
3. Build the application:
   ```bash
   go build -o password-gen
   ```

### Usage

```bash
./password-gen [options]
```

### Command Line Options

- `-length int`: Password length (default: 12)
- `-count int`: Number of passwords to generate (default: 1)
- `-upper`: Include uppercase letters (default: true)
- `-lower`: Include lowercase letters (default: true)
- `-digits`: Include digits (default: true)
- `-special`: Include special characters (default: false)
- `-no-similar`: Exclude similar characters (i, l, 1, L, o, 0, O) (default: false)
- `-no-ambiguous`: Exclude ambiguous characters ({ } [ ] ( ) / \ ' " ` ~ , ; : . < >) (default: false)

### Examples

1. Generate a simple 12-character password:
   ```bash
   ./password-gen
   ```

2. Generate a 16-character password with special characters:
   ```bash
   ./password-gen -length 16 -special
   ```

3. Generate 5 passwords without similar characters:
   ```bash
   ./password-gen -count 5 -no-similar
   ```

4. Generate a numeric-only password:
   ```bash
   ./password-gen -upper=false -lower=false -digits -special=false -length 8
   ```

5. Generate a complex password for system use:
   ```bash
   ./password-gen -length 20 -special -no-similar -no-ambiguous
   ```

### Features

- **Cryptographically Secure**: Uses Go's `crypto/rand` package for secure random number generation
- **Customizable Character Sets**: Choose which character types to include
- **Exclusion Options**: Remove similar or ambiguous characters to improve readability
- **Batch Generation**: Generate multiple passwords at once
- **Input Validation**: Comprehensive error checking and validation

### Security Notes

- The generator uses cryptographically secure random number generation
- Passwords are generated with uniform distribution across the selected character set
- No passwords are stored or logged by the application
- Consider using longer passwords (16+ characters) for critical applications

### Requirements

- Go 1.21 or later
- No external dependencies required