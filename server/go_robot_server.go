package server

import (
	context "context"
	"fmt"
	"github.com/go-vgo/robotgo"
	"google.golang.org/grpc"
	"log"
	"math"
	"net"
	pb "robotgo/service"
)

type robotServer struct{}


func (x *robotServer) MouseDoubleClick(ctx context.Context,mouseButton *pb.MouseButton) (*pb.Empty, error) {
	robotgo.MouseClick(getMouseButton(mouseButton.X), true)
	return &pb.Empty{}, nil
}

func (x *robotServer) MousePressed(ctx context.Context,mouseButton *pb.MouseButton) (*pb.Empty, error) {
	robotgo.MouseToggle("down",getMouseButton(mouseButton.X))
	log.Printf("mouse down: %d", mouseButton.X)
	return &pb.Empty{}, nil
}

func (x *robotServer) MouseReleased(ctx context.Context, mouseButton *pb.MouseButton) (*pb.Empty, error) {
	robotgo.MouseToggle("up",getMouseButton(mouseButton.X))
	log.Printf("mouse up: %d", mouseButton.X)
	return &pb.Empty{}, nil
}

func (x *robotServer) MouseScrolledUp(ctx context.Context, mouseScrolledRequest *pb.MouseScrolledRequest) (*pb.Empty, error) {
	robotgo.ScrollMouse(int(mouseScrolledRequest.Distance), "up")
	return &pb.Empty{}, nil
}

func (x *robotServer) MouseScrolledDown(ctx context.Context, mouseScrolledRequest *pb.MouseScrolledRequest) (*pb.Empty, error) {
	robotgo.ScrollMouse(int(mouseScrolledRequest.Distance), "down")
	return &pb.Empty{}, nil
}

func (x *robotServer) MouseMove(ctx context.Context, point *pb.Point) (*pb.Empty, error) {
	robotgo.MoveMouse(int(point.X), int(point.Y))
	log.Printf("mouse move: %d,%d", point.X,point.Y)
	return &pb.Empty{}, nil
}

func (x *robotServer) MouseClick(ctx context.Context, mouseButton *pb.MouseButton) (*pb.Empty, error) {
	robotgo.MouseClick(getMouseButton(mouseButton.X), false)
	return &pb.Empty{}, nil
}

func (*robotServer) CaptureScreen(captureScreenRequest *pb.CaptureScreenRequest, cps pb.GoRobot_CaptureScreenServer) error {
	width := captureScreenRequest.Width
	height := captureScreenRequest.Height
	var err error

	log.Printf("caputure screen")
	if width == 0 || height == 0 {
		bitmap := robotgo.CaptureScreen()
		defer robotgo.FreeBitmap(bitmap)
		bytes := robotgo.ToBitmapBytes(bitmap)
		err = sendImage(bytes, err, cps)
	} else {
		bitmap := robotgo.CaptureScreen(0, 0, int(captureScreenRequest.Width), int(captureScreenRequest.Height))
		defer robotgo.FreeBitmap(bitmap)
		bytes := robotgo.ToBitmapBytes(bitmap)
		err = sendImage(bytes,err,cps)
	}
	return err
}

func sendImage(bytes []byte, err error, cps pb.GoRobot_CaptureScreenServer) error {



	length := len(bytes)
	packLength := 1024 * 1024
	startIndex := 0

	for {
		endIndex := startIndex + packLength
		if endIndex > length {
			endIndex = length
		}

		err = cps.Send(&pb.CaptureScreenResponse{Bitmap: bytes[startIndex:endIndex]})

		startIndex = endIndex
		if startIndex >= length {
			break
		}
	}
	return err
}

func getMouseButton(value uint32) string {
	mouseButton := ""
	switch value {
	case 1024:
		mouseButton = "left"
		break
	case 2048:
		mouseButton = "center"
		break
	case 4096:
		mouseButton = "right"
		break
	default:
	}

	return mouseButton
}

func Run_Robot(ip_port string) error {
	sock, err := net.Listen("tcp", fmt.Sprintf("%s",ip_port))
	if err != nil {
		return err
	}

	var options = []grpc.ServerOption{
		grpc.MaxRecvMsgSize(math.MaxInt32),
		grpc.MaxSendMsgSize(1073741824),
	}
	s := grpc.NewServer(options...)

	myServer := &robotServer{}
	/*
	   见 go_robot.pb.go 中
	func RegisterGoRobotServer(s *grpc.Server, srv GoRobotServer) {
		s.RegisterService(&_GoRobot_serviceDesc, srv)
	}
	   前者是 grpc server，后者是实现了 GoRobotServer 所有 interface 的实例，即 &robotServer{}

	   感觉就是将 grpc server 和 handler 绑定在了一起的意思。

	   RegisterGoRobotServer 的命名很有意思，Register(固定) + MouseXXX(service MouseXXX {} in proto 文件) + Server(固定)

	*/
	pb.RegisterGoRobotServer(s, myServer)
	if err != nil {
		return err
	}

	/*
	   在 s.Serve(sock) 上监听服务

	*/
	if err := s.Serve(sock); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	return nil
}