# Learn Go - The Complete Guide

Learning Go from the [Udemy course: "Go - The Complete Guide"](https://www.udemy.com/course/go-the-complete-guide)

## 📚 Course Overview

This repository contains hands-on practice exercises for learning Go programming language. Each section covers different concepts and features of Go, progressing from basics to advanced topics like concurrency.

## 📂 Repository Structure

```
learn-go/
├── section2_1/              # Basic types and variables
├── section2_2-3/            # Variable types & file operations
├── section4/                # Control structures (pointers, if/else, loops)
├── section5_1/              # Structs basics
├── section5_2/              # Working with structs
├── section5_3-6_1/          # Interfaces & packages
│   ├── note/                # note package (Note struct, Save/Display methods)
│   ├── todo/                # todo package (Todo struct, Save/Display methods)
│   └── main.go              # Interface demonstration (saver, outputtable)
├── section6_2/              # Error handling
├── section7_1/              # Slices and maps
├── section7_2/              # List operations
├── section8_1/              # Functions as values, recursion, anonymous functions
├── section9_1/              # Concurrency & goroutines
│   ├── prices/              # Price calculation with tax rates
│   └── testdata/            # Test data files
└── section10_1/             # Goroutines & channel communication
```

## 🎯 Key Topics Covered

| Section | Topics |
|---------|--------|
| **2.1** | Basic types, variables, input/output |
| **2.2-3** | Type conversions, file operations |
| **4** | Pointers, control flow, loops |
| **5.1-2** | Structs, methods, embedding |
| **5.3-6** | **Interfaces**, packages, modular code |
| **6** | Error handling, custom errors |
| **7.1-2** | Arrays, slices, maps |
| **8** | Function values, recursion, closures |
| **9** | **Goroutines**, concurrency patterns |
| **10** | **Channels**, goroutine communication |

## 🚀 Running the Code

Each section is a standalone Go module. To run a specific section:

```bash
cd section5_3-6_1
go run main.go
```

Or build and run:

```bash
go build -o app
./app
```

## 📝 Notable Examples

### Interfaces (section5_3-6_1)
Demonstrates embedded interfaces and polymorphism:
```go
type saver interface {
    Save() error
}

type outputtable interface {
    saver
    Display()
}
```

Both `Note` and `Todo` implement these interfaces, allowing polymorphic behavior.

### Concurrency (section9_1)
Calculates prices with different tax rates concurrently:
```go
taxRates := []float64{0, 0.07, 0.1, 0.15}

for _, taxRate := range taxRates {
    priceJob := prices.NewTaxIncludedPriceJob(taxRate)
    priceJob.Process()
}
```

### Goroutines & Channels (section10_1)
Safe communication between goroutines using channels:
```go
done := make(chan bool)
go slowGreet("Hello!", done)
<-done  // Wait for completion
```

## 💡 Learning Tips

- Each section builds on previous concepts
- Run `go mod init` if starting a new module
- Use `go fmt` to format code
- Run `go vet` to check for potential issues
- Test with `go test` for test-driven development

## 📖 Course Link

[Go - The Complete Guide](https://www.udemy.com/course/go-the-complete-guide)

## 🛠️ Tools & Environment

- **Go Version:** 1.21.2+
- **Editor:** VS Code with Go extension
- **Build:** Standard Go toolchain

## ✅ Progress Tracking

- [x] Section 2: Types & Variables
- [x] Section 4: Control Structures
- [x] Section 5: Structs & Interfaces
- [x] Section 6: Error Handling
- [x] Section 7: Collections (Arrays, Slices, Maps)
- [x] Section 8: Advanced Functions
- [x] Section 9: Concurrency Basics
- [x] Section 10: Goroutines & Channels

---

**Last Updated:** October 2024