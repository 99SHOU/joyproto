rd cs /s /q
rd go /s /q
rd grpc /s /q

md cs
md go
md grpc

echo gen C# file
cd ./proto
protogen --csharp_out=../cs define.proto game_msg.proto framework_msg.proto
cd ..

echo gen Go file
cd ./proto
protoc --go_out=../go define.proto game_msg.proto framework_msg.proto
cd ..

call copy.bat


pause