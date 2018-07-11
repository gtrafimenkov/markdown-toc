markdown-toc
============

Insert and maintain a table of contents in Markdown files.

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Alternatives](#alternatives)
- [Changes](#changes)
- [License](#license)


Features
--------

markdown-toc generates a TOC (table of content) from Markdown files and inserts or
updates it in the given file.

- Cross-platform command line utility
  - Windows, Linux, Mac and *bsd (anything the Go compiler will handle)
- Anchor links are compatible to GFM (Github flavoured Markdown) and GitLab.  It is configured
  using `-style` command line option.
- Automatic "intelligent" insertion of new TOC
- Update existing TOCs without moving it
- Configurable max. and min. header depth
- Line endings are kept intact (LF or CRLF is detected and then used for new lines)


Installation
------------

```
go get -u github.com/gtrafimenkov/markdown-toc
go install github.com/gtrafimenkov/markdown-toc/cmd/markdown-toc
```

Usage
-----

The output of `markdown-toc -h` should be self explanatory:

```
markdown-toc [options] FILE...

  -d	print full diff to stdout
  -e	update only existing TOC (no insert)
  -indent string
    	string used for nesting (default "\t")
  -max int
    	maximum depth (default 99)
  -min int
    	minimum depth (default 1)
  -p	print full result to stdout
  -style string
    	style of the TOC: github or gitlab (default "github")
  -v	print version
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
