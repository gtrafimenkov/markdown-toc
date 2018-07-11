Change Log
==========

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/) and this
project adheres to [Semantic Versioning](http://semver.org/).

- [Unreleased](#unreleased)
- [0.0.1 - 2018-02-07](#001---2018-02-07)

<!--
Added      new features.
Changed    changes in existing functionality.
Deprecated once-stable features removed in upcoming releases.
Removed    deprecated features removed in this release.
Fixed      any bug fixes.
Security   invite users to upgrade in case of vulnerabilities.
-->

[Unreleased]
------------

### Added

- Ability to create GitLab styled TOC.  See `-style` command line option.

### Fixed

- Generation of TOC with unicode characters

[0.0.1] - 2018-02-07
--------------------

### Fixed
- Link generation for headers starting with unicode characters.

### Changed
- Don't exit with non-zero exit code in the update mode.
- Project forked from [nochso/tocenize](https://github.com/nochso/tocenize).
