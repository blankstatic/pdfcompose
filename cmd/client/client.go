package main

import (
	"context"
	"io/ioutil"
	"log"

	pdfcompose "github.com/blankstatic/pdfcompose/pkg/service-pdf-compose/pdfcomposeservice"

	"google.golang.org/grpc"
)

const (
	filename = "test.png"
	output   = "output.pdf"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	client := pdfcompose.NewPDFComposeServiceClient(conn)

	fileContent, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	request := &pdfcompose.FileRequest{
		FileName:    filename,
		FileContent: fileContent,
	}

	stream, err := client.SendFile(context.Background())
	if err != nil {
		log.Fatalf("Failed to send file: %v", err)
	}

	if err := stream.Send(request); err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Failed to receive response: %v", err)
	}

	err = ioutil.WriteFile(output, response.PdfFile, 0644)
	if err != nil {
		log.Fatalf("Failed to write PDF file: %v", err)
	}

	log.Println("PDF file saved successfully")
}
