## Installation

```
$ sudo su
$ echo 'deb [trusted=yes] https://apt.fury.io/kaginari/ /' >> /etc/apt/sources.list.d/fury.list
$ apt update
$ apt install pfcli
```

##### Init go module :
```
 $ go mod init gitlab.com/XXXXX
```
##### Run go project : 

```
$ go run main.go
```

##### Install go binary :

```
$ go install
```

##### lunch binary :

```
$ pfcli <COMMAND> <ARG>
```


## Lunch runner :

 you need to lunch runner he will automatically register
```
$ cd tmp_runner
$ terragrunt apply 
```
