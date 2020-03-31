客服系统开发者QQ交流群： 623661658

# 开源客服系统-机器人 v2.0.0
客服系统仓库地址： [前往>>>][1]

## 以下是v2.0.0版本的重要更新
- 对前面版本进行重构，分离业务逻辑与机器人的混搭运行弊端
- 做了大量优化，在很大程度上提升了性能，并代码松耦合，
- 业务系统支持负载均衡了，这是对v1.0的重大里程碑更新, 对接海量客户不再愁了
- 增加工单系统能力，无在线客服接待？不用怕，工单来给您解决一切问题
- 代码可读性大大提高，初学者都能看懂的代码，还有什么理由不学习一下呢
- 定时清理无接入人工记录的用户，避免数据沉淀
- H5客户端增加了重连机制
- 客户端只保留30天聊天记录，已分表处理

## 目录文件结构说明
|       |     |
| :-------- | :-------- |
| - conf  | 项目配置文件 |
| - grpcc  | rpc链接层 |
| - robot  | 机器人相关操作 |
| - services  | 服务提供者-通过rpc调取服务端获取数据 |
| - processguard_robot.sh | Linux 进程守护shell脚本 |


## 安装
- cd $GOPATH/src && git clone https://github.com/chenxianqi/kefu_go_robot
- cd kefu_go_robot && go get

## 启动
go run main.go


## 打包发布 linux (其它运行环境编译请自行search baidu)
  - CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
	编译完成将生成一个、main 可执行文件，改名成 kefu_go_robot
	新建一个robot目录将编译后改名的kefu_go_robot文件放进去，然后在robot目录创建一个conf文件夹，将conf/conf.yaml文件拷贝一份进来，最后将processguard_robot.sh可执行文件也拷贝进来，robot项目就打包完成了
	
- 启动：先启动服务端后 cd robot && nohub ./processguard_robot.sh &


## LICENSE

Copyright 2019 keith

Licensed to the Apache Software Foundation (ASF) under one or more contributor
license agreements.  See the NOTICE file distributed with this work for
additional information regarding copyright ownership.  The ASF licenses this
file to you under the Apache License, Version 2.0 (the "License"); you may not
use this file except in compliance with the License.  You may obtain a copy of
the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.  See the
License for the specific language governing permissions and limitations under
the License.



  [1]: https://github.com/chenxianqi/kefu_server 
  
