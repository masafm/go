package main

import (
	"go_test/greetings"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func main() {
	tracer.Start(
        tracer.WithEnv("host"),
        tracer.WithService("go-test"),
        tracer.WithServiceVersion("1.0"),
		tracer.WithLogStartup(true),
		tracer.WithDebugMode(true),
    )
	defer tracer.Stop()

	span := tracer.StartSpan("web.request", tracer.ResourceName("/posts"))
    // Set tag
    span.SetTag("test", "masafumi")
	defer span.Finish()

	greetings.HelloWorld()
}
