# go-minis
Mini Go programs, and useful Go patterns.

All contributions & suggestions are appreciated :)

## TODOs
* example of naming & unamming inherited structs
* Exported fields in Unexported struct
* recieve only channels (<-chan)
* example of my limiter
* example of getting on a closed channel
* returning stuff in defer (like errors)
* recovering from errors
* methods on types other than struct
* new keyword
* const map
* distinctions in defer: passing var OR using var from parent func scope
* error shadowing: named err error in func definition & redeclared err error in if statement that returns (also use defer to show true value with prints in caling func)
* switch patterns (https://yourbasic.org/golang/switch-statement/)
* check if type implements method (https://cs.opensource.google/go/go/+/refs/tags/go1.17.3:src/errors/wrap.go;l=14)
* demonstrate how global vars initialization happens before init funcs
* go slice tricks (https://ueokande.github.io/go-slice-tricks/)
* intentional garbage collection
* workers pool (multiple ways)
* an example of how one interface variable can be assigned to another interface variable

## Kinda unrelated TODOs
* benchmark comparing errors.New & fmt.Errorf
* comparing memory size of `chan bool` and `chan struct{}` or `chan interface{}`

