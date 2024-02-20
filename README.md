# go-blender

Your perfect static site cocktail maker. Currently powering
https://www.awalvie.me/

It takes `.md` files as input and renders them as `.html` files via go
templating.

### INSTALL

```bash
go install -v github.com/awalvie/go-blender@latest
```

### USAGE

```
Usage: go-blender [options]

option:
        init  PATH    initialize default go-blender project in PATH
        build PATH    builds project in currect directory
```

### DIRECTORY STRUCTURE

##### `index/`

This is where your markdown files will live. `go-blender` was designed
to specifically render the markdown files a certain way. For example if
you have the following directory structure:

```
index/
|-- index.md
|-- about/
    |-- index.md
    |-- blah.md
    |-- tools/
        |-- index.md
        |-- foo.md
        |-- bar.md
```

`go-blender` alongwith the correct templates can be used to render the
following navigation bar when you visit `bar.md`'s page:

```
about/   blah     foo
         tools/   bar
```

Think how your file explorer looks.

The `index.md` file is the root of the site. Each folder inside `index/`
needs to have an `_index.md` file. This is the file that will be used to
render the page for the folder itself. The files inside the folder
itself will be rendered as themselves.

Ref: https://github.com/awalvie/lyceum/tree/master/index

> **NOTE**: Currently all files are rendered as {file_name}.html. So avoid
> duplicate file names, or submit a PR if you'd like to add a feature to
> support it

#### `static/`

This is where your static files will live. These files will be copied to
the `build/` directory.

Ref: https://github.com/awalvie/lyceum/tree/master/static

#### `templates/`

This is where your Go templates will live.

> **TODO**: Add more documentation on how to use the templates and
> what variables are available.

Ref: https://github.com/awalvie/lyceum/tree/master/templates

#### `build/`

This is where the final site HTMLs will be built to.

### SETUP

Once installed, initialize `go-blender` in a new directory:

```bash
go-blender init path/to/folder
```

This will create the following folders:

```
site/
  |- build/
  |- index/
  |- static/
  |- templates/
```

Add your markdown files to `index/` and your static files to `static/`
and your templates to `templates/`.

Build the site using:

```
go-blender build path/to/folder
```
