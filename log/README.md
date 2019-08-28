# log
log module 

> Just a wrapper now, should be refactored later.

## Quick Start

``` golang
func init() {
	log.InitWithConfigFile("log/test.toml")
	log.Info("Hello")
	log.ErrorF("err: %v", "...")
}
```