package navigation

import (
	"fmt"
)

type Level int

const (
	Root        Level = 0
	Project     Level = 1
	Module      Level = 2
	Package     Level = 3
	StructInter Level = 4
)

func (s Level) String() string {
	switch s {
	case Root:
		return "Root"
	case Project:
		return "Project"
	case Module:
		return "Module"
	case Package:
		return "Package"
	case StructInter:
		return "Structure or Interface"
	}
	return "unknown"
}

type CWD struct {
	//active *projects.Project
	// everything under top
	path   map[Level]string
	cLevel Level
	top    string
}

func (cwd CWD) CurrentLevel() Level {
	return cwd.cLevel
}

func (cwd CWD) CurrentPathBuffer() map[Level]string {
	buff := make(map[Level]string)
	for lvl, name := range cwd.path {
		if lvl <= cwd.cLevel {
			buff[lvl] = name
		}
	}
	return buff
}

func (cwd CWD) CurrentItem() (string, error) {
	name, ok := cwd.path[cwd.cLevel]
	if !ok {
		return "", fmt.Errorf("not item found at current level [%s] - not in path", cwd.cLevel)
	}

	return name, nil
}

func (cwd CWD) String() string {
	/*
		switch cwd.CurrentLevel {
		case Root:
			return cwd.top
		case Project:
			return cwd.top + ""
		default:
			return cwd.top
		}
	*/
	// this will be /<root if name provided>
	s := cwd.top
	for lvl, name := range cwd.path {

		// this will be
		// /<root>/<project>/<module>/<package>/<struct-int>
		// that's it all we support
		if lvl <= cwd.cLevel {
			s = fmt.Sprintf("%s/%s", s, name)
		}
	}
	return s
}

func NewCWD(top string) CWD {
	return CWD{path: make(map[Level]string), cLevel: Root, top: top}
}

func (cwd CWD) SetLevel(lvl Level) CWD {
	cwd.cLevel = lvl
	return cwd
}

func (cwd CWD) IsTop() bool {
	return cwd.cLevel == Root
}

func (cwd CWD) Up() CWD {
	if cwd.IsTop() {
		// cannot go any higher stop at root
		cwd1 := cwd.SetLevel(Root)
		return cwd1
	} else {
		cwd1 := cwd.SetLevel(cwd.cLevel - 1)
		return cwd1
	}
	return cwd
}

func (cwd CWD) Down(name string) CWD {
	// we are at the top so do not move any higher
	cwd1 := cwd.SetLevel(cwd.cLevel + 1)
	cwd1.path[cwd1.cLevel] = name
	return cwd1
}

/*
func (cwd CWD) SetActive(active *projects.Project) CWD {
	cwd.active = active
	cwd1 := cwd.SetLevel(Project)
	return cwd1
}
*/
