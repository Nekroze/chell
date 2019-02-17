# chell

An interactive shell with a POSIX like syntax, `chell` will monitor commands executed and move them into a new terminal split if they take too long to allow continued execution of commands. Additionally `chell` provides exit reports for failed commands and suggestions on alternate commands or switches in various circumstances.

## Install

At present you will require a checkout of the source code, from the root of which you may run either of the following two commands to build and run `chell` either locally or in docker respectively:

```bash
make run
```

```bash
make docker
```

### Runtime Requirements

[Tmux][1] is required to run `chell` in its ideal mode of operation where to provide the terminal multiplexing functionality behind `chell` and its portals.

## Portals?

Yes portals! and by that I mean that long running commands are popped into their own terminal split and a new `chell` gets spawned for continued shell usage. When the commands in a portal finish an exit report is displayed, keeping the split open until a user acknowledges the exit report.

## Exit Reports?

Most shells by default silently record the exit code of command that are executed however this is not typically displayed unless the user sets up a custom prompt that displays it. When using `chell` this information is surfaced at the end of a portal command or when a normal command fails via an exit report along with other information such as how long the execution took
