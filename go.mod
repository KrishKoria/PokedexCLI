module github.com/KrishKoria/PokedexCli

go 1.23.5

replace github.com/KrishKoria/PokeApi v0.0.0 => ./internal/api
replace github.com/KrishKoria/PokeCache v0.0.0 => ./internal/cache

require (
github.com/KrishKoria/PokeApi v0.0.0
github.com/KrishKoria/PokeCache v0.0.0 //indirect
) 
