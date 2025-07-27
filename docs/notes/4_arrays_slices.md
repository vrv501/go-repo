### Arrays & Slices
- Arrays are comparable using `==` and `!=`
- nil represents lack of value for a type. it doesn't have any type
- slices aren't comparable(== and != isnt allowed)
- Slices can be converted to arrays using [4]int(slice) -- this will copy first four elements from slice to new array. Since it's a copy, any change made to original slice will not affect the new array
- Arrays are passed by value
- However if you type-cats slice as (*[4]int)(slice) -- no it will be pointer to underlying array of original slice which means any changes to slice will affect this pointer variable as well
- Full Slice expression: slice[start:end:last-position-in-slice-to-share]. Last index will be last position in original slice that will be shared with this new slice. Capacity of new slice will be lastposition-start. So if new elements are appended to this new sice then it will no longer share underlying array of original slice. Instead it will be copied to new array pointed by slice. Which means original slice will be unaffaceted when capacity of new slice is increased


- all arrays unlike C are fixed composite types. So when you say `x := [4]int{}` it literally means x is one whole composite(ncludes all 4 memory allocations & locations)  
  so if you assign array to some other var it lietrally copies everything
- arrays support comparision operators (same length and equal values)
- slice is **reference** to an uderlying array. It has capacity(number of elemnts in underlying array) & length(number of elemnts in slice)   
  ```
                                                        _ _ _ _ _ _ _
  array ---------------------------------------------> |_|_|_|_|_|_|_|
                                                        ^           ^
                                                        |           |
                                                        |           |
                                                        start       end ----> slice = array[start:end]
  slice is basically defined by start & end
  If you move(also fancy called reslice) end to left(<-------) although number of elemts in slice have reduced, the underlying array will still be unchanged.
  ie; if you move end again to right, the elemts since are still present in underlying array will not be lost and back in slice

  Keep in mind you cannot try to move end beyond capacity(also true for start, technically you cannot move start - index 0 beyond left) otherwise it will panic because the memory locations outside are not part of the array.

  However the moment start moves right the previous elemts in array are lost, thereby reducing capacity
  ```   
  ```go
  x := []int{1,2,3,4,5,6,7,8}
  x = x[:5] //prints {1,2,3,4,5} but underlying array is still present
            // len is 5 but cap is 8
  x = x[:8] // {1,2,3,4,5,6,7,8}
            // len 8 cap 8
  // x = x[:9] // panic end cannot be gretar than capacity

  x = x[2:] // will print {3,4,5,6,7,8} 
            // len 6 cap 6
  ```
- zero value of slice is nil with len0 & cap 0
- append to any slice will interanlly take care to increase capacity
- a fucntion arg with `funcName(arg ...int)` - here arg will be slice of type int
- both maps & slices are refernces to their underlying implementations. SO if you pass them as args to funcitons, or assign them to another variable, know that anymodifications will afect them all!!!!

### Pointers
- A pointer is actually represented by 4 byte or 8 byte address which stores address of variable it is pointing to
- When you pass the pointer to a function, you are actually copying the address pointed by pointer into a new variable   
```go
func f(y *int) {
  k := 2
  y = &k // useless since you are updating value pointed by y
          // but at caller function, value stored at x will still be original address of z
}
z := 2
var x *int = &z
f(x) // here let's say z's address is 0xc8d
    // then x's address will be 0xADF and the value stored at this address is z's address
    // When you do f(x), inside f() y will be holding address pointed by x
    // essentially y's address will be 0xBCD and value stored at that address will be z's address
```
- If you use new(T), it will return address which contains zero value of type   
  ```go
  x := new(int) // x is of type *int
                // x != nil and *x will be 0
  ```