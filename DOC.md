# Commands

## add
```
Add a new stew

Usage:
stew add [flags]

Required Flags:
-n, --name string   Name of the stew
-p, --path string   Path to the stew

Flags:
-d, --description string   Description of the stew (default "no description provided")
-g, --git                  If the stew uses git
-h, --help                 help for add
-n, --name string          Name of the stew
-p, --path string          Path to the stew (default "/Users/benjamin/Documents/GitHub/stew")
```

## edit
```
Edit an existing stew

Usage:
stew edit [flags]

Required Flags:
-i OR -s to get a stew to edit
at least one of --n , -p, or -d must be provided to make an edit

Flags:
  -d, --description string   The new description of the stew
  -h, --help                 help for edit
  -i, --id int               The id of the stew (default -1)
  -n, --name string          The new name of the stew
  -p, --path string          The new path of the stew
  -s, --stew string          The stew you want to edit

Global Flags:
      --config string   config file (default is $HOME/.config/stew/config.yaml)
```

## get
```
Get a stew from a given id or name

Usage:
  stew get [flags]

Required Flags:
at least one of --i or --n must be provided to get a stew from an id or name

Flags:
  -h, --help          help for get
  -i, --id int        The id of the stew (default -1)
  -n, --name string   The name of the stew

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

## remove
```
Remove a stew from a given id or name

Usage:
  stew remove [flags]

Required Flags:
at least one of --i or --n must be provided to remove a stew from an id or name

Flags:
  -h, --help          help for remove
  -i, --id int        The id of the stew (default -1)
  -n, --name string   The name of the stew

Global Flags:
  --config string   config file (default is $HOME/.config/stew/config.yaml)
```

### use
```
Use a stew from a given id or name

Usage:
  stew use [flags]

Required Flags:
at least one of --i or --n must be provided to use a stew from an id or name

Flags:
  -h, --help          help for use
  -i, --id int        The id of the stew (default -1)
  -n, --name string   The name of the stew

Global Flags:
  --config string   config file (default is $HOME/.config/stew/config.yaml)
```
