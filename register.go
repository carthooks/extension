package extension

import (
	"bytes"
	"context"
	"log"
	"net"
	"os"
	"path/filepath"

	pb "github.com/carthooks/extension/service"

	"google.golang.org/grpc"
)

type Server struct {
	pb.CarthooksExtensionServer
	publicDir string
}

func (s *Server) GetResourcesGetResource(c context.Context, r *pb.GetResourceRequest) (rsp *pb.GetResourcesResponse, err error) {
	//	*GetResourcesResponse_File
	//	*GetResourcesResponse_Directory
	fpath := filepath.Join(s.publicDir, r.GetPath())

	file, err := os.Open(fpath)
	if err != nil {
		return nil, err
	}
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	if fileInfo.IsDir() {
		return s.DirResource(file)
	}
	return s.FileResource(file)
}

func (s *Server) DirResource(f *os.File) (rsp *pb.GetResourcesResponse, err error) {
	dirctory := &pb.DirectoryResponse{
		Files:       []string{},
		Directories: []string{},
	}

	files, err := f.Readdir(-1)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			dirctory.Directories = append(dirctory.Directories, file.Name())
		} else {
			dirctory.Files = append(dirctory.Files, file.Name())
		}
	}

	return &pb.GetResourcesResponse{
		Body: &pb.GetResourcesResponse_Directory{
			Directory: dirctory,
		},
	}, nil
}

func (s *Server) FileResource(f *os.File) (rsp *pb.GetResourcesResponse, err error) {
	buf := bytes.NewBuffer([]byte{})
	_, err = buf.ReadFrom(f)
	if err != nil {
		return nil, err
	}

	return &pb.GetResourcesResponse{
		Body: &pb.GetResourcesResponse_File{
			File: &pb.FileResponse{
				Body: buf.Bytes(),
			},
		},
	}, nil
}

func StartServer(addr, publicDir string, serv pb.CarthooksExtensionServer) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	s := Server{
		CarthooksExtensionServer: serv,
		publicDir:                publicDir,
	}

	grpc_server := grpc.NewServer()
	pb.RegisterCarthooksExtensionServer(grpc_server, s)
	log.Printf("server listening at %v", lis.Addr())
	if err := grpc_server.Serve(lis); err != nil {
		return err
	}
	return nil
}
