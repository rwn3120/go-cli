package cli

type Arg struct {
    desc     string
    mandatory bool
}

func (a *Arg) String() string {
    if a.mandatory {
        return "<" + a.desc + ">"
    }
    return "[" + a.desc + "]"
}

func Argument(desc string) *Arg {
    return &Arg{desc: desc}
}

func Mandatory(arg *Arg) *Arg {
    arg.mandatory = true
    return arg
}

type Cmd struct {
    name      string
    desc      string
    opts      map[string]*Option
    shortOpts map[string]*Option
    args      []*Arg
    handler   func([]string) error
}

func CommandWithoutHelp(handler func([]string) error, name string, description string, options ...*Option) *Cmd {
    command := &Cmd{
        name:      Escape(name),
        desc:      Sentence(description),
        opts:      make(map[string]*Option),
        shortOpts: make(map[string]*Option),
        handler:   handler}
    command.AddOptions(options...)
    return command
}

func Command(handler func([]string) error, name string, description string, options ...*Option) *Cmd {
    return CommandWithoutHelp(handler, name, description, options...).AddHelp()
}

func (c *Cmd) AddHelp() *Cmd {
    return addHelp(c).(*Cmd)
}

func (c *Cmd) AddOptions(options ...*Option) *Cmd {
    return addOptions(c, options...).(*Cmd)
}

func (c *Cmd) AddArguments(arguments ...*Arg) *Cmd {
    for _, argument := range arguments {
        c.args = append(c.args, argument)
    }
    return c
}

func (c *Cmd) Usage() {
    usage(c)
}

func (c *Cmd) arguments() []*Arg {
    return c.args
}

func (c *Cmd) options() map[string]*Option {
    return c.opts
}

func (c *Cmd) shortOptions() map[string]*Option {
    return c.shortOpts
}

func (c *Cmd) groups() map[string]*Grp {
    // command has no sub-groups
    return nil
}

func (c *Cmd) commands() map[string]*Cmd {
    // command has no sub-commands
    return nil
}

func (c *Cmd) trigger() string {
    return c.name
}

func (c *Cmd) description() string {
    return c.desc
}
