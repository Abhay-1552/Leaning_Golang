![Custom Go Gopher](https://github.com/Abhay-1552/Leaning_Golang/blob/main/asset/1.png)

Here are some common terminal commands frequently used in Go development:

1. **go run:** Compiles and executes one or more Go source files. For example:
   ```bash
   go run main.go
   ```

2. **go build:** Compiles Go source files and generates an executable binary. For example:
   ```bash
   go build main.go
   ```

3. **go install:** Compiles and installs packages and dependencies. For executables, it places them in the `$GOPATH/bin` directory. For example:
   ```bash
   go install
   ```

4. **go get:** Downloads and installs packages and dependencies from remote repositories. For example:
   ```bash
   go get github.com/example/package
   ```

5. **go mod init:** Initializes a new module and creates a new `go.mod` file. For example:
   ```bash
   go mod init module_name
   ```

6. **go mod tidy:** Cleans up the `go.mod` file by adding missing modules and removing unnecessary ones. For example:
   ```bash
   go mod tidy
   ```

7. **go test:** Runs tests in the current package directory. For example:
   ```bash
   go test
   ```

8. **go fmt:** Formats Go source code files according to Go's style guidelines. For example:
   ```bash
   go fmt file.go
   ```

9. **go vet:** Reports suspicious constructs in Go code. For example:
   ```bash
   go vet
   ```

10. **go clean:** Removes object files and cached files from the workspace. For example:
    ```bash
    go clean
    ```

11. **go doc:** Displays documentation for Go packages. For example:
    ```bash
    go doc fmt
    ```

These are some of the commonly used terminal commands for Go development. They are indispensable for managing dependencies, compiling, testing, and maintaining Go projects.
