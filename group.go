package cli

type Grp struct {
    name      string
    desc      string
    opts      map[string]*Option
    shortOpts map[string]*Option
    cmds      map[string]*Cmd
    grps      map[string]*Grp
}

func GroupWithoutHelp(name string, description string, commands ...*Cmd) *Grp {
    group := &Grp{
        name:      Escape(name),
        desc:      Sentence(description),
        opts:      make(map[string]*Option),
        shortOpts: make(map[string]*Option),
        cmds:      make(map[string]*Cmd),
        grps:      make(map[string]*Grp)}
    group.AddCommands(commands...)
    return group
}

func Group(name string, description string, commands ...*Cmd) *Grp {
    return GroupWithoutHelp(name, description, commands...).AddHelp()
}

func (g *Grp) AddHelp() *Grp {
    return addHelp(g).(*Grp)
}

func (g *Grp) AddGroups(groups ...*Grp) *Grp {
    return addGroups(g, groups...).(*Grp)
}

func (g *Grp) AddOptions(options ...*Option) *Grp {
    return addOptions(g, options...).(*Grp)
}

func (g *Grp) AddCommands(commands ...*Cmd) *Grp {
    return addCommands(g, commands...).(*Grp)
}

func (g *Grp) Usage() {
    usage(g)
}

func (g *Grp) options() map[string]*Option {
    return g.opts
}

func (g *Grp) shortOptions() map[string]*Option {
    return g.shortOpts
}

func (g *Grp) groups() map[string]*Grp {
    return g.grps
}

func (g *Grp) commands() map[string]*Cmd {
    return g.cmds
}

func (g *Grp) trigger() string {
    return g.name
}

func (g *Grp) description() string {
    return g.desc
}

func (g *Grp) arguments() []*Arg {
    return nil
}
