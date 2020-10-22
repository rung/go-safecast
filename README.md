# go-safecast
Library for safe type conversion in Go

## Usage

```go
import "rung/go-safecast"
```

### Convert int to int32
```go
	i := 2147483647
	converted, err := safecast.Int32(i) // convert int to int32 in a safe way
	if err != nil {
		return err
	}
```

### Convert string to int32
```go
	s := "2147483647"
	i, err := safecast.Atoi32(s) // convert string to int32 in a safe way
	if err != nil {
		return i
	}
```
