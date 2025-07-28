- Dont name packages as generic nouns
- Instead ensure directory where package resides matches with package Name as much as possible
- Its fine to have punctuation in directory Names, however use cameCase for package names
- Each module has multiple packages
- If a group of functions have relatred fucntionality, instead try to group them on what they are targeting and shorten the functionNames. For example if you have funcitonsLike MatchName & ExtractName targeting names slice, you can instead have package called name and functions Extract, Match
- Similarly When you have sentinel errors, try to Name them starting with Err
- Go modules follow semantic versioning(major.minor.patch)
- When you are increasing minor or patch no speaicla considerations are required
- However if you decide to create major release, you need to follow couple of steps:
 - First make changes you want on the current branch
 - Then checkout to branch called v2
 - In this branch's go.mod update module name to moduleName/vN
 - Create a release with vN
- To embed stuff into go binary:   
  ```go
  package main

  import (
  	 _ "embed"
 	 "fmt"
 	 "os"
 	 "strings"
  )

  //go:embed passwords.txt
  var passwords string

  func main() {
   	  pwds := strings.Split(passwords, "\n")
	  if len(os.Args) > 1 {
		  for _, v := range pwds {
			  if v == os.Args[1] {
				  fmt.Println("true")
				  os.Exit(0)
			  }
		  }
		  fmt.Println("false")
	  }
  }
  ```    
  Here passwords.txt must be in same directory of this program file. Then import embded package as shown. Then you must create avriable of type string, []byte or embed.FS(Virtual filesystem). This variable must be annotated with `go:embed` directive, so that the content of file is loaded into that variable.
- To load directories you have to annotate `//go:embed directory-name` over variable of type embed.FS. Then this variable will be holding directory specified by name. You can iterate over this variable as if you iterate over real filesystem. All io/fs methods can be used on this variable
- `//go:generate` annotations specified once in random source files in project will be used during `go generate ./...` to run arbitary tools. Syntax: `go:generate toolName --tool-options`
