## Getting Started

Once you cloned the repo you can just execute the Makefile and will bring up all the go dependencies needed to execute this command line

 `myself@pc:/path/to/$make`

## You need

To be able to run this you should have [go](https://golang.org/doc/install) installed in your system and clone this in the `GOPATH`.

## Disclamer

I would appretiate a follow up on the discussion in the hash.go comments. I leave there an explenation about what I think that happened.

## TODO

- Maybe it would be nice to have an output bar that gives feedback to the user about the status of the download. So it does not think the program stalled
- Support HTTP Basic Auth/FTP/SFPT/S3, maybee define subcommands for everyone of them
