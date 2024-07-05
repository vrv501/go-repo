- defer evaluates once fucntion around it is done. defer calls are stacked(LIFO)
- A deferred function's arguments are evaluated when the defer statement is evaluated.(not execution, evaluation)   
  evaluation of args - in the current routine they are decided and bound to defer func, but defer func will not be executed. This is also true for go-routines. Any arguments to go-routines are evaluated immediately, but go-routine may not be immediately executed
  ```go
  func a() {
    i := 0
    defer fmt.Println(i) // after a is done will still print 0
    i++
    return
  }
  ```
- if any panics are encountered, before system exits all deferred functions are executed
- panics can be used to stop normal system execution however if there are any deferred functions so far, they will be executed before panic is done
- recover can be used to handle panics and resume normal system operation, however recover is a non-blocking fucntion which means if we call recover and there's no panic it moves on. But panic can happen beyond the call of recover, which is why its **required** to be called in a defer func  
  ```go
  defer func() {
    if r:= recover(); r != nil {
        fmt.Println("recovered: %v", r) // will print recovered: panicked!
    }}()
  panic("panicked!")
  ```
