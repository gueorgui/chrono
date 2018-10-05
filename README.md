![Alt chrono](./.github/full-logo-with-tagline.png)

## Overview ##

Chrono is a time tracking tool written in Go.
It is fast and simple to use.

Want to know what your time? Chrono will help you with that. Track how
long you spend on a project. Add notes so you know exactly what
you did.

Chrono can manage thousands of frames in less than a second.

To begin using Chrono, you can start tracking a project with `chrono start [project name] [tags]`

To stop tracking the project use `chrono stop`.

#### Supported Architectures ####

We provide pre-built Chrono binaries for Windows, Linux, and macOS (Darwin) for x64, i386 and ARM architectures.

Hugo can be compiled from source where ever the Go compiler tool chain can run.

**For more information on which architectures you can install Chrono on, check out the [Go documentation](https://golang.org/doc/install).**

## Choose How to Install ##

The simplest way to install Chrono is to download the latest binary from the [releases page](https://github.com/JordanKnott/chrono/releases).
The binaries have no external dependencies.

To contribube to the Chrono project or documentation, you should [fork the GitHub project](https://github.com/jordanknott/chrono#fork-destination-box) and clone it to your machine.

Alternatively you can install Chrono by building it yourself. This ensures you're running the absolute bleeding edge version.

### Builld & Install the Binaries from Source (Advanced Install) ###

#### Prequisite Tools ####

* [Git](https://git-scm.com/)
* [Go (at least Go 1.11)](https://golang.org/dl/)


#### Downloading the source ####

As of right now, Chrono uses [dep](https://github.com/golang/dep) to manage dependencies. We'll be moving to Go Modules in the near future.

The easiest to get the source is to clone Chrono in a directory outside of `GOPATH`, for example:

``` bash
mkdir $HOME/src && cd $HOME/src
git clone https://github.com/jordanknott/chrono.git
cd chrono
go install
```

**If you are a Windows user, substitute the `$HOME` environment variable above with `%USERPROFILE%`.**

## Contributing to Chrono

For a complete guide to contributing to Chrono, see the [Contribution Guide](CONTRIBUTING.md).

We welcome contributions of many kinds from updating documentation, feature requests, bug reports & issues,
feature implementation, pull requests, answering other users questions, etc.

### Asking Support Questions

We currently don't have a discussion forum. Please create an issue on the issue tracker with a label
of `question`.

### Reporting Issues

If you believe you have found an issue or bad documentation, use
the GitHub issue tracker to report the problem to the Chrono maintainers.

If you're not sure if it's an issue, create an issue with a label of `question`.

When reporting an issue, please provide the version of chrono is use (`chrono version`)
