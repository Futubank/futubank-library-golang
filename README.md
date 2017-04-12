# futubank-library-golang

Модуль для вычисления криптографической подписи, используемой при отправке запросов к процессинговому центру futubank.com

## Установка
    go get github.com/futubank/futubank-library-golang

## Пример использования
```golang
package main

import "github.com/futubank/futubank-library-golang"
import "fmt"

func main() {
        data := map[string]string{
                "amount":      "100.10",
                "description": "Order Description",
        }
        signature := futubank.CalculateSignature(data, "secret_key")

        fmt.Printf("Signature: %s\n", signature)
}
```

