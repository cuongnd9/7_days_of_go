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

## documents

- [How To Install Go and Set Up a Local Programming Environment on macOS](https://www.digitalocean.com/community/tutorials/how-to-install-go-and-set-up-a-local-programming-environment-on-macos)
