package cli

import (
    "testing"
    "fmt"
    "strings"
)

var uppercase = false
var name = ""

func cmdHandler(args []string) error {
    if uppercase {
        fmt.Println("Hello", strings.ToUpper(name))
    } else {
        fmt.Println("Hello", name)
    }
    fmt.Printf("args (%d):\n", len(args))
    for i, arg := range args {
        fmt.Printf("\t[%d] %s\n", i, arg)
    }
    fmt.Println()
    return nil
}

func TestCli(t *testing.T) {
    myCli := New("my CLI", "x.y")
    myCli.AddOptions(
        FlagOpt(&uppercase, "uppercase", 'U', "sets uppercase"),
        RequiredStringOpt(&name, "name", 'n', "sets name"))
    myCli.AddCommands(
        Command(cmdHandler, "greetings", "command description"))

    err := myCli.Handle([]string{"--name", "sir", "greetings", "arg1", "arg2"})
    if err != nil {
        t.Error(err.Error())
    }

    err = myCli.Handle([]string{"--name", "madam", "-U", "greetings"})
    if err != nil {
        t.Error(err.Error())
    }
}
