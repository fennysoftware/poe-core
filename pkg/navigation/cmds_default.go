package navigation

import (
	"fmt"
	"strings"

	types "github.com/fennysoftware/poe-core/pkg/types"
)

const (
	UPDIR = ".."
)

func exit(n *Navigator, args []string) (error, []string, bool) {
	return nil, []string{}, true
}

func printhlp(n *Navigator, args []string) (error, []string, bool) {
	res := []string{}
	res = append(res, " ")
	res = append(res, "Available Commands:")
	res = append(res, "Commands at any level")
	res = append(res, "1. exit - Exit the tool")
	res = append(res, "2. cd - change depth ")
	res = append(res, "3. ls - list")
	res = append(res, "4. help - command help")
	res = append(res, "***")
	res = append(res, " ")
	res = append(res, "Commands at Root (/poe)")
	res = append(res, "5. mkproj - Creates a Project")
	res = append(res, "***")
	res = append(res, " ")
	res = append(res, "Commands at Project (/poe/<project>)")
	res = append(res, "6. mkmod - Create a modules")
	res = append(res, "***")
	res = append(res, " ")
	res = append(res, "Commands at at module level (/poe/<project>/<modulename>)")
	res = append(res, "7. mkpkg - Creates a package under that modules")
	res = append(res, "8. mkmain - Create a main package")
	res = append(res, "***")
	res = append(res, " ")
	res = append(res, "Commands at at package level: (/poe/<project>/<modulename>/<packagename>)")
	res = append(res, "9. mkfunc - Create a new function")
	res = append(res, "10. mkvar - Create a new variable")
	res = append(res, "11. mkinter - Create a new interface")
	res = append(res, "12. mktype - Create a new type")
	res = append(res, "13. mkstr - Create a new structure")
	res = append(res, "***")
	res = append(res, " ")
	res = append(res, " ")

	return nil, res, false
}

func pwd(n *Navigator, args []string) (error, []string, bool) {
	res := []string{}
	res = append(res, fmt.Sprintf("%s", n))
	return nil, res, false
}

func lvl(n *Navigator, args []string) (error, []string, bool) {
	res := []string{}
	res = append(res, fmt.Sprintf("%s", n.Level()))
	return nil, res, false
}

func changepos(n *Navigator, args []string) (error, []string, bool) {
	res := []string{}
	// moving up!
	if len(args) == 1 && args[0] == UPDIR {
		n.cwd = n.cwd.Up()
		return nil, res, false
	}

	path := args[0]

	//we are rooted!
	relative := path[0] != '/'
	subs := strings.Split(path, "/")
	if len(subs) == 1 {
		res = append(res, fmt.Sprintf("%s %d %t", subs[0], len(subs), relative))
		if relative {
			n.cwd = n.cwd.Down(subs[0])
		} else {
			n.cwd = n.cwd.Down(subs[0])
		}
	} else {
		for _, v := range subs {
			if len(v) > 0 {
				res = append(res, fmt.Sprintf("%s %d %t", v, len(subs), relative))
				if relative {
					n.cwd = n.cwd.Down(v)
				} else {
					n.cwd = n.cwd.Down(v)
				}
			}
		}
	}

	/*
	   // moved up so set active

	   	if n.cwd.CurrentLevel == Project {
	   		setactive(n, parent.Name)
	   	}



	   // fine the name
	   _, ok := n.projs[name]

	   	if !ok {
	   		fmt.Printf("Project %s does not exist", name)
	   	} else {

	   		child := path.NewInfoCustomSeperator('/', TOP+"/"+name)
	   		// stack oveflow?
	   		//child = path.NewInfoCustom(TOP, '/', TOP+"/"+projname)
	   		n.cwd = child
	   		n.CWD.Level = Project
	   		setactive(n, name)
	   	}

	   switch cwd.CurrentLevel {
	   case Project:

	   	// looking for modules
	   	projname := n.cwd.Name
	   	_, ok := n.projs[projname]
	   	if !ok {
	   		fmt.Printf("Project %s does not exist.\n", projname)
	   	} else {
	   		child := path.NewInfoCustomSeperator('/', TOP+"/"+projname+"/"+name)
	   		n.cwd = child
	   		n.level = 2
	   	}

	   case Module:

	   		// looking for packages
	   		modname := n.cwd.Name
	   		child := n.cwd
	   		parent := child.Parent
	   		_, err := n.active.GetModule(modname)
	   		if err != nil {
	   			fmt.Printf("Module %s does not exist.\n", modname)
	   		} else {
	   			child := path.NewInfoCustomSeperator('/', parent.ParsedPath+"/"+name)
	   			n.cwd = child
	   			n.level = 2
	   		}
	   	}
	*/

	return nil, res, false
}

func list(n *Navigator, args []string) (error, []string, bool) {
	res := []string{}
	if n.cwd.IsTop() {
		for _, v := range n.projs {
			res = append(res, v.Name)
		}
	} else {

		buffer := n.cwd.CurrentPathBuffer()
		projname := buffer[Project]
		active, err := getActiveProject(n, projname)
		if err != nil {
			return err, res, false
		}

		switch n.cwd.CurrentLevel() {
		case Project:
			for _, name := range active.ListModules() {
				res = append(res, name)
			}
		case Module:
			modname := buffer[Module]
			mod, err := active.GetModule(modname)
			if err != nil {
				return err, res, false
			}
			for _, name := range mod.ListPackages() {
				res = append(res, name)
			}
		case Package:
			modname := buffer[Module]
			mod, err := active.GetModule(modname)
			if err != nil {
				return err, res, false
			}

			pkgname := buffer[Package]
			pkg, err := mod.GetPackage(pkgname)
			if err != nil {
				return err, res, false
			}

			for _, mems := range pkg.Members {
				if mems.GetType() == types.Mt_Structures {
					for _, mem := range mems.GetAllPointers() {
						res = append(res, mem.GetName())
					}

					for _, mem := range mems.GetAllReferences() {
						res = append(res, mem.GetName())
					}
				}

				if mems.GetType() == types.Mt_Interfaces {
					for _, mem := range mems.GetAllPointers() {
						res = append(res, mem.GetName())
					}

					for _, mem := range mems.GetAllReferences() {
						res = append(res, mem.GetName())
					}
				}
			}
		case StructInter:
			modname := buffer[Module]
			mod, err := active.GetModule(modname)
			if err != nil {
				return err, res, false
			}

			pkgname := buffer[Package]
			pkg, err := mod.GetPackage(pkgname)
			if err != nil {
				return err, res, false
			}

			//siname := buffer[StructInter]
			for _, mems := range pkg.Members {
				/*
					mem, err := mems.GetMember(siname)
					if err != nil {
						return
					}
					if mem != nil {
						for _, name := range mem.Members {
							fmt.Println(name)
						}
					}
				*/

				if mems.GetType() == types.Mt_Structures {
					for _, mem := range mems.GetAllPointers() {
						res = append(res, mem.GetName())
					}

					for _, mem := range mems.GetAllReferences() {
						res = append(res, mem.GetName())
					}
				}
				if mems.GetType() == types.Mt_Interfaces {
					for _, mem := range mems.GetAllPointers() {
						res = append(res, mem.GetName())
					}

					for _, mem := range mems.GetAllReferences() {
						res = append(res, mem.GetName())
					}
				}
			}
		}
	}

	return nil, res, false
}

func pathbuf(n *Navigator, args []string) (error, []string, bool) {
	res := []string{}
	for lvl, name := range n.cwd.CurrentPathBuffer() {
		res = append(res, fmt.Sprintf("%s : %s \n", lvl, name))
	}
	return nil, res, false
}
