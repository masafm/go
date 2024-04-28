package main

import (
	"go_test/greetings"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	httptrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"
	"fmt"
	"net/http"
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

	span := tracer.StartSpan("sample", tracer.ResourceName("main"))
    span.SetTag("test", "masafumi")
	defer span.Finish()

	// URLを設定
	url := "https://www.datadoghq.com/ja/"
	// http.Clientのインスタンスを作成
	client := &http.Client{}
	traceClient := httptrace.WrapClient(client)
	
	// HEADリクエストを実行
	resp, err := traceClient.Head(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close() // 必ずクローズする

	// ヘッダーを取得し、標準出力に出力する
	for key, values := range resp.Header {
		for _, value := range values {
			fmt.Printf("%s: %s\n", key, value)
		}
	}
	
	greetings.HelloWorld()
}
