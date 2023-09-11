# My Go CLI

I'm currently developing this CLI to automate some usual tasks that I have (and also some fun experiments)

## Notes

To build the cli functionalities locally, you'll have to run the `go build cmd/main.go` command on your machine. Then, you'll use the `main` file generated like you would usually use a CLI 😀

Instead, if you do not want to build it locally, you can use the binaries from the [release](https://github.com/luisotaviodesimone/my-go-cli/releases)

### Examples

```
file/path/to/main --help
```

```
file/path/to/main speak "my go cli" -v en
```

### Notes
- For the `git` command to work you have to have the `sensible-info.json` file set up and the executable file in the same directory

### To-Do

- Commands
  - Install node lts with ease
