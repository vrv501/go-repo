All literal types are not having any default types. There are 4 types: integer, decimal, rune, string
### Integer literals
- Only integer literals support bitwise, logical operators
- support all arithmetic, comparision operators
- byte is alias for uint8 and rune for int32

### Decimal literals
- Decimal literals represented using float32, float64(default) can be used with all arithmetic operators except %.You can also use comparsion operators
- Dont support bitwie & logical operators
- Division of floating wih 0 either yields +Inf or -Inf(if float is < 0). Divinding float32/64(0) by 0 yeilds NaN
- enerally while comparig floats its recommonded to calculate variance and see if it's less than threshod rather than directly comparing floats given their limited precision

### Rune and string literals
- Strings are interpreted literals. They interpret rune
- You can use all comparision operators
- Strings are immutable 
- Strings can **only** be concatenated using +

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

### Strings
- They are interpreted literals of slice of runes. Each rune in inturn represented by slice of bytes. So len(string) will give total number of bytes in string, not rune
- When you use index on a string, you are grabbing byte at that index not rune. 
- When you slice string, you're slicing slice of bytes which means if rune which is represented by more than one byte gets sliced then there are chances it will be incorrectly cut somewhere in between which might not be final expected result. Each index represents one byte
`[255 123 45 6 89]`
- You can typecast rune, byte as string(var). Similarly you can type-cast string as []rune(str) or []byte(str)
- for-range loop loops over rune in strings. So if you use `for i, v := range str`, the v is rune while i is starting index of rune. so dont be surprised if i is not consecutive as you loop over. You might need to type-cast string(v) to get string representation