stdin
=====

**stdin** is a simple utility to make it even easier to read in lines of text in Golang. 

Requires Go 1.1.

Usage:

	import "github.com/earthboundkid/stdin"
	
	func f() {
		for line := range stdin.Lines() {
			//Do something.
		}
	}