package extension

import (
	"bytes"
	"context"
	"log"
	"net"
	"os"
	"path/filepath"

	pb "github.com/carthooks/extension/service"
	"github.com/gin-gonic/gin"

	"google.golang.org/grpc"
)

type Server struct {
	pb.CarthooksExtensionServer
	publicDir     string
	HttpGinEngine *gin.Engine
}

func (s Server) GetAssets(c context.Context, r *pb.GetAssetsRequest) (rsp *pb.GetAssetsResponse, err error) {
	//	*GetAssetsResponse_File
	//	*GetAssetsResponse_Directory
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
		return s.DirAssets(file)
	}
	return s.FileAssets(file)
}

func (s *Server) DirAssets(f *os.File) (rsp *pb.GetAssetsResponse, err error) {
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

	return &pb.GetAssetsResponse{
		Body: &pb.GetAssetsResponse_Directory{
			Directory: dirctory,
		},
	}, nil
}

func (s *Server) FileAssets(f *os.File) (rsp *pb.GetAssetsResponse, err error) {
	buf := bytes.NewBuffer([]byte{})
	_, err = buf.ReadFrom(f)
	if err != nil {
		return nil, err
	}

	return &pb.GetAssetsResponse{
		Body: &pb.GetAssetsResponse_File{
			File: &pb.FileResponse{
				Body: buf.Bytes(),
			},
		},
	}, nil
}

type Config struct {
	Addr          string
	PublicDir     string
	Server        pb.CarthooksExtensionServer
	HttpGinEngine *gin.Engine
}

func StartServer(config *Config) error {
	lis, err := net.Listen("tcp", config.Addr)
	if err != nil {
		return err
	}

	s := Server{
		CarthooksExtensionServer: config.Server,
		publicDir:                config.PublicDir,
		HttpGinEngine:            config.HttpGinEngine,
	}

	grpc_server := grpc.NewServer()
	pb.RegisterCarthooksExtensionServer(grpc_server, s)
	log.Printf("server listening at %v", lis.Addr())
	if err := grpc_server.Serve(lis); err != nil {
		return err
	}
	return nil
}
