- A panic in goroutine will also abruptly terminate main program
- Use errors.Join(..errS) to join multiple errors
- Use %w along with fmt.Errorf to wrap errors with additional errors/context. Use errors.Unwrap() to unwrap a wrapped error and get error embedded inside it

