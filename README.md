# go-safecast
Library for safe type conversion in Go

## What is this
The type of `int` equals `int64` on 64-bit machine in Go.  
When you convert `int`(`int64`) to `int32`, `int8` or `int6`, Your code could have Integer Overflow vulnerability.

In 2019, [Kubernetes](https://kubernetes.io/) had the vulnerability. and the vulnerability was found on [Security Audit Project](https://github.com/kubernetes/community/blob/master/sig-security/security-audit-2019/findings/Kubernetes%20Final%20Report.pdf) by Trail of Bits.

This library is created to prevent the vulnerability creation.

**(This library is heavily inspired by Kubernetes's Security Audit Report by Trail of Bits)**

## Usage

```go
import "github.com/rung/go-safecast"
```

### Convert int to int32
```go
	i := 2147483647
	i32, err := safecast.Int32(i) // convert int to int32 in a safe way
	if err != nil {
		return err
	}
```
This library also have `safecast.Int16` and `safecast.Int8`. You can use the functions in the same way as `safecast.Int32`

### Convert string to int32
```go
	s := "2147483647"
	i, err := safecast.Atoi32(s) // convert string to int32 in a safe way
	if err != nil {
		return i
	}
```
This library also have `safecast.Atoi16` and `safecast.Atoi8`. You can use the functions in the same way as `safecast.Atoi32`


## What happens when overflows
### Range of each integer
|       | int32 - signed 32bit integer         | int16 - 16bit signed integer | int8 - 8bit signed integer | 
| :---: | :----------------------------------: | :--------------------------: | :------------------------: | 
| Range | From -2,147,483,648 to 2,147,483,647 | From -32,768 to 32,767       | From -128 to 127           | 

### When using native int32(), the code cause overflows
![native int32](img/native-int32.png)  
Link: [Go Playground](https://play.golang.org/p/YsKayPgMX7k)

### When using this library, your code is safe
![native int32](img/safecast-int32.png)  
This library can detect integer overflow. so you can convert integer in a safe way.  
Link: [Go Playground](https://play.golang.org/p/VUkXLIijq6Q)

## License
[MIT License](LICENSE)
