# kotrcli
kotr is api client for kotr.

```go
package main

import (
	"fmt"
	"github.com/kotrcli/kotrcli"
)

func main() {
  token := "YOURTOKEN"
  userToke := "YOURUSERTOKEN"
  result := kotrcli.Dakoku(kotrcli.SYUKKIN, token, userToken)
  fmt.Println(result)
}
```
