package cli

import (
    "strings"
    "fmt"
    "errors"
    "github.com/rwn3120/go-table"
    "os"
    "unicode"
    "regexp"
)

type cmdInfo interface {
    trigger() string
    description() string
    options() map[string]*Option
    shortOptions() map[string]*Option
    groups() map[string]*Grp
    commands() map[string]*Cmd
    arguments() []*Arg
    Usage()
}

func addHelp(cli cmdInfo) cmdInfo {
    return addOptions(cli, FlagOptFunc(func() error {
        cli.Usage()
        os.Exit(0)
        return nil
    }, "help", 'h', "Show help and exit"))
}

func addOptions(cli cmdInfo, options ...*Option) cmdInfo {
    for _, option := range options {
        if option != nil {
            checkDuplicates(cli, option.long)
            cli.options()[option.long] = option
            checkDuplicates(cli, option.short)
            cli.shortOptions()[option.short] = option
        }
    }
    return cli
}

func addGroups(cli cmdInfo, categories ...*Grp) cmdInfo {
    for _, group := range categories {
        if group != nil {
            checkDuplicates(cli, group.name)
            cli.groups()[group.name] = group
        }
    }
    return cli
}

func addCommands(cli cmdInfo, commands ...*Cmd) cmdInfo {
    for _, command := range commands {
        if command != nil {
            checkDuplicates(command, command.name)
            cli.commands()[command.name] = command
        }
    }
    return cli
}

func checkMissingOptions(cli cmdInfo) error {
    var missingOptions []string
    for _, option := range cli.options() {
        if option.required && !option.used {
            missingOptions = append(missingOptions, option.long+" "+option.expects())
        }
    }
    if len(missingOptions) > 0 {
        return errors.New("missing required options: " + strings.Join(missingOptions, ","))
    }
    return nil
}

func checkDuplicates(cli cmdInfo, name string) {
    if _, exists := cli.groups()[name]; exists {
        panic(fmt.Sprintf("Duplicit group %s", name))
    }
    if _, exists := cli.commands()[name]; exists {
        panic(fmt.Sprintf("Duplicit command %s", name))
    }
    if _, exists := cli.options()[name]; exists {
        panic(fmt.Sprintf("Duplicit option %s", name))
    }
}

func usage(cli cmdInfo) {
    Infof("Usage: ")

    if cli.commands() != nil {
        fmt.Printf("%s [OPTIONS] <COMMAND> [ARGS]...\n", cli.trigger())
    } else {
        if len(cli.arguments()) > 0 {
            fmt.Printf("%s [OPTIONS]", cli.trigger())
            for _, arg := range cli.arguments() {
                fmt.Printf(" %s", arg.String())
            }
            fmt.Println()
        } else {
            fmt.Printf("%s [OPTIONS]\n", cli.trigger())
        }

    }
    Important("\n" + cli.description())

    options := table.New(96, 16, indentSize, false)
    for _, option := range cli.options() {
        options.Row(option.trigger(), option.description())
    }

    if options.Size() > 0 {
        Info("\nOptions:")
        options.Print()
    }

    commands := table.New(96, 16, indentSize, false)
    for _, group := range cli.groups() {
        commands.Row(group.trigger(), group.description())
    }

    for _, command := range cli.commands() {
        commands.Row(command.trigger(), command.description())
    }

    if commands.Size() > 0 {
        Info("\nCommands:")
        commands.Print()
    }
}

func process(cli cmdInfo, args []string, requiresArg bool) error {
    for index := 0; index < len(args); index++ {
        arg := args[index]
        // options
        if option, found := cli.options()[arg]; found {
            if option.argType == flag {
                if err := option.set("true"); err != nil {
                    return err
                }
            } else {
                if index+1 >= len(args) {
                    return errors.New("Missing " + arg + " value")
                }
                index++
                if err := option.set(args[index]); err != nil {
                    return err
                }
            }

            // short options
        } else if option, found := cli.shortOptions()[arg]; found {
            if option.argType == flag {
                if err := option.set("true"); err != nil {
                    return err
                }
            } else {
                if index+1 >= len(args) {
                    return errors.New("Missing " + arg + " value")
                }
                index++
                if err := option.set(args[index]); err != nil {
                    return err
                }
            }

            // groups
        } else if group, found := cli.groups()[arg]; found {
            if err := checkMissingOptions(cli); err != nil {
                return err
            }
            if err := checkMissingOptions(cli); err != nil {
                return err
            }
            return process(group, args[index+1:], true)

            //commands
        } else if command, found := cli.commands()[arg]; found {
            if err := checkMissingOptions(cli); err != nil {
                return err
            }
            if err := process(command, args[index+1:], false); err != nil {
                return err
            }
            if err := checkMissingOptions(cli); err != nil {
                return err
            }
            if err := checkMissingOptions(command); err != nil {
                return err
            }
            return command.handler(args[index+1:])
        } else {
            if requiresArg {
                return errors.New("Unknown argument: " + arg)
            } else {
                break
            }
        }
    }

    if requiresArg {
        cli.Usage()
        return nil
    }
    return nil
}

func Sentence(format string, args ...interface{}) string {
    text := fmt.Sprintf(format, args...)
    text = strings.TrimSpace(text)
    if !strings.HasSuffix(text, ".") ||
        !strings.HasSuffix(text, "?") ||
        !strings.HasSuffix(text, "!") {
        text = text + "."
    }
    words := strings.Fields(text)
    if len(words) > 0 {
        words[0] = strings.Title(words[0])
    }
    return strings.Join(words, " ")
}

func Escape(format string, args ...interface{}) string {
    name := fmt.Sprintf(format, args...)
    re := regexp.MustCompile("^-*")
    name = re.ReplaceAllString(name, "")
    var result string
    for _, char := range name {
        if unicode.IsUpper(char) && !strings.HasSuffix(result, "-") {
            result = fmt.Sprintf("%s-%s", result, strings.ToLower(string(char)))
        } else {
            result = fmt.Sprintf("%s%s", result, strings.ToLower(string(char)))
        }
    }
    return result
}
