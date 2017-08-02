libgosu
-------

Same functional logic as https://github.com/tianon/gosu but restructured to make it easier to utilize the functionality in other go projects

While the `gosu` cli utility can be built from this project, it is recommended to always use the cli utility from https://github.com/tianon/gosu

Example Lib Usage
```go
package myapp

import (
	"fmt"

	"github.com/emacski/libgosu"
)

err := libgosu.Exec("nobody", []string{"redis-server", "--maxmemory", "2mb", "6379"})
if err != nil {
	fmt.Errorf("error: %v", err)
}
```
