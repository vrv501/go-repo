- A method is a special function with reciever argument. in between func and functionName   
  ```
  func(a myType) functinName(){}
  ```
- Any method has to be defined on a type. type can be anything that's not part of stdlib or other packages or that's not defined(has to be xplicity defined using type not inline(ex: map[int]any))
  ```go
  type myInt int
  type Circle struct{}
  ```
  In the above example you can create methods for myInt, Circle but not for int, struct
- Keep in mind if you declare your onw type on an inbuilt type, it doesn't inherit the original type's methods - no inheritance in go 😄
- If you're type is implementing a interface and you want to ensure it always implements interface and accidentally prevent issues in runtime, you can use below to handle errors at compile time   
  ```go
  var _ InterfaceName = (*typeName)(nil) // assuming *typeName implements interface
  ```
- An interface underneath is a tuple holding (value, concrete-type)   
  ```go
  var i interface = TypeName{} // Here i will technically be ({}, TypeName)
  ```
- Empty interface `interface{}` basically means no methods. So any type implements an empty interface{}. This can be usefukl when your fucntion expects arguments of any dataType
  ```go
  var i interface{} = "34234"
  fmt.Println(i)
  i = 56
  fmt.Println(i)
  j, ok := i.(int) // use type assertions to get the underlying value if interface is of type .(type)
                   // always use booleans
  // type assertion switch
  switch v := i.(type) {
    case int:
      fmt.Println("value:", v)
    case string:
    fmt.Println("value:", v)
    default://if all above fail, val will still be of type interface{}
  }
  ```