echo copy file

del ..\joyserver\common\pb\*.go
del ..\joyclient\Assets\scripts\pb\*.cs

copy .\cs\*.cs ..\joyclient\Assets\scripts\pb
copy .\go\*.go ..\joyserver\common\pb
copy .\grpc\*.go ..\joyserver\common\pb