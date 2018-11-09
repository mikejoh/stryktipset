# Stryktipset CLI

Serves the `stryktipset` package through a command-line tool called `st`.

### Compile and run

Run `make build` to compile the binary, use `make build-linux` for x-compiling to Linux based OS:es. To compile the binary for Linux through Docker run: `make build-docker-linux`. See the `Makefile` for more options.

Example of running (in working directory) `st`:
```
$> ./st -sek 192
Full: 1
Half: 6
```
