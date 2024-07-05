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
- select lets you wait on multiple channels. The case however is chosen at random   
  ```go
  select {
    case <-c:
    case <-c2:
    default: // if nothing recieved from any channel execute default case
  }
  ```
- sync provides mutex package which can be used to access a shared resources between multiple go routines. waitGroups can be used to wait until all go routines are executed in main routine