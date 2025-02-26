# GRPC

```sh
cmd
    - api
        - main.go
config
    - local.yaml
internal | pkg
    - config
        - config.go
    - http
        - handlers
            - posts
                - posts.go
    - types
        - types.go
```

> [**cleanenv**](https://github.com/ilyakaznacheev/cleanenv)  
> [**graceful shutdown**](https://www.freecodecamp.org/news/graceful-shutdowns-k8s-go/)  
> [**structured logging**](https://betterstack.com/community/guides/logging/logging-in-go/)  
> [**dependecny injection**](https://stackoverflow.com/questions/41900053/is-there-a-better-dependency-injection-pattern-in-golang)
>
> [**go-blueprint**](https://github.com/Melkeydev/go-blueprint)

```sh
syntax = "proto3";

package hello;

message HelloRequest {
  string name = 1;
}
message HelloReply {
  string message = 1;
}

service Greeter {
  rpc SayHello (HelloRequest) returns (HelloReply) {}
}
```

```go
package main

import (
    "context"
    "log"
    "net"

    "google.golang.org/grpc"
    pb "path/to/your/protobuf/package"
)

type server struct {
    pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
    return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
    lis, _ := net.Listen("tcp", ":50051")

    s := grpc.NewServer()
    pb.RegisterGreeterServer(s, &server{})

    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
```

```go
package main

import (
    "context"
    "log"
    "time"

    "google.golang.org/grpc"
    pb "path/to/your/protobuf/package"
)

func main() {
    conn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
    defer conn.Close()

    c := pb.NewGreeterClient(conn)

    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    r, _ := c.SayHello(ctx, &pb.HelloRequest{Name: "world"})
    log.Printf("Greeting: %s", r.GetMessage())
}
```

> [**When RESTful architecture isn't enough..**.](https://www.youtube.com/watch?v=_4TPM6clQjM)  
> [**Complete Golang and gRPC Microservices**](https://www.youtube.com/watch?v=ea_4Ug5WWYE)  
> [**How to Implement Server-Side Streaming (w/ Go, gRPC, and Redis)**](https://www.youtube.com/watch?v=ZK4baSgQ1ks)  
> [**2024 gRPC Golang Tutorial - The tutorial I wish I had when I was learning**](https://www.youtube.com/watch?v=mPESsBfUKkc)

---

> [**Documentation**](https://grpc.io/docs/languages/go/basics/)  
> [**Golang gRPC In-Depth Tutorial**](https://www.golinuxcloud.com/golang-grpc/)  
> [**Interceptor into Golang gRPC server**](https://medium.com/@mahes0/adding-interceptor-into-golang-grpc-server-0a5ea4d12f27)  
> [**GRPC with Auth Interceptor, Streaming and Gateway**](https://dev.to/truongpx396/golang-grpc-with-auth-interceptor-streaming-and-gateway-in-practice-24b8)  
> [**A Complete Guide to Implement gRPC using Golang**](https://reliasoftware.com/blog/golang-grpc)
