- first generics are not some data-sturctures. genrics are used for parameterizing types of arguments that can be passed to a fucntion or while creating your own type
  Ex: 
  ```go
  func isEqual(s []any, v any) // in this func you can pass any args and write your own logic on those arguments
                             // but you need to do type assertions when you need their conceret type and value
                            // however there's no guarantee the both arguments belong to same concrete type
  
  func isEqual[T comparable](s []T, v T) // Here you can be rest assured that both args belong to same dataType - int, or string etc not mix match
                                         // This ensures they all are restricted to same dataType
                                         // but you still need to do type assertions to know their concrete value
  
  ```
- Keep in mind, interface type also satisfies comaparable. A variable of type interface can be used with isEqual funciton. However if underlying concrete type that implements the interface isn't comparable then isEqual will panic at runtime where you are using == or !=
- generic functions. comparable basically restrict types which implement == & != and they only allow to use == and != in genric fucntions. You cannot use other operators with comparable. Ex: <, > inside generic function between `comparable` args
  ```go
  func myfuncName[T comparable](arg1 []T, arg2 T) { // you might wonder why cant we use any here
                                                    // you can however that also means you have to do type assertions and ony then you can use == & !=
                                                    // whereas with comparable you can directly start using == & != only
                                                    // its still ok to use any depending on usecase

  }

  // when calling the above funciton you do
  myFuncName([]int{}, int) // ensures in compile time itself they both belong to same dataType
  myFuncName[int](2,[]int{}) // if no args you need to explicitly infer T by specifying it when calling the function. optional when args are present
  ```
- In a generic funciton, if you're returning a value which is also of type T, then since you dont know the zero value of that type, you can instead do:   
  ```go
  var zeroValue T
  return zeroValue // will be decided in runtime :)
  ```
- You can take a argument of one type and return argument of different type which are not known at compile time using
  ```go
  func Foo[T1, T2 any](v []T1) T2 {}

  // Here you can have Foo[int](v []int) string
  // the above declaration is supported due to T1, T2
  // If you use only T1 then both arg & eturn type must be same
  ```
- generic types.
  ```go
  type myNewType[T any] struct { // you can use comparable here as well
    attr1 T
    attr2 []T
    attr3 *myNewType[T] //notice this when referncing generic Type use [T]
  }

  // when instanting the above genric type
  ////////////// notice the below concreteType in [type]
  g := myNewType[int]{attr1: int, attr2: []int{}} // compared to this if you use any without type parameterization
                                                  // it will not guarantee all of them belong to same dataType

  // that being said if you use [any] while instantiating genericType,(myNewType[any]{attr1: string, attr2: []int{}}) 
  // you will lose the benefits of typeSafety given by generics
  // you can instead combine genric fucntions & types and restrict all args to belong to same dataType
  // notice im not exporting struct
  type myNewType[T any] struct { // you can use comparable here as well
    attr1 T
    attr2 []T
    attr3 *myNewType[T] //notice this when referncing generic Type use [T]
  }

  func instantiate[T any](arg T, arg2 T) mynewType[T]  {
    return myNewType[T]{attr1: arg, attr2: []T{arg2}}
  }

  instantiate(int, int) // will compile
  instantiate(int, string) // compile-error

  ```
- generic methods **ARE NOT** possible. 
- generics functions basically ensure all types of arguments are same & not mix match

### Generics Syntax
- Generic Types:   
  ```go
  type Circle[T any] struct {
      r T
      next *Circle[T]
  }

  type Differ[T] interface {
    fmt.Stringer
    Diff(T) float64
  }

  func Foo[T Differ[T]](pair1, pair2 T) // here this is a self refernecing type constarint 
  // where T is not just any type, but any type that also implements the interface Differ
  // pair1, pair2 args are of type Differ[T] and not T
  ```
- Generic Functions:    
  ```go
  func Foo[T1, T2 any](val T1) T2 {}
  ```
- Type Terms: Can only be used as type constarints not as variable type, parameter type, return valueType    
  ```go
  type Integer interface {
    ~int | ~int16 | ~int32 // any data type that has these as underlying concrete types
                            // If you dont use ~, then only types that can satisfy 
                            // these are actual prmitive data-types. Tring to use type which implements these types
                            // will cause compile time error
    String() string // the types having above as underlying types must also implement this method
  }

  type Comparable interface {
    int | int8 | int16 | int32 | int64 | uint
  } // this interface can be used to retsrict what types can be accepted by a type constraint
    // to also allow other tpes that have underlying types a sthose metioned, you put ~ before them 

  type Ordered[T integer] interface {
    Compare(v1, v2 T) T
  }
  ```
- Type inference at compile time is impossible for functions that return another type which is determined at runtime. So you will have to explicitly specify when calling the function
- Take Comparable interface for example. There are certain limits when using constants with above interface. Let's say you want to add a constant 1000. This is a compile time error since int8 max value is 255. However you can add 100 to type comparable since 100 satisfies all the types
