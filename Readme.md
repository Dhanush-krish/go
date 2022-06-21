## go
*   created by google in 2007 and open-sourced in 2009
*   server-side, compiled language
*   modern c language :)

### advantages
*   multi-threading
*   concurrency is cheap and easy
*   faster than python :(
*   automatic garbage collection

### application
*   cli-tools
*   perfomant application on cloud platform
*   automation application

### differences with python
*   pointers is present in go
*   only post increment is allowed in go
*   while loop is not present
*   exponent operator is not present
*   conditional expression requires relational operators **like 1%2 == 1 not 1%2**
*   not allowed to define functions inside the function
*   changing type of a variable is not allowed in go once declared
*   **division operator /** gives int for( int operands) values not float values
*   1 based indexing in  range function

initialize the module file
```
go mod init basics
```
run a go file 
```
go run  <file name>
```
list all the go environment variable
```
go env
```

###build for specific os
linux
```
GOOS=linux go build <filename>
```
windows
```
GOOS=windows go build <filename>
```
mac os
```
GOOS=darwin go build <filename>
```

reference
*   https://go.dev/doc/tutorial/getting-started
*   https://pkg.go.dev/
*   https://www.programiz.com/golang/