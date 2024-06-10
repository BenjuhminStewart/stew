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

## Examples of Use Cases
Personally, there are certain templates I like to start with for markdown files, react projects, etc., so why not create a way to easily create, store and use templates from within your terminal?

A good way to do this is to have some folder of templates you like for example in `$HOME/.config/stew/templates` or `$HOME/Documents/templates` and add new stews directed toward those paths.

For example, say you make a templates directory at `$HOME/Documents/templates` and you made a basic markdown template for when you take notes or configure a README at `$Home/Documents/templates/markdown`. You can then use `stew add markdown -p $HOME/Documents/templates/markdown -d "basic README layout"` to save it in your stews and later when you want to use it in a directory simply run `stew use markdown` and your markdown file will be copied over for easy use!

You may be wondering how this is any different from running a simple `cp -r dir1 dir2` and it comes down to organization and ease of use. I simply prefer being able to run `stew list` and see all my created stews and be able to set descriptions and configure it the way I want than to just use `cp -r`. I plan to add many more features to take this above and beyond but I need help coming up with those ideas. So if you have any requests please create an Issue or contribute yourself by cloning the repo and submitting a PR. I will do my best to review promptly.
