package greetings

import (
	"fmt"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

// HelloWorld prints "Hello, World!" to the standard output.
func HelloWorld() {
    span := tracer.StartSpan("sample", tracer.ResourceName("HelloWorld"))
    span.SetTag("test", "masafumi")
    defer span.Finish()
	
    fmt.Println("Hello, World!")
}
