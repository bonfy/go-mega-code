# Go-Mega Code

Source Code for [Go-Mega Tutorial](https://github.com/bonfy/go-mega)

## 如何使用

> 由于 Go 包的特殊引用方式， 最好将 code 都放到 $GOPATH/src/github.com/bonfy/go-mega-code 文件夹中再运行


#### 准备

请先安装 Go 环境，可以参考 [Prepare](https://go-mega.bonfy.im/00-prepare)

```cmd
# 安装完成后 查看版本
$ go version
go version go1.11 darwin/amd64
```

#### Git 方式

```cmd
# 建立文件夹

$ cd $GOPATH/src
$ mkdir -p github.com/bonfy
$ cd github.com/bonfy

# clone

$ git clone git@github.com:bonfy/go-mega-code.git
$ cd go-mega-code

# pull 指定分支

$ git pull origin 01-Hello-World
$ git checkout 01-Hello-World
$ go run main.go
```

#### Go Get 方式

```cmd
$ go get -v github.com/bonfy/go-mega-code

# 会提示 No go files 不过不要紧，继续就可以了

# cd 到 刚才get到的文件夹

$ cd $GOPATH/src/github.com/bonfy/go-mega-code

# pull 指定分支

$ git pull origin 01-Hello-World
$ git checkout 01-Hello-World
$ go run main.go

```

#### Download Zip方式

下载每章的 zip，并解压，进入文件夹 运行 

```cmd
$ go run main.go
```


## Branch列表

| Branch | Source Code | Zip |
| :--- | :--- | :--- |
| 01-Hello-World | [View](https://github.com/bonfy/go-mega-code/tree/01-Hello-World) | [Download](https://github.com/bonfy/go-mega-code/archive/v0.1.zip)|
| 02-Template-Basic | [View](https://github.com/bonfy/go-mega-code/tree/02-Template)  | [Download](https://github.com/bonfy/go-mega-code/archive/v0.2.zip)|
| 03-Template-Advance | [View](https://github.com/bonfy/go-mega-code/tree/03-Template-Advance) | [Download](https://github.com/bonfy/go-mega-code/archive/v0.3.zip) |

## 支持我

![support](https://github.com/bonfy/go-mega/blob/master/images/sponsor.jpg?raw=true)


