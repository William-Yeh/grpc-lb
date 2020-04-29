# gRPC mock server

Use [GripMock](https://github.com/tokopedia/gripmock) as an example to demonstrate the use of gRPC mock server.


## Mock server

Start a mock server according to proto files and a **stub** file in json format:

```bash
docker run  \
    -p 4770:4770  -p 4771:4771  \
    -v $(pwd)/src/addr:/proto   \
    tkpd/gripmock:latest  \
      --imports="/protobuf,/proto" \
      --stub="/proto"  \
    /proto/addr.proto
```

- GRPC server serves on tcp://localhost:4770.

- Stub server serves on http://localhost:4771.


## Test client

Inspect the stub info of the mock server:

```bash
curl localhost:4771
```

Make a gRPC request to the mock server:

```bash
grpcurl -plaintext    \
    -import-path "src/addr"  \
    -proto addr.proto   \
    localhost:4770  \
    addr.Addr.GetAddr
```
