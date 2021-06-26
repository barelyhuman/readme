# go-cli-template

This is what I use for my own command-line tools built with go 


## Get Started

- fork this repo and use it as a template to create a new repo

**or**

- Use the tar archive of this repo to unwrap the code into a folder

```sh
$ mkdir <project-name>
$ cd <project-name>
$ curl -fsSL https://github.com/barelyhuman/go-cli-template/archive/dev.tar.gz | tar -xz --strip-components=1
$ rm -rf go.mod
$ go mod init <project-name>
```

- Edit the main.go file to use the `<project-name>/cmd` instead of `github.com/barelyhuman/go-cli-template/cmd` 
- Then run the code with `go run .` to see `Hello World` 


