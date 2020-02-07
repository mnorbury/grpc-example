all: grpc-example

clean:
	@rm -f grpc-example calculatorpb/calculator.pb.go

grpc-example: main.go calculatorpb/calculator.pb.go
	go build -o grpc-example main.go

calculatorpb/calculator.pb.go: proto/calculator.proto
	protoc -I proto/ --go_out=plugins=grpc:calculatorpb/ proto/calculator.proto
