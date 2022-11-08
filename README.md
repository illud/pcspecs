# PcSpecs

## Get system info

\
PcSpecs is a package to gather the system specs.


## Features
- Hostname
- Platform
- OsNumber
- CPU
- GPU
- RAM
- Disk
- MAINBOARD

## Installation

Gojira requires [Go](https://golang.org/) v1.11+ to run.

Install the dependencies.

```sh
go get github.com/saturnavt/pcspecs
```
Or

```sh
go install github.com/saturnavt/pcspecs@latest
```
## How to use

Import:

```go
pcs "github.com/saturnavt/pcspecs"
```

Example:

```go
package main

import pcs "github.com/saturnavt/pcspecs"

func main(){
    fmt.Println(pcs.Spec())
    /* Output
    DESKTOP-4DA5M71 
    Microsoft Windows 10
    Pro 10 
    Intel(R) Core(TM) i7-10700K CPU @ 3.80GHz
    NVIDIA GeForce RTX 2060 SUPER
    24 
    134
    TUF GAMING B460M-PLUS (WI-FI*/
}
```

Or

```go
package main

import pcs "github.com/saturnavt/pcspecs"

func main(){
    fmt.Println(pcs.Spec().GPU)
    /* Output
    NVIDIA GeForce RTX 2060 SUPER
    */
}
```

## License

MIT

Gojira is [MIT licensed](LICENSE).