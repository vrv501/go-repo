- a unbuffered channel(make(chan dataType)) can be used to communicate between different go routines
- in an unbuffred channel send blocks until someone r=is receiving and reciever blocks until someone sends   
  ```go
  c := make(chan int)
  go func () {c <- 3]()
  v := <- c
  fmt.Println(v)
  ```
- you can create buffred channel using (make(chan int, 3)) which basically has behaviour:
  - send is not blocked until chan is full ie; number of elements in channel match with its buffer length
  - recieve will not block unless channel is empty ie; number of elements in channel is 0   
  ```go
  c := make(chan int, 2)
  c <- 1
  c <- 2
  
  v := <-c
  v1 := <-c
  fmt.Println(v,v1)
  ```
- however there is a convenient way to know sender has completed tranmmission on channel. Just let the sender close it(close(c))
  ```go
  val, ok := <-c // ok is false if c is closed by sender
  for i := range c{} // recieved values from channel until its closed. then loop is auto terminated
  ```
- Closing, writing to already closed buffered or unbuffered channel causes panic. Reading from already closed buffered returns whatever values are left in the channel or zero value if nothing is present. Reading from already close unbuffered channel returns zero value of the channel
- comma ok idion works when reading from channel to determine if channel is open or closed, ok returns false when channel is closed
- Raed/write on nil channel will hang. close will cause panic
- select lets you wait on multiple channels. The case however is chosen at random   
  ```go
  select {
    case <-c:
    case <-c2:
    default: // if nothing recieved from any channel execute default case
  }
  ```
- making channels nil which is used in  acase in select will disable that since no read/writes will be done/hung
- `for v := range c` will keep streaming values from channel c. Th eloop terminates when c is closed
- sync.Once, sync.WaitGroup, sync.Mutex should be passed a spointers if required
- break in switch, select will break from them
- sync provides mutex package which can be used to access a shared resources between multiple go routines. waitGroups can be used to wait until all go routines are executed in main routine

## HTTP Handler
- First use Header() to set any custom headers. Then use WriteHeader(int) to set status code. Finally Write([]byte) to write response body. This order is to be strictly respected

## Context
- Common pattern to save context keys
- Create your package inside that have own type and iota combination
  ```go
  // use this approach if you have multiple keys to be added to context
  type userKey int

  const (
    _ userKey = iota
    key1 // actual keys
    key2 

  )
  

  // Use this approach if you have single key to be adde din context
  type userKey struct{}   

    
  // Helper functions to add to context
  func ContextWithUser(ctx context.Context, user string) context.Context {
    return context.WithValue(ctx, key, user)
  }

  func UserFromContext(ctx context.Context) (string, bool) {
      user, ok := ctx.Value(key).(string)
      return user, ok
  }

  ```
- Calling Done() on ctx that is not yet cancelled, returns nil
