package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var version = "dev"

const usage = `machina -- machine layout for agents

usage: machina [flags] <command> [args]
       disk|mem|load|home-map [flags]

commands:
  status             what is present, what is missing
  install [step]     model + dirs + bin + skills  (or one of those words)
  disk               filesystems, sizes, percent full
  mem                memory and swap
  load               load averages and cpu count
  home               top-level home map (also: home-map)
  where PATH         which work dir a path belongs in
  tree               print TREE.md
  version            print the version

flags:
  -home PATH   treat PATH as $HOME
  -n           dry run; print what would happen and write nothing
  -f           overwrite existing AGENTS.md templates
  -a           disk: include virtual filesystems
  -s           home-map: compute directory sizes

exit status:
  0  ok
  1  the command failed
  2  usage error
`

func main() {
	log.SetPrefix("machina: ")
	log.SetFlags(0)
	os.Exit(run())
}

func run() int {
	home, err := os.UserHomeDir()
	if err != nil {
		home = os.Getenv("HOME")
	}
	var (
		dry   bool
		force bool
		all   bool
		sizes bool
	)

	args := os.Args[1:]
	cmd := invokedName(os.Args[0])
	if cmd == "machina" {
		cmd = ""
	}

	var rest []string
	for i := 0; i < len(args); i++ {
		a := args[i]
		switch {
		case a == "-n":
			dry = true
		case a == "-f":
			force = true
		case a == "-a":
			all = true
		case a == "-s":
			sizes = true
		case a == "-home":
			if i+1 >= len(args) {
				log.Print("-home needs a path")
				return 2
			}
			i++
			home = args[i]
		case strings.HasPrefix(a, "-home="):
			home = strings.TrimPrefix(a, "-home=")
		case a == "-h" || a == "-help" || a == "--help":
			fmt.Fprint(os.Stderr, usage)
			return 0
		case a == "-version" || a == "--version":
			fmt.Printf("machina %s\n", version)
			return 0
		case strings.HasPrefix(a, "-"):
			log.Printf("unknown flag %s", a)
			fmt.Fprint(os.Stderr, usage)
			return 2
		default:
			if cmd == "" {
				cmd = a
			} else {
				rest = append(rest, a)
			}
		}
	}

	if home == "" {
		log.Print("cannot resolve home; set HOME or pass -home")
		return 1
	}
	home, err = filepath.Abs(home)
	if err != nil {
		log.Print(err)
		return 1
	}

	opt := options{home: home, dry: dry, force: force, all: all, sizes: sizes}

	switch cmd {
	case "":
		fmt.Fprint(os.Stderr, usage)
		return 2
	case "version":
		fmt.Printf("machina %s\n", version)
		return 0
	case "status":
		if err := status(os.Stdout, opt); err != nil {
			log.Print(err)
			return 1
		}
		return 0
	case "install":
		step := "all"
		if len(rest) > 0 {
			step = rest[0]
		}
		if err := install(os.Stdout, opt, step); err != nil {
			log.Print(err)
			return 1
		}
		return 0
	case "disk":
		if err := reportDisk(os.Stdout, opt.all); err != nil {
			log.Print(err)
			return 1
		}
		return 0
	case "mem":
		if err := reportMem(os.Stdout); err != nil {
			log.Print(err)
			return 1
		}
		return 0
	case "load":
		if err := reportLoad(os.Stdout); err != nil {
			log.Print(err)
			return 1
		}
		return 0
	case "home", "home-map":
		if err := reportHome(os.Stdout, opt); err != nil {
			log.Print(err)
			return 1
		}
		return 0
	case "where":
		if len(rest) < 1 {
			log.Print("where needs a path")
			return 2
		}
		reportWhere(os.Stdout, opt.home, rest[0])
		return 0
	case "tree":
		if err := printTree(os.Stdout); err != nil {
			log.Print(err)
			return 1
		}
		return 0
	default:
		log.Printf("unknown command %q", cmd)
		fmt.Fprint(os.Stderr, usage)
		return 2
	}
}

func invokedName(arg0 string) string {
	base := filepath.Base(arg0)
	return strings.TrimSuffix(base, ".exe")
}

type options struct {
	home  string
	dry   bool
	force bool
	all   bool
	sizes bool
}

func printTree(w io.Writer) error {
	b, err := readScaffold("TREE.md")
	if err != nil {
		return err
	}
	_, err = w.Write(b)
	return err
}
