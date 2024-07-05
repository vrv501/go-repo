- A closure is a function value that references variables from outside its body.  
  ```go
  func adder() func(int) int {
	sum := 0
	return func(x int) int { //this is a closure
		sum += x
		return sum
	}
  }
  ```
- once adder returns above closure, vars are still bound to the closure(just like defer funcs). which means any change to sum inside closure will affect original var sum
- This means sum var is like attribute that belongs to the closure as long as it's alive.    
  ```go
  func main() {
    ad := adder()
    for i:=0; i < 5; i++ {
      fmt.Println(ad(i)) // prints 0, 1, 3, 6, 10
                         // sum is incremented with each iteration, and its alive as long as closure is present
                        // once main terminates closure & sum var are both gone
    }   
  }                
  ```
- if multiple closure is accessing pkg level variable, they are all get same refernce to the pkg level variable, not copies. So careful when using closures
- one good use of deferring closures is to dyncamically update named returns of a function since any return of defreed fucntionc annot be accessed, we can combine these two concepts to get results of deferred function