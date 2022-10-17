A golang module which provide terminfo database access.

still need some time to be ready.

## original content for README.md

This package represents the parent for all terminals.

In older versions of tcell we had (a couple of) different
external file formats for the terminal database. Those are
now removed. All terminal definitions are supplied by
one of two methods:

1. Compiled Go code

2. For systems with terminfo and infocmp, dynamically
   generated at runtime.

The Go code can be generated using the mkinfo utility in
this directory. The database entry should be generated
into a package in a directory named as the first character
of the package name. (This permits us to group them all
without having a huge directory of little packages.)

It may be desirable to add new packages to the extended
package, or -- rarely -- the base package.

Applications which want to have the large set of terminal
descriptions built into the binary can simply import the
extended package. Otherwise a smaller reasonable default
set (the base package) will be included instead.

## Addition information

For alpine user, use the following command to add terminfo database:
`gen.sh` script need terminfo database.

```sh
% apk add foot-extra-terminfo rxvt-unicode-terminfo ncurses-terminfo wezterm-extra-terminfo ncurses-terminfo-base
```

## Todo list

The default/base terminfo is determined by `apk info -L ncurses-terminfo-base`. which is:

```
ncurses-terminfo-base-6.3_p20221008-r0 contains:
etc/terminfo/a/alacritty
etc/terminfo/a/ansi
etc/terminfo/d/dumb
etc/terminfo/g/gnome
etc/terminfo/g/gnome-256color
etc/terminfo/k/konsole
etc/terminfo/k/konsole-256color
etc/terminfo/k/konsole-linux
etc/terminfo/l/linux
etc/terminfo/p/putty
etc/terminfo/p/putty-256color
etc/terminfo/r/rxvt
etc/terminfo/r/rxvt-256color
etc/terminfo/s/screen
etc/terminfo/s/screen-256color
etc/terminfo/s/st-0.6
etc/terminfo/s/st-0.7
etc/terminfo/s/st-0.8
etc/terminfo/s/st-16color
etc/terminfo/s/st-256color
etc/terminfo/s/st-direct
etc/terminfo/s/sun
etc/terminfo/t/terminator
etc/terminfo/t/terminology
etc/terminfo/t/terminology-0.6.1
etc/terminfo/t/terminology-1.0.0
etc/terminfo/t/terminology-1.8.1
etc/terminfo/t/tmux
etc/terminfo/t/tmux-256color
etc/terminfo/v/vt100
etc/terminfo/v/vt102
etc/terminfo/v/vt200
etc/terminfo/v/vt220
etc/terminfo/v/vt52
etc/terminfo/v/vte
etc/terminfo/v/vte-256color
etc/terminfo/x/xterm
etc/terminfo/x/xterm-256color
etc/terminfo/x/xterm-color
etc/terminfo/x/xterm-kitty
etc/terminfo/x/xterm-xfree86
```
