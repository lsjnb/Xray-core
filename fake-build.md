# fake-main


> 这个版本只增加了授权模块

```bash
go build -o build_assets/core.exe -trimpath -buildvcs=false -gcflags="all=-l=4" -ldflags="-X github.com/xtls/xray-core/core.build=$(git describe --always --dirty) -s -w -buildid=" -v ./main


GOOS=linux GOARCH=amd64 go build -o build_assets/core -trimpath -buildvcs=false -gcflags="all=-l=4" -ldflags="-X github.com/xtls/xray-core/core.build=$(git describe --always --dirty) -s -w -buildid=" -v ./main
```