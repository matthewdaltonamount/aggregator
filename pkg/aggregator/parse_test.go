package aggregator

import (
	"context"
	"fmt"
	"testing"
)

func TestParseLogs(t *testing.T) {

	t.Run("it parses logs", func(t *testing.T) {
		ctx := context.Background()
		log := "log.txt"
		ctx = context.WithValue(ctx, "log", log)

		responseAgg, err := ParseLogController(ctx)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(responseAgg)
	})
}
