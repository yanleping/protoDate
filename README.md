# protoDate
creative、campaign数据从mongo读取到redis中（proto)

# 执行步骤
1、将改项目代码拉下来后，将文件protobuf、redis、util三个包以及flushCampaignpb.go、flushcreativepb.go拷贝到 
$GOPATH/src目录下

2、执行命令：go build flushcreativepb.go, 会在当前目录生成一个 flushcreativepb 可执行文件。

3、执行命令：./flushcreativepb -limit=2000

建议：在执行这个命令前:1、先测试当前测试机是否可连接远程的mongo  2、测试当前测试机的有没有安装redis,且有没有写的权限。

