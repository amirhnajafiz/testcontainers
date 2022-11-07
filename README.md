<h1 align="center">
	Test Containers
</h1>

Example of using Test containers in Golang. Creating a **redis** cluster with
containers and test redis cluster with Golang SDK while running unit tests.

We use _test-containers_ to run our dependent services like databases in a container
hense we won't need to mock them or use real external services. This makes our uint tests more
efficient and more reliable. Let's see an example.

## Redis container

All we need for running a redis container with ```docker-compose``` is
an image name and image version.

These data are all in ```pkg/storage/redis/container.go```.

```go
const (
	// redis image information
	imageName = "redis"
	imageTag  = "latest"
)
```

Now we build our container with a request.

```go
// container request
req := testcontainers.ContainerRequest{
    Image:        imageName + ":" + imageTag,
    ExposedPorts: []string{"6379/tcp"},
    WaitingFor:   wait.ForLog("Ready to accept connections"),
}
```

Now you can access container in your source codes.

```go
ctx := context.Background()

// creating a new container.
redisC, err := redis.CreateRedisContainer()
if err != nil {
    t.Error(err)

    return
}

// get container connection.
redisConnection, err := redisC.Endpoint(ctx, "")
if err != nil {
    t.Error(err)

    return
}
```

## Test

Execute redis container test with following command:

```shell
go test -v ./...
```
