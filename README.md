# dotgo

A rewrite of `dotbot`, but in Go so I can store the binary in my dotfiles and not have to worry about the local Python installation.

# what's complete?

Currently implemented the logging package `messenger`. Still need to do plugins as the main part, then the config follows from that.

# todo

* messenger done
* config done (maybe?)
* plugins need heavy working
* cli needs some refactor, but that requires metavar change to the package im using
* Actually add the custom parsing and all that, potentially allow config to have validation