# 啪唧工具包 paji-sdk

Golang 的開發工具包

## 使用

```go
package main

import "github.com/marco79423/paji-sdk.go/converter"

func main() {
	var i = 0
	var j *int = converter.ConvertIntToIntRef(i)
}
```

## 功能模組

| 模組        | 功能           |
|-------------|----------------|
| converter    | 型態轉換 |
