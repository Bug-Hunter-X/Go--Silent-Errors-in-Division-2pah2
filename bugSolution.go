func Calculate(a, b int) (int, error) {
  if b == 0 {
    return 0, errors.New("division by zero")
  }
  return a / b, nil
}

func main() {
  result, err := Calculate(10, 0)
  if err != nil {
    fmt.Println("Error:", err) // Print error details
    return // Exit early if error occurs
  }
  fmt.Println("Result:", result)
} 
