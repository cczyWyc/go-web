package onclass

import (
	"fmt"
	"time"
)

type FilterBuilder func(next Filter) Filter

type handlerFunc func(c *Context)

type Filter func(c *Context)

var _ FilterBuilder = MetricsFilterBuilder

func MetricsFilterBuilder(next Filter) Filter {
	return func(c *Context) {
		start := time.Now().Nanosecond()
		next(c)
		end := time.Now().Nanosecond()
		fmt.Printf("use %d ns", end-start)
	}
}
