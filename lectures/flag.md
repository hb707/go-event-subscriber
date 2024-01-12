# flag 모듈 사용

main.go 파일에

```go
var configFlag = flag.String("config", "./config.toml", "toml env file not found") // 실행시 flag값을 추가해서 값을 받아올 수 있게 해줌
```

라는 flag를 생성해주고,

```go
func main() {
	flag.Parse()
	config.NewConfig(*configFlag)
	a := app.NewApp()
	fmt.Println(a)
}
```

형태로 flag.Parse()를 해줌.

이렇게 하면,

```
go run main.go
```

라는 cli로 실행하게 될 때, 플래그를 입력하여 유동적으로 받아올 수 있도록 해줌.
flag에 들어간 value가 default로 들어가지만 -config ~~~ 형태로도 지정이 가능

```
go run app/main.go -config ./env/config.toml
```

이런형태로 직접 cli에 입력하여 경로를 변경할 수 있음.
