package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func run(part2 bool, input string) any {
	if input == "" {
		return "skip"
	}
	var e *exec
	execs := []*exec{}
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		if strings.HasPrefix(line, "$ ") {
			command := strings.Split(line[2:], " ")
			e = &exec{command: command}
			execs = append(execs, e)
		} else {
			e.output = append(e.output, line)
		}
	}
	fs := newFileSystem()
	for _, e := range execs {
		fs.simulate(e)
	}
	const draw = false
	if draw {
		fmt.Println(fs.root.tree())
	}
	if part2 {
		return fs.MinDeleteFor(updateSize)
	}
	return fs.root.Size100KB()
}

type exec struct {
	command []string
	output  []string
}

func (e exec) prog() string {
	return e.command[0]
}

func (e exec) arg(i int) string {
	return e.command[1+i]
}

type fileSystem struct {
	root  *dir
	wd    *dir
	total int64
}

const updateSize = 30000000

func newFileSystem() *fileSystem {
	root := newDir("", nil)
	return &fileSystem{
		root:  root,
		wd:    root,
		total: 70000000,
	}
}

func (fs *fileSystem) simulate(e *exec) {
	switch e.prog() {
	case "cd":
		fs.changeDir(e.arg(0))
	case "ls":
		fs.listDir(e.output)
	default:
		panic("unknown command")
	}
}

func (fs *fileSystem) unused() int64 {
	return fs.total - fs.root.Size()
}

func (fs *fileSystem) changeDir(target string) {
	switch target {
	case "/":
		fs.wd = fs.root
	case "..":
		if fs.wd.parent == nil {
			panic("nil parent")
		}
		fs.wd = fs.wd.parent
	default:
		fs.wd = fs.wd.subDir(target)
	}
}

func (fs fileSystem) listDir(output []string) {
	for _, line := range output {
		pair := strings.SplitN(line, " ", 2)
		name := pair[1]
		if pair[0] == "dir" {
			fs.wd.subDir(name)
			continue
		}
		size, err := strconv.ParseInt(pair[0], 10, 64)
		if err != nil {
			panic(err)
		}
		fs.wd.file(name, size)
	}
}

func (fs fileSystem) MinDeleteFor(target int64) int64 {
	required := target - fs.unused()
	if required < 0 {
		panic("no min required")
	}
	min := int64(math.MaxInt64)
	fs.root.forEach(func(d *dir) {
		s := d.Size()
		if s < required {
			return
		}
		// d is a candidate
		if s < min {
			min = s
		}
	})
	if min == math.MaxInt64 {
		panic("no min found")
	}
	return min
}

type dir struct {
	name     string
	children map[string]any
	parent   *dir
}

func newDir(name string, parent *dir) *dir {
	return &dir{
		name:     name,
		children: map[string]any{},
		parent:   parent,
	}
}

func (d *dir) subDir(name string) *dir {
	x, ok := d.children[name]
	if !ok {
		dir := newDir(name, d)
		d.children[name] = dir
		return dir
	}
	dir, ok := x.(*dir)
	if !ok {
		panic("file/dir mismatch")
	}
	return dir
}

func (d dir) tree() string {
	sb := strings.Builder{}
	fmt.Fprintf(&sb, "%s (%d) {\n", d.path(), d.Size())
	for _, c := range d.children {
		var out string
		switch v := c.(type) {
		case *dir:
			out = v.tree()
		case *file:
			out = v.String()
		default:
			panic("invalid child")
		}
		ls := strings.Split(out, "\n")
		for i := range ls {
			ls[i] = "  " + ls[i]
		}
		indented := strings.Join(ls, "\n")
		sb.WriteString(indented)
		sb.WriteRune('\n')
	}
	sb.WriteString("}")
	return sb.String()
}

func (d *dir) file(name string, size int64) {
	d.children[name] = &file{
		name: name,
		size: size,
	}
}

func (d *dir) path() string {
	if d.name == "" {
		return "/"
	}
	p := []string{d.name}
	w := d.parent
	for w != nil {
		p = append([]string{w.name}, p...)
		w = w.parent
	}
	return strings.Join(p, "/")
}

func (d dir) Size() int64 {
	sum := int64(0)
	for _, c := range d.children {
		switch v := c.(type) {
		case *dir:
			sum += v.Size()
		case *file:
			sum += v.size
		default:
			panic("invalid child")
		}
	}
	return sum
}

func (d *dir) forEach(fn func(*dir)) {
	fn(d)
	for _, c := range d.children {
		sub, ok := c.(*dir)
		if !ok {
			continue
		}
		sub.forEach(fn)
	}
}

func (d dir) Size100KB() int64 {
	const limit = 100000
	sum := int64(0)
	d.forEach(func(sub *dir) {
		size := sub.Size()
		if size < limit {
			sum += size
		}
	})
	return sum
}

type file struct {
	name string
	size int64
}

func (f file) String() string {
	return fmt.Sprintf("%s (%d)", f.name, f.size)
}
