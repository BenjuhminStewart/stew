package types

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/BenjuhminStewart/stew/util"
	"os"
	"time"
)

const (
	green       = "\033[32m"
	red         = "\033[31m"
	property    = "\033[36m"
	description = "\033[35m"
	path        = "\033[33m"
	reset       = "\033[0m"
)

type Stew struct {
	Name        string
	Description string
	Path        string
	CreatedAt   time.Time
}

type Stews []Stew

func (st Stews) Len() int {
	return len(st)
}

func (st Stew) Print() {
	fmt.Printf("\n%vName%v        -> %s\n", property, reset, st.Name)
	fmt.Printf("%vDescription%v -> '%v%s%v'\n", property, reset, description, st.Description, reset)
	fmt.Printf("%vPath%v        -> '%v%s%v'\n", property, reset, path, st.Path, reset)
	fmt.Printf("%vCreated At%v  -> %s\n\n", property, reset, util.FormatTime(st.CreatedAt))
}

func (st Stew) PrintAdded() {
	fmt.Printf("\n`%v%s%v` has been added to your stews 🎉\n\n", green, st.Name, reset)
	fmt.Printf(" %vName%v        -> %s\n", property, reset, st.Name)
	fmt.Printf(" %vDescription%v -> '%v%s%v'\n", property, reset, description, st.Description, reset)
	fmt.Printf(" %vPath%v        -> '%v%s%v'\n", property, reset, path, st.Path, reset)
	fmt.Printf(" %vCreated At%v  -> %s\n\n", property, reset, util.FormatTime(st.CreatedAt))
}

func (st Stew) PrintRemoved() {
	fmt.Printf("\n`%v%s%v` has been removed from your stews 👋\n", red, st.Name, reset)
}

func (st *Stews) Add(name string, description string, path string, usesGit bool) {
	if st.doesStewExist(path) {
		fmt.Printf("\n`%v%s%v` already exists in your stews\n", red, name, reset)
		return
	}

	stew := Stew{
		Name:        name,
		Description: description,
		Path:        path,
		CreatedAt:   time.Now(),
	}
	*st = append(*st, stew)

	stew.PrintAdded()
}

func (st *Stews) doesStewExist(path string) bool {
	for _, s := range *st {
		if s.Path == path {
			return true
		}
	}
	return false
}

func (st *Stews) Remove(i int) error {
	ls := *st
	if i < 0 || i >= len(ls) {
		return errors.New("invalid index")
	}

	*st = append(ls[:i], ls[i+1:]...)
	ls[i].PrintRemoved()
	return nil
}

func (st *Stews) RemoveByName(name string) error {
	for i, s := range *st {
		if s.Name == name {
			return st.Remove(i)
		}
	}
	err := fmt.Sprintf("\n%vstew with name `%s` not found%v", red, name, reset)
	return errors.New(err)
}

func (st Stews) Get(i int) (*Stew, error) {
	if i < 0 || i >= len(st) {
		err := fmt.Sprintf("\n%vstew of index %d not found%v", red, i, reset)
		return &Stew{}, errors.New(err)
	}

	return &st[i], nil
}

func (st Stews) GetByName(name string) (*Stew, error) {
	for i, t := range st {
		if t.Name == name {
			return st.Get(i)
		}
	}
	err := fmt.Sprintf("\n%vstew with name `%s` not found%v", red, name, reset)
	return &Stew{}, errors.New(err)
}

func (s *Stew) Edit(name, description, path string) error {

	if name != "" {
		fmt.Printf("\nName Changed: %s -> %v%s%v\n", s.Name, green, name, reset)
		s.Name = name
	}

	if description != "" {
		fmt.Printf("\nDescription Changed: %s -> %v%s%v\n", s.Description, green, description, reset)
		s.Description = description
	}

	if path != "" {
		fmt.Printf("\nPath Changed: %s -> %v%s%v\n", s.Path, green, path, reset)
		s.Path = path
	}

	return nil
}

func (st *Stews) List() {
	// create a table of the stews aligned in columns of ID, Name, Description, CreatedAt
	if len(*st) == 0 {
		fmt.Println("\nNo stews found")
		return
	}

	fmt.Printf("\n%v%v%v\n", property, "Stews:", reset)
	for i, s := range *st {
		fmt.Println()
		fmt.Printf(" %v%v%v [%v%v%v]", green, s.Name, reset, property, i, reset)
		fmt.Printf(": '%v%s%v'\n", description, s.Description, reset)
		fmt.Printf(" %v%s%v\n", path, s.Path, reset)
		fmt.Printf(" %s\n", util.FormatTime(s.CreatedAt))
	}

}

func (st *Stews) Load(path string) error {
	file, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return err
	}

	err = json.Unmarshal(file, st)
	if err != nil {
		return err
	}

	return nil

}

func (st *Stews) Save(path string) error {
	file, err := json.Marshal(st)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, file, 0644)
	if err != nil {
		return err
	}

	return nil
}
