### Struct Types & Methods
- Types can be declared at any level, scope in Go
- User declared types based on built in primitive types can be assigned values which are directly primitive data-types without explicit type-casting. All the built in methods for those prmitive datat-types will work as well for these user-defined types based on primitve datat-types

- Attributes of a struct that are not exported will no tbe modifyable outside the package where its declared. Same oes true for its methods. 
- Encoding packages ignore all unexported attrutes of a struct
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
- You can declare mix of pointer & value receivers on methods, but dont. When you have pointer receiver, but declare a value instance for type. You can stll accs smethods having pointer receivers because go automatically converts into rigt format. The reverse also holds true. A Pointer instance will be able to acces smethods declared using value receivers
- Same might not be true when you decide what type of (value/pointer) receivers implement methods of an interface and assign it to a variable of type intreface. There you have to be explicit

- Keep in mind if you declare your onw type on an inbuilt type, it doesn't inherit the original type's methods - no inheritance in go 😄
- You must also declare methods associated to a type in same package
- If you have a struct and some methods associated with it using pointer receiver, then if you declare a variable which is pointer to the struct holding nil. Now if you call any methods on this variable it will stll work(lol?). Which is why its better to handle nil in your pointer receiver methods. However those methods which have value receivers will panic snce there's no value being pointed by that variable

- You can embed any type not just struct in nother struct. This makes all the methods available on the embedded type to be promoted to this new struct. This is called composition
- This is useful for implementing a interface by a struct by simply embedding another type which implemnets the interface. Then the parent struct automatically implements the interface !!

### Interface
- Interfaces can be embedded inside another interfaces essentially making anything that implements this new interface also implement original interface. This will help avoid suplicate method names
- If you're type is implementing a interface and you want to ensure it always implements interface and accidentally prevent issues in runtime, you can use below to handle errors at compile time   
  ```go
  var _ InterfaceName = (*typeName)(nil) // assuming *typeName implements interface
  ```
- An interface underneath is a struct holding twp ointer fields (value, concrete-type) 
  - If value is non nil, type is always non nil which means interface is non nil
  - However if value is nil but type is non nil(example: a pointer to a data-type), interface is still non-nil
  - Both type and value must be nil to say interface is non nil. This happens if you for example declare a vaiable of tpe interface and not assin it to any type thats implementing this interface
  ```go
  var i interface = TypeName{} // Here i will technically be ({}, TypeName)
  ```
- Unlike nil instances for types having methods where you can still invoke methods, invoking methods on a nil interface will trigger panic, because type is nil in case of nil interfaces
- Interfaces are comparable as long as two variables are of same type and same value. However if types are not comparable, then interfaces also arent comparable(== and !=)
- Interface comparision follows three step-process. First types are compared. If they match then we check if both individual types are comparable. If yes, then their values are compared. However remember if pointer types is implementing an interface trying to compare two instances of this pointer type will return false. Because the interface looks like(*pointerType, address). Pointers hold addres sof variable they are pointing to. Addresses can never be equal
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
- Type assertions hsould only be used with interface types and only reveal the immediate concrete type implementing the interface. If the immediate concrete type is an alias/based on another concrete type, trying to type assert the intreface with this underlying concrete type will cause panics/fail the assertion

- You can embed a interface in a sruct so that when this attribute of type interface is prperly initialised, the enclosing struct can directly use methods on the interface. Similarly with interfaces on exception is you can embed interface inside a interface. This ensures any type implementing the enclosing interface can also implement inside interface
- When you use a type assertion in Go with the syntax interfaceVariable.(Type), the Type can be either a concrete type (such as a struct, int, string, etc.) or an interface type (another interface the value may implement)
