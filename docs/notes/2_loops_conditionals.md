- Only one loop in go. All the three statements in for loop are optional
  ```
  for init; condition; counter {

  }
  ```    
  ```go
  for i := 0; i < 4; i++ {

  }
  ```  
  Note: If any vars are initailised using `:=` in above for loop they are scope bounded to that loop alone
- Conditional statements in go include if, else if, else. Unlike other languages they can expect an optional initalisation & condition statements.
  Any vars initalised again are scope bounded to entire if-else(including full ladder) ladder alone
  ```
  if condition {

  }

  if initalisation; condition {

  }
  ```  
  ```go
  if i :=0; i == 4 {

  } else if i == 6 {

  } else {

  }
  ```

### Switch statements
- You can use break to break out of any case but it snot required as it happens implicitly
- You can use multiple values in case seperated by commas(case 3,4,5,6:). You can also have blank switch statement
- You can have expressions ot be evaluated in switch case

- switch statement just like above has initialisation & variable as starting point which are both optional. Each case however can be constant, function too   
  ```go
  switch optional-var-initialisation; optional-var {
    case constant-or-expression:
    case constant-or-expression:
    default:
  }
  ```