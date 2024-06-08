package types

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"text/tabwriter"
	"time"
)

const (
	StewPath  = ".stews.json"
	green     = "\033[32m"
	red       = "\033[31m"
	property  = "\033[36m"
	in_quotes = "\033[35m"
	reset     = "\033[0m"
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
	fmt.Printf("%vDescription%v -> '%v%s%v'\n", property, reset, in_quotes, st.Description, reset)
	fmt.Printf("%vPath%v        -> '%v%s%v'\n", property, reset, in_quotes, st.Path, reset)
	fmt.Printf("%vCreated At%v  -> %s\n\n", property, reset, formatTime(st.CreatedAt))
}

func (st Stew) PrintAdded() {
	fmt.Printf("\n`%v%s%v` has been added to your stews 🎉\n\n", green, st.Name, reset)
	fmt.Printf(" %vName%v        -> %s\n", property, reset, st.Name)
	fmt.Printf(" %vDescription%v -> '%v%s%v'\n", property, reset, in_quotes, st.Description, reset)
	fmt.Printf(" %vPath%v        -> '%v%s%v'\n", property, reset, in_quotes, st.Path, reset)
	fmt.Printf(" %vCreated At%v  -> %s\n\n", property, reset, formatTime(st.CreatedAt))
}

func (st Stew) PrintRemoved() {
	fmt.Printf("\n`%v%s%v` has been removed from your stews 👋\n\n", red, st.Name, reset)
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
	for _, t := range st {
		if t.Name == name {
			return &t, nil
		}
	}
	err := fmt.Sprintf("\n%vstew with name `%s` not found%v", red, name, reset)
	return &Stew{}, errors.New(err)
}

func (s *Stew) Edit(name, description, path string) error {

	if name != "" {
		s.Name = name
	}

	if description != "" {
		s.Description = description
	}

	if path != "" {
		s.Path = path
	}

	return nil
}

func formatTime(t time.Time) string {
	// format yyyy-mm-dd hh:mm:ss
	return t.Format("2006-01-02 15:04:05")

}

func (st *Stews) List() {
	// create a table of the stews aligned in columns of ID, Name, Description, UsesGit, CreatedAt
	if len(*st) == 0 {
		fmt.Println("\nNo stews found")
		return
	}

	table := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)
	fmt.Fprintln(table, "ID\tName\tDescription\tPath\tCreatedAt")
	fmt.Fprintln(table, "--\t----\t-----------\t----\t-------\t---------")
	for i, s := range *st {
		fmt.Fprintf(table, "%d\t%s\t'%s'\t%s\t%s\n", i, s.Name, s.Description, s.Path, formatTime(s.CreatedAt))

	}

	table.Flush()
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
