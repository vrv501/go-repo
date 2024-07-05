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