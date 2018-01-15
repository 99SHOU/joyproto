rd cs /s /q
rd go /s /q
rd grpc /s /q

md cs
md go
md grpc

echo gen C# file
cd ./proto
protogen --csharp_out=../cs EventDefine.proto MsgDefine.proto
cd ..

echo gen Go file
cd ./proto
protoc --go_out=../go EventDefine.proto MsgDefine.proto
protoc --go_out=plugins=grpc:../grpc grpc_test.proto
cd ..

call copy.bat


pause