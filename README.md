# go-cli library


## usage
```

var myCliFlag = false

func cmdHandler(args []string) error{
    println("Hello", args...)
    println("myCliFlag", myCliFlag)
    return nil
}

...
myCli := New("myCli CLI", "x.y")
myCli.AddOptions(FlagOpt(&myCliFlag, "flag", 'f', "sets flag"))
myCli.AddCommands(
    Command(cmdHandler, "hello", "prints hello"))

err := myCli.Handle([]string{"hello", "sir"})

```
