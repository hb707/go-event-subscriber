go mod init
go mod tidy

## 사용하는 패키지

```
require (
  github.com/ethereum/go-ethereum v1.13.5
  github.com/naoina/toml v0.1.2-0.20170918210437-9fafd6967416
  go.mongodb.org/mongo-driver v1.13.0
)
```

go get github.com/naoina/toml
go get go.mongodb.org/mongo-driver/mongo
go get github.com/ethereum/go-ethereum

\*\* ethclient import 시에 에러가 발생
go get github.com/ethereum/go-ethereum/rpc@v1.13.10
위 패키지 추가 인스톨 해줌!
