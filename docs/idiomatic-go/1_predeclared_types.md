- All predeclared types are assigned defauklt values when declared but not initialised. 0 for numeric types, "" for strings, false for bool, nil for pointers, slices, maps, fucntions, interfaces, channels. Any struct will have all its memebres initialised with their default values
- Be careful when using `:=` to declare vars. A variable that is declared using `:=` same Name in different blocks/scopes will not override the original valye even though they have same names   
  ```go
  func main() {
    x := 3
    if x > 8 {
        x := 8
    }
    fmt.Print(x) // will still print 3 because x was redeclared inside if block and will get different address 
  }
  ```
- A rune literal(alias for int32) reprsents a unicdoe character enclosed in single quotes.(Single & double quotes are not interchangeable). Rune represented by unicdoe character in single quotes('😀', 'a'), 8bit octal number ('\141'), 8bit hexadecimal number ('\x678'), 16bit hexadecimal number('\u0061'), 32bit unicode number('\U00016')
- strings are interpreted literals. They contain zero or more rune literals. Always enclosed in double quotes or backticks(where they are interpreted as is without bothering escape sequences)
  - what this vbasically means string is a slice of bytes or slice of rune. You can typecase vice-versa
  - in slice of bytes, we go across each byte representation of unicdoe characters. this ejmans unicode whi is 4bytes long cannot be simply indexed using slices of bytes. iterating over string is same as iterating over slice of bytes. when you print you will get byte value (0-255)
  - iterating over string is same as iterating over slice of rune
  - string(byte) will give you the unicoede character. so do not confuse when string(12) doesnt give me '12' but unicode characrter represented by 12
  - however when you typecast string as slice of rune, and ietrate you will get each unicode character4 codepoint ie; unicode utf-8 number. When you typecast this back to string you will get unicode character. Directly printing the rune, will simply give you the unicode number not character
  - that's why strings are interpreted literals. you interpret what character is behind the rune literal
- Never use == and != to compare two float numbers. All mathemaical operators are allowed between floats except %
- slices cont support == and !=(basically not comparable). However arrays are comparable
- clear function takes a slice and sets all its elements to its zero value 
- hash maps are not comparable. keys for hashmap should be comparable
- structs instances comparision can only be comparable if its members are comparable, otherwise not - compiler will complain
- functions, channels, maps, slices and structs that contain before types are not comparable 
- named return values can be used with assigning to them explicity   
  ```go
  func Foo(arg int) (value int, err error) {
      if arg == 0 {
        return value, err
      }
      return 0, errors.New("") // this will automatically assign them to named return vars and finally those named vars are returned
  }
  ```