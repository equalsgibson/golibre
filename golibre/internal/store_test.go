package internal_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/equalsgibson/golibre/golibre"
	"github.com/equalsgibson/golibre/golibre/internal"
)

func BenchmarkStoreOperations(b *testing.B) {
	// Set up
	keys := []golibre.UserID{}
	items := map[golibre.UserID]internal.Item[[]golibre.GraphGlucoseMeasurement]{}

	for i := range 100 {
		k := golibre.UserID(fmt.Sprint(i))
		keys = append(keys, k)
		items[k] = internal.Item[[]golibre.GraphGlucoseMeasurement]{}
	}

	store := internal.NewStore[golibre.UserID, []golibre.GraphGlucoseMeasurement]()
	store.Set(items)

	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			store.GetMultiple(context.Background(), keys)
		}
	})
}
