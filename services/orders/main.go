package main

func main() {

	// run the grpc server ...
	grpcServer := NewGRPCServer(":9000")
	grpcServer.Run()
}
