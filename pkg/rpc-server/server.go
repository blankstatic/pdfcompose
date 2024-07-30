package server

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"

	"github.com/blankstatic/pdfcompose/pkg/composer"

	pdfcompose "github.com/blankstatic/pdfcompose/pkg/service-pdf-compose/pdfcomposeservice"

	"google.golang.org/grpc"
)

type Server struct {
	pdfcompose.UnimplementedPDFComposeServiceServer
}

func (s *Server) SendFile(stream pdfcompose.PDFComposeService_SendFileServer) error {
	var fileReaders []io.ReadCloser

	for {
		fileRequest, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		fileReader := io.NopCloser(bytes.NewReader(fileRequest.GetFileContent()))
		fileReaders = append(fileReaders, fileReader)

		fmt.Printf("Received file: %s\n", fileRequest.GetFileName())
	}

	composedFile, err := composer.ComposeFromFiles(fileReaders)
	if err != nil {
		return err
	}

	response := &pdfcompose.FileResponse{}

	response.PdfFile, err = ioutil.ReadAll(composedFile)
	if err != nil {
		return err
	}

	if err := stream.SendAndClose(response); err != nil {
		return err
	}

	return nil
}

func RunRPCServer() error {
	grpcServer := grpc.NewServer()
	pdfcompose.RegisterPDFComposeServiceServer(grpcServer, &Server{})

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}

	log.Println("gRPC server is running on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		return err
	}
	return nil
}
