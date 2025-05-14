# GoMall 微服务电商平台

## 项目简介
GoMall是一个基于Go语言开发的微服务电商平台，采用云原生架构设计，使用Kitex框架实现RPC服务。该平台提供完整的电商业务流程，包括用户管理、商品管理、购物车、订单处理、支付系统等功能。项目采用前后端不分离的开发模式，使用Hertz框架提供Web服务。

## 技术栈
### 后端技术
- 框架：Kitex (字节跳动开源RPC框架)
- Web框架：Hertz
- 语言：Go
- 通信协议：gRPC
- 数据库：MySQL
- 缓存：Redis
- 服务发现：Consul
- 消息队列：NATS (用于邮件服务)
- 日志：Kitex Logger

### 前端技术
- 模板引擎：Go Template
- UI框架：Bootstrap
- 会话管理：Hertz Session

## 系统架构

### 核心服务
1. **用户服务 (User Service)**
   - 位置：`/app/user`
   - 接口定义：`/idl/user.proto`
   - 功能：
     - 用户注册
     - 用户登录/登出
     - JWT Token认证
     - 用户信息管理
   - 技术特点：
     - 使用JWT进行身份认证
     - 密码加密存储
     - Session管理

2. **商品服务 (Product Service)**
   - 位置：`/app/product`
   - 接口定义：`/idl/product.proto`
   - 功能：
     - 商品信息管理
     - 商品分类
     - 商品搜索
     - 商品库存管理

3. **购物车服务 (Cart Service)**
   - 位置：`/app/cart`
   - 接口定义：`/idl/cart.proto`
   - 功能：
     - 购物车创建
     - 商品添加/删除
     - 购物车信息查询
     - 购物车清空

4. **订单服务 (Order Service)**
   - 位置：`/app/order`
   - 接口定义：`/idl/order.proto`
   - 功能：
     - 订单创建
     - 订单状态管理
     - 订单查询
     - 订单定时取消

5. **支付服务 (Payment Service)**
   - 位置：`/app/payment`
   - 接口定义：`/idl/payment.proto`
   - 功能：
     - 支付处理
     - 支付状态管理
     - 订单查询
     - 模拟自动下单

6. **结账服务 (Checkout Service)**
   - 位置：`/app/checkout`
   - 接口定义：`/idl/checkout.proto`
   - 功能：
     - 订单结算
     - 价格计算
     - 优惠券处理

7. **邮件服务 (Email Service)**
   - 位置：`/app/email`
   - 接口定义：`/idl/email.proto`
   - 功能：
     - 邮件通知
     - 营销邮件
     - 使用NATS进行消息队列处理

8. **前端服务 (Frontend Service)**
   - 位置：`/app/frontend`
   - 接口定义：`/idl/frontend/*`
   - 功能：
     - 页面渲染
     - 用户认证中间件
     - 会话管理
     - API聚合

## 目录结构
```
gomall/
├── app/                    # 微服务应用目录
│   ├── user/              # 用户服务
│   │   ├── biz/          # 业务逻辑
│   │   ├── conf/         # 配置文件
│   │   └── utils/        # 工具函数
│   ├── product/          # 商品服务
│   ├── cart/             # 购物车服务
│   ├── order/            # 订单服务
│   ├── payment/          # 支付服务
│   ├── checkout/         # 结账服务
│   ├── email/            # 邮件服务
│   └── frontend/         # 前端服务
│       ├── middleware/   # 中间件
│       ├── templates/    # 页面模板
│       └── utils/        # 工具函数
├── idl/                    # 接口定义文件
│   ├── user.proto         # 用户服务接口定义
│   ├── product.proto      # 商品服务接口定义
│   ├── cart.proto         # 购物车服务接口定义
│   ├── order.proto        # 订单服务接口定义
│   ├── payment.proto      # 支付服务接口定义
│   ├── checkout.proto     # 结账服务接口定义
│   ├── email.proto        # 邮件服务接口定义
│   ├── api.proto          # 通用API定义
│   └── frontend/          # 前端服务接口定义
└── README.md              # 项目说明文档
```

## 认证与安全

### 用户认证
- 使用JWT进行身份认证
- Session管理用户状态
- 全局认证中间件
- 路由保护中间件

### 安全特性
- 密码加密存储
- Token过期机制
- Session安全控制
- 请求验证和过滤

## 开发指南

### 环境要求
- Go 1.18+
- MySQL 8.0+
- Redis 6.0+
- Consul
- NATS
- Protocol Buffers 编译器
- Kitex 工具链

### 快速开始
1. 克隆项目
```bash
git clone [项目地址]
cd gomall
```

2. 安装依赖
```bash
go mod tidy
```

3. 配置环境变量
```bash
# 创建.env文件
cp .env.example .env
# 编辑.env文件，配置必要的环境变量
```

4. 启动服务
```bash
# 启动用户服务
go run app/user/main.go

# 启动前端服务
go run app/frontend/main.go

# 启动其他服务
go run app/[service_name]/main.go
```

### 服务生成
1. 生成新服务
```bash
bash gen_service.sh servicename
```

2. 生成前端页面
```bash
bash gen_frontend.sh service_pagename
```

## 部署说明

### 开发环境
- 使用Consul进行服务发现
- 本地Redis缓存
- 本地MySQL数据库
- 本地NATS消息队列

### 生产环境
- 使用Consul集群
- Redis集群
- MySQL主从
- NATS集群
- 负载均衡
- 监控告警

## 错误处理
项目使用统一的错误处理机制：
- 使用`utils.MustHandleError`进行错误处理
- 统一的错误码定义
- 详细的错误日志记录
- 优雅的错误恢复机制

## 贡献指南
1. Fork 项目
2. 创建特性分支
3. 提交变更
4. 推送到分支
5. 创建 Pull Request

## 许可证
[待定]

## 联系方式
[待定]
