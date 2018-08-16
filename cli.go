package cli

import (
    "os"
    "path/filepath"
    "fmt"
)

const (
    DefaultVersion = "v0.0.1"
)

type Cli struct {
    bin       string
    name      string
    desc      string
    version   string
    opts      map[string]*Option
    shortOpts map[string]*Option
    cmds      map[string]*Cmd
    grps      map[string]*Grp
}

func Default(description string, options ...*Option) *Cli {
    cli := &Cli{
        bin:       filepath.Base(os.Args[0]),
        version:   DefaultVersion,
        desc:      Sentence(description),
        opts:      make(map[string]*Option),
        shortOpts: make(map[string]*Option),
        cmds:      make(map[string]*Cmd),
        grps:      make(map[string]*Grp)}
    cli.AddOptions(options...)
    return cli
}

func New(description string, version string, options ...*Option) *Cli {
    return Default(description, options...).AddHelp().AddVersion(version)
}

func (c *Cli) AddVersion(version string) *Cli {
    c.version = version
    return addOptions(c, FlagOptFunc(func() error {
        fmt.Println("version:", InfoStr(c.version))
        os.Exit(0)
        return nil
    }, "version", 'v', "Show version and exit")).(*Cli)
}

func (c *Cli) AddHelp() *Cli {
    return addHelp(c).(*Cli)
}

func (c *Cli) AddOptions(options ...*Option) *Cli {
    return addOptions(c, options...).(*Cli)
}

func (c *Cli) AddGroups(groups ...*Grp) *Cli {
    return addGroups(c, groups...).(*Cli)
}

func (c *Cli) AddCommands(commands ...*Cmd) *Cli {
    return addCommands(c, commands...).(*Cli)
}

func (c *Cli) Handle(args []string) error {
    return process(c, args, true)
}

func (c *Cli) Usage() {
    usage(c)
}

func (c *Cli) options() map[string]*Option {
    return c.opts
}

func (c *Cli) shortOptions() map[string]*Option {
    return c.shortOpts
}

func (c *Cli) groups() map[string]*Grp {
    return c.grps
}

func (c *Cli) commands() map[string]*Cmd {
    return c.cmds
}

func (c *Cli) trigger() string {
    return c.bin
}

func (c *Cli) description() string {
    return c.desc
}

func (c *Cli) arguments() []*Arg {
    return nil
}

func (c *Cli) Exit(code int, errors ...error) {
    for _, err := range errors {
        Error(err.Error())
    }
    os.Exit(code)
}