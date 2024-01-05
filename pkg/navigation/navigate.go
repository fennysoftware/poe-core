package navigation

import (
	"fmt"
	"strings"

	modules "github.com/fennysoftware/poe-core/pkg/modules"
	packages "github.com/fennysoftware/poe-core/pkg/packages"
	projects "github.com/fennysoftware/poe-core/pkg/projects"
	"golang.org/x/exp/slices"
)

type Navigator struct {
	cwd     CWD
	projs   map[string]*projects.Project
	methods []Command
}

func (n *Navigator) AddCommand(c Command) {
	mds := n.methods
	mds = append(mds, c)
	n.methods = mds
}

func (n *Navigator) ReadCommand() []string {
	command := readline(n.String() + " % ")
	cmds := strings.Split(command, " ")
	if len(cmds[0]) == 0 {
		return []string{}
	}
	return cmds
}

func (n *Navigator) RunCommand(cmd string, args []string) (error, []string, bool) {
	res := []string{}
	if len(cmd) == 0 {
		return nil, res, false
	}
	for _, v := range n.methods {
		if v.command == cmd {
			if len(v.supportedlevels) > 0 && !slices.Contains(v.supportedlevels, n.cwd.CurrentLevel()) {
				return fmt.Errorf("Command %s not supported at level: %s", cmd, n.cwd.CurrentLevel()), res, false
			}

			msg, err := v.Check(args)
			if err != nil {
				return err, []string{msg}, false
			}
			return v.Method(n, args)
		}
	}

	return fmt.Errorf("Command %s not found", cmd), res, false
}

func (n *Navigator) CheckCreate(pname string) error {

	if n.cwd.IsTop() {
		proj := projects.NewProject(pname)
		_, ok := n.projs[pname]
		if !ok {
			n.projs[pname] = proj
		}
		return nil
	} else {

		buffer := n.cwd.CurrentPathBuffer()
		projname := buffer[Project]
		active, err := getActiveProject(n, projname)
		if err != nil {
			return err
		}

		switch n.cwd.CurrentLevel() {

		case Project:
			mod, err := active.GetModule(pname)
			if err == nil && mod != nil {
				return fmt.Errorf("module %s already exists", pname)
			} else {
				newmod := modules.NewModule(pname)
				active.UpdateModule(newmod)
			}
		case Module:
			modname := buffer[Module]
			mod, err := active.GetModule(modname)
			if err != nil {
				return fmt.Errorf("package %s does not exist", pname)
			} else {
				pkg := packages.NewPackage(pname)
				mod.UpdatePackage(pkg)
				active.UpdateModule(*mod)
			}

		default:
			return fmt.Errorf("unknown Level %d", n.cwd.CurrentLevel())

		}
		return nil
	}
}

func (n *Navigator) String() string {
	return n.cwd.String()
}

func (n *Navigator) Level() string {
	return n.cwd.CurrentLevel().String()
}

func NewNavigator(top string) *Navigator {
	nav := &Navigator{projs: make(map[string]*projects.Project), cwd: NewCWD(top)}
	nav.AddCommand(Command{command: "ls", maxnumberofargs: 0, needsargs: false, Method: list})
	nav.AddCommand(Command{command: "cd", maxnumberofargs: 0, needsargs: false, Method: changepos})
	nav.AddCommand(Command{command: "pwd", maxnumberofargs: 0, needsargs: false, Method: pwd})
	nav.AddCommand(Command{command: "lvl", maxnumberofargs: 0, needsargs: false, Method: lvl})
	nav.AddCommand(Command{command: "help", maxnumberofargs: 0, needsargs: false, Method: printhlp})
	nav.AddCommand(Command{command: "exit", maxnumberofargs: 0, needsargs: false, Method: exit})
	nav.AddCommand(Command{command: "path", maxnumberofargs: 0, needsargs: false, Method: pathbuf})
	return nav
}

func getActiveProject(n *Navigator, name string) (*projects.Project, error) {
	active, ok := n.projs[name]
	if !ok {
		return nil, fmt.Errorf("project %s does not exist", name)
	}
	return active, nil
}
