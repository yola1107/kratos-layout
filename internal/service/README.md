# Service


## 三层职责与依赖关系
```text
[ service ]         // 请求适配层（HTTP/gRPC）
      ↓ 调用
[  biz   ]          // 业务逻辑层（Usecase + 接口定义）
      ↑ 依赖接口
[  data  ]          // 数据实现层（接口实现 + DB/缓存）
```

## 解耦的机制
```
1>接口依赖反转
    biz 定义接口（如 GreeterRepo）
    data 实现接口
    最终由 wire 注入到 biz
    
    // biz/greeter.go
    type GreeterRepo interface {
        Save(ctx context.Context, g *Greeter) error
    }
    
2>wire 依赖注入

3>结构体注入而不是 new() 构造
```

## 目录结构
```
app/
├── service/        // gRPC / HTTP 服务层
├── biz/            // Usecase + 接口定义
├── data/           // 接口实现，数据库操作

```


## service 层：
```
1>负责接收请求、参数校验、调用业务逻辑
2>一般对应 controller / handler（如 HTTP、gRPC）
3>仅负责路由、调用 biz.Usecase
4>不直接操作数据或数据库

type GreeterService struct {
    uc *biz.GreeterUsecase
}
```


## biz 层：核心业务逻辑（领域层）
```
1>包含领域实体、用例逻辑
2>依赖抽象的 repository 接口，而不关心具体的数据实现
3>由 data 层提供实现并注入

type GreeterUsecase struct {
    repo GreeterRepo  // 是接口
}
```

## data 层：实现数据访问逻辑
```
1>实现 biz 层定义的接口
2>依赖数据库驱动、缓存库、三方接口等
3>完全被 biz 所抽象

type greeterRepo struct {
    data *Data  // 封装 DB/Redis 等客户端
}

func (r *greeterRepo) Save(ctx context.Context, g *biz.Greeter) error {
    // 具体实现
}

```


