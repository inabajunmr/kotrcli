# kotrcli
kotr is api client for KING OF TIME My Recorder.

# install
```
go get github.com/inabajunmr/kotrcli
```

# Usage
```go
package main

import (
	"fmt"
	"github.com/inabajunmr/kotrcli"
)

func main() {
  token := "YOURTOKEN"
  userToken := "YOURUSERTOKEN"
  result := kotrcli.Dakoku(kotrcli.SYUKKIN, token, userToken)
  fmt.Println(result)
}


```
