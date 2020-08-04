# day 1

## install Go on Mac (with homebrew)

### install Go

```sh
brew update && brew install golang
```

### setup workspace

```sh
mkdir -p $HOME/go/{bin,src}
```

- src: The directory that contains Go source files. A source file is a file that you write using the Go programming language. Source files are used by the Go compiler to create an executable binary file.
- bin: The directory that contains executables built and installed by the Go tools. Executables are binary files that run on your system and execute tasks. These are typically the programs compiled by your source code or another downloaded Go source code.

### setup environment

```sh
nano ~/.zshrc
```

```sh
# ~/.zshrc
...
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

```sh
source ~/.zshrc
```

## 📒 notes

### 🥚 basic types

```
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```

|  type  	|                    limit                   	|
|:------:	|:------------------------------------------:	|
| uint8  	| 0 – 255                                    	|
| uint16 	| 0 – 65535                                  	|
| uint32 	| 0 – 4294967295                             	|
| uint64 	| 0 – 18446744073709551615                   	|
| int8   	| -128 – 127                                 	|
| int16  	| -32768 – 32767                             	|
| int32  	| -2147483648 – 2147483647                   	|
| int64  	| -9223372036854775808 – 9223372036854775807 	|

## documents

- [How To Install Go and Set Up a Local Programming Environment on macOS](https://www.digitalocean.com/community/tutorials/how-to-install-go-and-set-up-a-local-programming-environment-on-macos)
- [fly-with-go](https://github.com/103cuong/fly-with-go)

<img style="margin-left: 40px" alt="fly-with-go" src="https://github-readme-stats.vercel.app/api/pin/?username=103cuong&repo=fly-with-go" />

- [Golang for Node.js Developers](https://github.com/miguelmota/golang-for-nodejs-developers)

<img style="margin-left: 40px" alt="golang-for-nodejs-developers" src="https://github-readme-stats.vercel.app/api/pin/?username=miguelmota&repo=golang-for-nodejs-developers" />
