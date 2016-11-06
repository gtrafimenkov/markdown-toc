tocenize
========

Insert and maintain a table of contents in Markdown files.

- [Features](#features)
- [Installation](#installation)
	- [Pre-compiled binaries](#pre-compiled-binaries)
	- [From source](#from-source)
- [Usage](#usage)
- [Alternatives](#alternatives)
- [Changes](#changes)
- [License](#license)


Features
--------

tocenize generates a TOC (table of content) from Markdown files and inserts or
updates it in the given file.

- Anchor links are compatible to GFM (Github flavoured Markdown)
- Automatic "intelligent" insertion of new TOC
- Update existing TOC
- Configurable max. and min. header depth
- Line endings are kept intact (LF or CRLF is detected and then used for new lines)


Installation
------------


### Pre-compiled binaries

Compiled binaries are available on the [releases page][releases].


### From source

If you have a working Go environment, simply run:

```
go install github.com/nochso/tocenize/cmd/tocenize
```


Usage
-----

The output of `tocenize -h` should be self explanatory:

```
tocenize [options] FILE...

  -V    print version
  -d    print full diff to stdout
  -max int
        maximum depth (default 99)
  -min int
        minimum depth (default 1)
  -p    print full result to stdout
  -u    update existing file (default true)
  -v    verbose output
```


Alternatives
------------

- [github.com/axelbellec/gotoc](https://github.com/axelbellec/gotoc) inserts on
  top, doesn't update existing TOC, doesn't support setext-style headers


Changes
-------

All notable changes to this project will be documented in the [changelog].

The format is based on [Keep a Changelog](http://keepachangelog.com/) and this
project adheres to [Semantic Versioning](http://semver.org/).


License
-------

This project is released under the [MIT license](LICENSE).


[changelog]: CHANGELOG.md
[releases]: https://github.com/nochso/tocenize/releases