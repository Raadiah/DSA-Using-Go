module github.com/Raadiah/DSA-Using-Go

go 1.23.4

retract (
    v0.1.0 // Test Publish.
    v0.2.0 // Test Publish.
    v0.3.0 // Test Publish.
    v0.4.0 // Test Publish.
    v0.6.0 // Test Publish. All the previous versions were added to test how this works.
)

require github.com/fatih/color v1.18.0

require (
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	golang.org/x/sys v0.25.0 // indirect
)
