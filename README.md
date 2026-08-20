# animated-computing-machine

A tiny Go command-line project that prints a short text animation.

## What this project is

This repository is a minimal Go application used to demonstrate a simple
executable plus unit test setup.

## What it is doing

When you run the program, it prints a sequence of ASCII-style frames:

- `[=     ]`
- `[==    ]`
- `[===   ]`
- `[ ==== ]`
- `[  === ]`
- `[   == ]`
- `[    = ]`

These frames represent a basic "progress bar" animation shown one line at a
time.

## Build

```bash
go build ./...
```

## Run

```bash
go run .
```