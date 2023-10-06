# Building a REST API Client in Go

A client program takes input from the user and executes some logic. For developing a Go client, we use the Go built-in library called _Flag_ for writing CLI tooles. It refers to the command line flags.
Command-line tools are important as web user interfaces to perform system tasks. In _business-to-business(B2B)_ companies, the software is packaged as a single binary instead of having multiple different packages.
The flag package has multiple functions, such as `Int` and `String`, to handle the respective type input that's supplied as a command-line flag

To do this, we can use the `flag.String` method, as shown in the following code snippet:

```
import "flag"
var name = flag.String("name", "stranger", "your wonderful name")
```

In Windows, `.exe` will be generated when we build a `.go` file. After that, from the command line, we can run the program by calling the program name

The CLI package provides an intuitive API for creating command-line applications with ease. It allows a Go program to collect arguments and flags. It's quite handy for designing complex applications.
The cli package simplifies client application development. It is much faster and intuitive than the internal flag package.
The cli package provides three major elements:
_App_
_Flag_
_Action_

_Apps_ are used for defining namespaces in applications. A _flag_ is an actual container that stores options passed. _Action_ is a function that executes on collected options.

### Cobra, an advanced library

Like cli, cobra is a package for writing client binaries but takes a different approach. In cobra, we have to create separate commands and use them in our main app.
We can install cobra using:

```
go install github.com/spf13/cobra-cli@latest
```

Go will automatically install it in your $GOPATH/bin directory which should be in your $PATH.
Example:

```
var rootCmd = &cobra.Command{
Use:        "details",
Short:      "This project takes student information",
Long:       `A long string about description`,
Args:        cobra.MinimumNArgs(1),
Run: func(cmd *cobra*.Command, args []string) {
        name := cmd.PersistentFlags().Lookup("name").Value
        age := cmd.PersistentFlags().Lookup("age").Value
        log.Printf("Hello %s (%s years), Welcome to the command line
        world", name, age)
    },
}
```
