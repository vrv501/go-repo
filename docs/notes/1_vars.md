- Any variable that exists outside functions must use var keyword followed by one or more variableName(s) followed by optional dataType(otherwise implicit decision is taken based on values assigned - not recommonded)  
  ```go
  package main
  
  var (
    x int32
    y = 98.2 // float64
    z int = 78 //int64 - int is 32 bits on x86 & 64bits on x64
  )


  func main() {
    var x, y, z int
  }
  ```

- Inside functions, you can follow the above convention or use `:=` shorthand notation without var keyword. Again this would decide dataType implicitly
  ```go
  func main() {
    var x,
    x := "sdfsdf" // string
    y := 45 //int64
  }
  ```

- byte is alias for uint8, rune also called as Unicode code point is alias for int32
- Explicit type-casting is required in go, you cannot perform operations on different dataTypes. However take care when explicit type-casting as if number exceeds max range of type-casting dataType it will be rounded back(can sometimes result in logical errors)
- const can be used to declare constants just like vars, but cannot be initialised with `:=`
- Any functionName, variable outside the function beginning with a capitalLetter is considered exported or package level variable