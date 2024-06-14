# Commands

## add
```
Add a new stew to saved templates

Usage:
stew add <name_of_stew> [flags]

Flags:
-d, --description string   Description of the stew (default "no description provided")
-h, --help                 help for add
-p, --path string          Path to the stew (defaults to current directory)
```

## edit
```
Edit the values of a saved stew

Usage:
stew edit <name_of_stew> [flags]

Required Flags:
at least one of --n , -p, or -d must be provided to make an edit

Flags:
  -d, --description string   The new description of the stew
  -h, --help                 help for edit
  -n, --name string          The new name of the stew
  -p, --path string          The new path of the stew

Global Flags:
      --config string   config file (default is $HOME/.config/stew/config.yaml)
```

## get
```
Get values of a saved stew

Usage:
  stew get <name_of_stew> [flags]

Flags:
  -h, --help          help for get
  -t, --tree          Print the tree of the stew

Global Flags:
      --config string   config file (default is $HOME/.config/stew/config.yaml)
```

## list
```
List all stews

Usage:
  stew list [flags]

Flags:
  -h, --help   help for list

Global Flags:
  --config string   config file (default is $HOME/.config/stew/config.yaml)
```

### new
```
Create a new instance of a stew from a given name

Usage:
stew new <name_of_stew> [flags]

Flags:
  -h, --help          help for new
  -p, --path string   The path to the stew (defaults to current directory)
  -f, --force         Force the creation of a new stew even if the stew already exists

  Global Flags:
--config string   config file (default is $HOME/.config/stew/config.yaml)
```

## remove
```
Remove a stew

Usage:
  stew remove <name_of_stew> [flags]

Flags:
  -h, --help          help for remove

Global Flags:
  --config string   config file (default is $HOME/.config/stew/config.yaml)
```

## replace
```
Replace all instances of a string in a project

Usage:
stew replace <old_string> <new_string> [flags]

Flags:
  -h, --help          help for replace
  -p, --path string   The path to the stew (defaults to current directory)
  -i, --ignore-case   Ignore case when searching for the old string

  Global Flags:
--config string   config file (default is $HOME/.config/stew/config.yaml)
```
