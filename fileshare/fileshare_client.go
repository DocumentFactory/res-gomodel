package fileshare

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"

	"github.com/pnocera/res-gomodel/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type FileshareClient struct {
	service pb.FileshareServiceClient
}

func NewFileshareClient(cc *grpc.ClientConn) *FileshareClient {
	service := pb.NewFileshareServiceClient(cc)
	return &FileshareClient{service}
}

// func (c *FileshareClient) DownloadFile(req *pb.DownloadFileRequest, filename string) error {
// 	stream, err := c.service.DownloadFile(context.Background(), req)
// 	if err != nil {
// 		return err
// 	}

// 	return streamToFile(stream, filename)

// }

// func (c *FileshareClient) DownloadStream0(ctx context.Context, req *pb.DownloadFileRequest, iowriter io.Writer) error {

// 	log.Printf("DownloadStream: %v", req)

// 	stream, err := c.service.DownloadFile(ctx, req)
// 	if err != nil {
// 		return err
// 	}
// 	writer := bufio.NewWriter(iowriter)

// 	defer writer.Flush()

// 	size := 0

// 	for {
// 		res, err := stream.Recv()
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			log.Printf("DownloadStream stream.Recv error: %v", err)
// 			stream.CloseSend()
// 			return fmt.Errorf("cannot receive chunk from server %v", err)
// 		}
// 		bytes := res.GetChunkData()
// 		_, err = writer.Write(bytes)
// 		size = size + len(bytes)
// 		if err != nil {
// 			log.Printf("DownloadStream writer.write error: %v", err)
// 			stream.CloseSend()
// 			return fmt.Errorf("cannot write chunk to file %v", err)
// 		}
// 	}
// 	log.Printf("DownloadStream downloaded : %s", NiceByteSize(size))
// 	return stream.CloseSend()

// }

func (c *FileshareClient) DownloadStream(ctx context.Context, req *pb.DownloadFileRequest, iowriter io.Writer) error {

	stream, err := c.service.DownloadFile(ctx, req)
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(iowriter)

	defer writer.Flush()

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("cannot receive chunk from server %v", err)
		}

		_, err = writer.Write(res.GetChunkData())
		if err != nil {
			return fmt.Errorf("cannot write chunk to file %v", err)
		}
	}
	return nil

}

func (c *FileshareClient) UploadStream(ctx context.Context, req *pb.UploadFileRequest, ioreader io.Reader) (*pb.UploadFileResponse, error) {

	log.Printf("UploadStream: %v", req)

	metadata.AppendToOutgoingContext(ctx, "dapr-app-id", "fileshare", "dapr-stream", "true")

	stream, err := c.service.UploadFile(ctx)
	if err != nil {
		return nil, err
	}

	err = stream.Send(req)
	if err != nil {
		return nil, err
	}

	reader := bufio.NewReader(ioreader)
	buffer := make([]byte, 1024)

	size := 0

	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("cannot read chunk from file %v", err)
		}
		size = size + n

		req := &pb.UploadFileRequest{
			Data: &pb.UploadFileRequest_ChunkData{
				ChunkData: buffer[:n],
			},
		}

		err = stream.Send(req)
		if err != nil {
			return nil, fmt.Errorf("cannot send chunk to server %v %v", err, stream.RecvMsg(nil))
		}
	}
	log.Printf("UploadStream uploaded : %s", NiceByteSize(size))

	return stream.CloseAndRecv()

}

// func (c *FileshareClient) UploadFile(req *pb.UploadFileRequest, filename string) (*pb.UploadFileResponse, error) {

// 	stream, err := c.service.UploadFile(context.Background())
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = stream.Send(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return fileToStream(stream, filename)

// }

// func fileToStream(stream pb.FileshareService_UploadFileClient, filename string) (*pb.UploadFileResponse, error) {
// 	file, err := os.Open(filename)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer file.Close()

// 	reader := bufio.NewReader(file)
// 	buffer := make([]byte, 64*1024)

// 	for {
// 		n, err := reader.Read(buffer)
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			return nil, fmt.Errorf("cannot read chunk from file %v", err)
// 		}

// 		req := &pb.UploadFileRequest{
// 			Data: &pb.UploadFileRequest_ChunkData{
// 				ChunkData: buffer[:n],
// 			},
// 		}

// 		err = stream.Send(req)
// 		if err != nil {
// 			return nil, fmt.Errorf("cannot send chunk to server %v %v", err, stream.RecvMsg(nil))
// 		}
// 	}

// 	return stream.CloseAndRecv()
// }

// func streamToFile(stream pb.FileshareService_DownloadFileClient, filename string) error {
// 	file, err := os.Create(filename)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	writer := bufio.NewWriter(file)

// 	defer writer.Flush()

// 	for {
// 		res, err := stream.Recv()
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			return fmt.Errorf("cannot receive chunk from server %v", err)
// 		}
// 		bytes := res.GetChunkData()
// 		if len(bytes) > 0 {
// 			_, err = writer.Write(bytes)
// 			if err != nil {
// 				_ = stream.CloseSend()
// 				return fmt.Errorf("cannot write chunk to file %v", err)
// 			}
// 		}

// 		// _, err = writer.Write(res.GetChunkData())
// 		// if err != nil {
// 		// 	return fmt.Errorf("cannot write chunk to file %v", err)
// 		// }
// 	}
// 	return nil
// }

func NiceByteSize(size int) string {
	if size < 1024 {
		return fmt.Sprintf("%d bytes", size)
	} else if size < 1024*1024 {
		return fmt.Sprintf("%.2f KB", float64(size)/1024)
	} else if size < 1024*1024*1024 {
		return fmt.Sprintf("%.2f MB", float64(size)/1024/1024)
	} else {
		return fmt.Sprintf("%.2f GB", float64(size)/1024/1024/1024)
	}
}
