# kotrcli

[![GoDoc](https://godoc.org/github.com/inabajunmr/kotrcli/github?status.svg)](https://godoc.org/github.com/inabajunmr/kotrcli)

kotrcli is api client for KING OF TIME My Recorder.

# install
```
go get github.com/inabajunmr/kotrcli
```

# Usage
```go
ackage main
  
import (
        "fmt"
        "github.com/inabajunmr/kotrcli"
)

func main() {
  token := "YOURTOKEN"
  userToken := "YOURUSERTOKEN"
  result, _ := kotrcli.Dakoku(kotrcli.SYUKKIN, token, userToken)
  fmt.Println(result)
}
```
