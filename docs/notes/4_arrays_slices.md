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
