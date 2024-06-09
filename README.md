<img width="642" alt="image" src="https://github.com/BenjuhminStewart/stew/assets/82689821/94145b53-e0e2-4beb-b9ad-a34fae888875">

Introducing `stew` 🍲🎉. A CLI for creating, storing, and using templates to reduce boilerplate.

## Installation
```
go install github.com/BenjuhminStewart/stew@latest
```

## Usage
- `stew add`: add a new stew template
- `stew get`: get a stew given a certain id or name
- `stew list`: list all stew templates
- `stew remove`: remove a stew template
- `stew use`: use a stew template

For more information on usage, checkout the [DOC.md](https://github.com/BenjuhminStewart/stew/blob/main/DOC.md) file.

## Configuration

The current options for configuration are:

- `stewsPath`: the path to the file where stews are stored (default: `$HOME/.stews.json`)
- `timeFormat`: the format for the time that stews are created (default: `2006-01-02 15:04:05`)

Here is an example of a basic configuration file you could make up:

```yaml
stewsPath: /home/ben/.stews.json
timeFormat: Jan 2, 2006 @ 3:04pm
```

If you have any suggestions or issues, feel free to open an issue or PR. Enjoy! 🎉
