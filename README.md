# Go cli application
- Learning techniques, principles, and libraries needed to create user-friendly command-line interfaces and command suites 
- Building your own cli app

*Covers*
design and build `commands`
working with and *parsing* `flags`, *config files* and *remote config systems*
and how to work with `environment variables` `12-factor apps`

## Index
- [ ] CLI application design *excalidraw*
- [ ] Go
- [ ] Cobra CLI framework
- [ ] Build Todo app

## User Experience of CLI's
**Human Interface Guidelines for CLI's**

> [Unix Philosophy]
> - Simple
> - Clear
> - Composable
> - Extensible
> - Modular
> - Small
>
> - Ken

### Commands
Commands do something. These are the commands that user types in command line

```sh
➜  go-cli ls
#  pwd    cmd
```

- Abbrevieted
- Compact/Short
- Context aware
- Has human friendly language comprised of `Verbs`

### Args
**These are inputs to commands**

```sh
➜  go-cli cp   [file-name] [new-file-name]
#  cwd    cmd  arg1        arg2
```

- Order of args matters
- Separated by spaces
- Usually is a `Noun`

### Options / Flags
**Modify operation performed by command**

- Space separated
- Prefixed with `-` for short, `--` for long, `/`
- Common options are often shortened
- Short options are stacked
- Usually is an `Adverb`

### Cli apps

- Pronounced as `NOUNS`

### Subcommands
**CLI apps do multiple things.** These are represented as Subcommands

- Apps are just group of commands
- These subcommands also have flags and args
- All above rules apply

## Features

- [ ] Add Todo
- [ ] List Todos
- [ ] Mark as *done*
- [ ] Search/Filter
- [ ] Priorities
- [ ] Archive
- [ ] Edit
- [ ] Create Dates
- [ ] Due Dates
- [ ] Tags
- [ ] Projects

## Interface Design

### Add todo
```sh
td add "My task" -p"H" -d"49/4/20"

td add "My task" -P"Mits8" -p"H" -d"49/4/20" \
add "Another task" -p"L" \
add "Another one"
```

### Assign priority
will use `-p` flag and **HIGH ,MEDIUM ,LOW** strategy
*Default priority* - MEDIUM

### Listing

```sh
td ls -P"Mits8" # specify project
td ls # all projects
```

### Filtering

#### By status
```sh
td ls -s"DONE"
td ls -s"INPROGRESS"
```

#### By date
```sh
td ls --created-before="" # Before today
td ls --created-after="1/2/20" # after this yr's FEB
td ls --due # due today
```

### Updating
```sh
td mark done 2
td mark inprogress 2
```

### Editing
```sh
td edit 1 "New task title"
td edit 1 -pLOW
td edit 1 --due"+3d4h"
td edit 1 --created"1/2/19"
td edit 1 -P"MitsX"

# Batch editing
td edit 1 2 3 4 -pLOW
```
## Start Coding

```sh
mkdir -p ./project-repo && cd ./project-repo
go mod init github.com/hrutvikyadav/go-cli
# go get -u github.com/spf13/cobra@latest
go install github.com/spf13/cobra-cli@latest
# inside a go module, run
cobra-cli init --author "Hrutvik Yadav hrutvikyadav@gmail.com"

go run
# OR
go build
./go-cli
```

```sh
# Add new command
cobra-cli add add
tree .
.
├── LICENSE
├── README.md
├── cmd
│   ├── add.go
│   └── root.go
├── go-cli
├── go.mod
├── go.sum
└── main.go

1 directory, 8 files

go build 
./go-cli add
add called

# Another command
cobra-cli add list
```
