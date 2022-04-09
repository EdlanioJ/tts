package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/EdlanioJ/tts/infra/grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	var t, l string
	file := flag.String("file", "", "File to read text from *.txt")
	backend := flag.String("backend", "localhost:50051", "Address of the grpc server")
	text := flag.String("text", "", "Text to be spoken")
	language := flag.String("lang", "pt-BR", "Language to be spoken")
	output := flag.String("out", "", "mp3 file where the output will be written")
	help := flag.Bool("help", false, "Show help")

	flag.Parse()
	if *text == "" && *file == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if *text != "" {
		t = *text
		l = *language
	}

	if *file != "" {
		b, err := ioutil.ReadFile(*file)
		if err != nil {
			errorf("Error reading file: %v", err)
		}

		fmt.Println(b)
		mimeType := http.DetectContentType(b)
		if !strings.HasPrefix(mimeType, "text/plain") {
			errorf("Invalid file type: %s", mimeType)
		}
		t = string(b)
		l = *language
	}

	if *output == "" {
		*output = fmt.Sprintf("AUDIO-%v.mp3", time.Now().Unix())
	}

	conn, err := grpc.Dial(*backend, grpc.WithInsecure())
	if err != nil {
		errorf("Could not connect to grpc server: %v", err)
	}

	defer conn.Close()

	client := pb.NewTextToSpeechClient(conn)
	req := &pb.Request{
		Text:     t,
		Language: l,
	}
	response, err := client.Say(context.Background(), req)
	if err != nil {
		errorf("could not say: %v", err)
	}

	if err := ioutil.WriteFile(*output, response.Audio, 0666); err != nil {
		errorf("could not write file: %v", err)
	}
}

func errorf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(2)
}
