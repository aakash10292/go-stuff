## Exporting

Packages contain the basic unit of compiled code. They define a scope for the identifiers that are declared within them. Exporting is not the same as public and private semantics in other languages. But exporting is how we provide encapsulation in Go.

## Notes

* Code in go is compiled into packages and then linked together.
* Identifiers are exported (or remain unexported) based on letter-case.
* We import packages to access exported identifiers.
* Any package can use a value of an unexported type, but this is annoying to use.
