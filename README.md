# tcpweather.go

I was just messing around with reading from and writing to TCP connections.

This was some R&D for [Project Divinity](https://github.com/HDN-1D10T/divinity), where I will probably implement this type of configuration
to replace the current implementation of the Google's goexpect package and spawning a native telnet client, which will kill 2 birds with 1 stone
by knocking out a couple of dependencies and allowing full TCP integration with The Divinity Framework for macOS because of breaking functionality
caused by the goexpect package. 
