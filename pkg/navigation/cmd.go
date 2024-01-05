package navigation

import (
	"bufio"
	"fmt"
	"os"
)

type RunMethod func(n *Navigator, args []string) error

type Command struct {
	command         string
	zeroargsmsg     string
	maxnumberofargs int
	needsargs       bool
	Method          RunMethod
	supportedlevels []Level
	helptext        string
}

func MakeAllLevelCommand(cmd string, helptext string, meth RunMethod) Command {
	return Command{command: cmd, maxnumberofargs: 0, helptext: helptext, needsargs: false, Method: meth}
}

func (c Command) Check(args []string) (string, error) {
	if c.needsargs {
		input := ""
		if len(args) > c.maxnumberofargs {
			return "", fmt.Errorf("%s has too many arguments", c.command)
		}
		if len(args) == 0 && c.needsargs {
			input = readline(c.zeroargsmsg)
		}

		if len(args) == 1 {
			input = args[0]
		}

		return input, nil
	}

	return "", nil
}

func readline(msg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	if len(msg) > 0 {
		fmt.Print(msg)
	}
	scanner.Scan()
	return scanner.Text()
}
