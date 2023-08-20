# Facebook chat bot


#### Compile
`go version` 1.21rc
```golang

$ go build
```
or 
```shell

$ docker build -t fbchatbot .
$ docker run -it --name bot -p 8008:8008 -d fbchatbot
```
