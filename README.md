# Learn Go with Tests

https://quii.gitbook.io/learn-go-with-tests

Snippets of testable Go code, commented and
ready to be copy/pasted around.

## Basics
* Test files are named _xxx_test.go_, eg.: hello_test.go
* Test cases are prefixed with _Test_, eg.: TestHello(t *testing.T) 
* Run tests by typing: `go test`. (Single test: CTRL+SHIFT+F10)
* Offline docs: `go doc fmt` or download pkgsite:
```
go install golang.org/x/pkgsite/cmd/pkgsite@latest
pkgsite -open .
```

## Refactoring and your tooling

Big emphasis is the importance of refactoring and
trusting your tooling to help.

You should be familiar enough with your editor to
perform the following with a simple key combination:
* **Extract/Inline variable.** Taking magic values and 
giving them a name lets you simplify your code quickly.

* **Extract method/function.** It is vital to be able 
to take a section of code and extract functions/methods.

* **Rename.** You should be able to rename symbols 
across files confidently.

* **go fmt.** Go has an opinioned formatter called go fmt.
Your editor should run this on every file saved.

* **Run tests.** You should be able to do any of the above
and then quickly re-run your tests to ensure your
refactoring hasn't broken anything.

In addition, to help you work with your code, you should be able to:

* **View function signature.** You should never be unsure how to 
call a function in Go. Your IDE should describe a function in
terms of its documentation, its parameters and what it returns.

* **View function definition.** If it's still unclear what a
function does, you should be able to jump to the source code and
try and figure it out yourself.

* **Find usages of a symbol.** Understanding a function's
context can help you make decisions when refactoring.

* **Jump to terminal**. ALT+F12

* **Comment block**. (custom CTRL+SHIFT+L)

Mastering your tools will help you concentrate on the code and reduce context switching.
