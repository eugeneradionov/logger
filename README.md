# Logger

### Description
Logger is the wrapper for uber/zap logger with custom configurations and presets

### Usage
```go
import "github.com/github.com/eugeneradionov/logger"

func LoadDevelopment() (*zap.Logger, error) {
    return logger.Load(logger.Development)
}

func LoadProduction() (*zap.Logger, error) {
    return logger.Load(logger.Production)
}
```
