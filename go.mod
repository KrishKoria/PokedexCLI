module github.com/KrishKoria/PokedexCli

go 1.23.5

replace github.com/KrishKoria/PokeApi v0.0.0 => ./internal/api

replace github.com/KrishKoria/PokeCache v0.0.0 => ./internal/cache

require github.com/KrishKoria/PokeApi v0.0.0

require (
	github.com/KrishKoria/PokeCache v0.0.0 //indirect
	github.com/mattn/go-runewidth v0.0.3 // indirect
	github.com/peterh/liner v1.2.2 
	golang.org/x/sys v0.0.0-20211117180635-dee7805ff2e1 // indirect
)
