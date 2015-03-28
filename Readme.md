# goshorty

A small utility to customize your command prompt. Use as a replacement for `\w` in your prompt.

Turns this:

```
~/gocode/src/github.com/sklise/goshorty
```

into

```
~/g/s/g/s/goshorty
```

Inspired by the gif on [gizak/termui](https://github.com/gizak/termui). And thus there is probably some other way of accomplishing this.

## Installation

Clone this repo into your `$GOPATH` and `cd` into the folder. There are no packages to `get` so you can run `go install` and it should install `goshorty` as an executable.

To add to your command prompt:

```
export PS1='$(pwd | goshorty)'
```