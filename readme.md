# Exe - simple executables for Go

## Usage

    import "github.com/hekonsek/exe"
    
    var outputLines []string
    outputLines, err := exe.New().run("echo Hello world!")
    