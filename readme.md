# Exe - simple executables for Go

## Usage

    import "github.com/hekonsek/exe"
    
    var outputLines []string
    outputLines, err := exe.New().run("echo Hello world!")
    
## License

This project is distributed under Apache 2.0 license.