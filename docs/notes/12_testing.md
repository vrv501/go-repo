- Inside a test function as you are executing test cases one by one, use t.FatalF if one of test case fails and you dont want to proceed further. Use t.Errorf if you want to continue since test cases are independent of each other

- You can create a function `func TestMain(m *testing.M)` to cleanup and teardown any pre test setup procedures and post test setup procedures   
  ``` go
  func TestMain(m *testing.M) {
      
      // pre setup before we starting executing tests one by one

      exitVal := m.Run()

      // post all tests cleanup here    

      os.Exit(exitVal)
  }
  ```
- TestMain is invoked once, not before and after each individual test. Also be aware that you can have only one TestMain per package.
- It's convention to have test files in same directory of source files. The package name can be:
  - same as source file package name, in which case you will be able to access all vars, functions
  - if its packagename suffixed with "_test", then you will only be able to access exported vars, functions
- Use t.Run to run table tests
- Mark any test function with t.Parallel() so that it can be executed without waiting for other test sfuncitons in the package. If a test is marked as parallel, using t.Setenv can cause panic or inconsistencies
- When writing fuzz tests name of fuzz function should start with Fuzz followed by argument of FuzzTest(f * testing.F). An example: https://github.com/learning-go-book-2e/file_parser/blob/a1274d69ad4ebd7ffbe3de32d5f8e85129e0f50f/file_parser_test.go#L71
- In the above hyperlink you can just concentrate on fuzz add and how data is passed to your function. Rest of stuff after can be ignored
- Fields in corpus(slice of entries added to fuzzer using f.Add) used by fuzzer are currently restricted to any int, any float, string, []byte, bool. You can mix any of these data types while adding to corpus
- Arguments to f.Fuzz must be a function (func(_, _, _) )whose arguments are:
  - t *testing.T
  - same type, order, count of arguments to actual funciton you are fuzzing 
- Once fuzz cases are written you can use `go test -fuzz={{ FuzzFuncName }}`. Only one funciton can be fuzzed at any time. In the same directory go compiler will create subdirectory: `testdata/fuzz/{{ FuzzFuncName }}/{{ randomFileName }}` for every testcase where it has seen failures. Inside this you will have lot of files showing what go compiler did and what was the inputs sent to your function
- The essense is that if your code panics or takes too long, it will counted as failure by the fuzzer
- Benchmarking example: https://github.com/learning-go-book-2e/ch15/blob/23ccc11d9e388e1ed3aea232b4568ddd7d65f1ea/sample_code/bench/bench_test.go#L55
- Benchmarking can be done using: `go test -bench=. -benchmem` -bench runs all benchmarks, -becnhmem includes memory allocation info in the output
- Profiling notes: https://jvns.ca/blog/2017/09/24/profiling-go-with-pprof/
- You can use `go test -race` to test data race conditions when executing UT's for a function

