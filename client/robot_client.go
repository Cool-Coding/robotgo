package main

import (
	"context"
	"fmt"
	"github.com/go-vgo/robotgo"
	"google.golang.org/grpc"
	"log"
	pb "robotgo/service"
	"time"
)

func main() {
	recvSize := 10 * 1024 * 1024
	// 通过 grpc.Dial 获得一条连接
	conn, err := grpc.Dial("127.0.0.1:12345", grpc.WithInsecure(), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(recvSize)))
	// 如果要增加 Recv 可以接受的一个消息的数据量，必须增加 grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(100000000))
	//conn, err := grpc.Dial("unix:///var/lib/test.socket", grpc.WithInsecure(), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(100000000)))
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewGoRobotClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10000*time.Second)
	defer cancel()

	//point := pb.Point{X:800,Y:20}
	//point2 := pb.Point{X:800,Y:100}
	//mouseButton := pb.MouseButton{X:1024}

	//_, err = client.MouseMove(ctx, &point)
	//_, err = client.MousePressed(ctx, &mouseButton)
	//_, err = client.MouseMove(ctx, &point2)
	//_, err = client.MouseReleased(ctx, &mouseButton)

	request :=pb.CaptureScreenRequest{
		Width:  0,
		Height: 0,
	}
	screen, err := client.CaptureScreen(ctx,&request)
	if err != nil {
		fmt.Printf("error %v", err)
		return
	}

	data, err :=screen.Recv()
	if err != nil {
		fmt.Printf("error %v", err)
		return
	}
	fmt.Printf("%v", data)
	robotgo.SaveImg(data.Bitmap,"E:/robot.png")
}
