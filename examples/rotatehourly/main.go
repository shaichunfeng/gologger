package main

import (
	"log"
	"time"

	"github.com/shaichunfeng/gologger"
	"github.com/shaichunfeng/gologger/writer"
)

func main() {
	logsOptions := writer.DefaultFileWithRotationOptions
	logsOptions.Rotate = true
	logsOptions.Compress = true
	logsOptions.RotateEachHour = true
	filewriterWithRotation, err := writer.NewFileWithRotation(&logsOptions)
	if err != nil {
		log.Fatal(err)
	}
	gologger.DefaultLogger.SetWriter(filewriterWithRotation)
	for {
		time.Sleep(100 * time.Millisecond)
		gologger.Print().Msgf("%s\n", time.Now().String())
	}
}
