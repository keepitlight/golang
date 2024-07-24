package hotfix_test

import (
	"context"
	"fmt"

	"github.com/keepitlight/golang/hotfix"
)

type HistoryWrapper struct {
	fixed    hotfix.Version
	recorder func(v hotfix.Version, summary []string)
}

func (h *HistoryWrapper) Fixed() hotfix.Version {
	return h.fixed
}

func (h *HistoryWrapper) Record(v hotfix.Version, summary []string) {
	if h.recorder != nil {
		h.recorder(v, summary)
	}
}

func ExamplePatch() {
	const (
		V3 hotfix.Version = 3
		V2 hotfix.Version = 2
		V1 hotfix.Version = 1
	)

	// STEP 1 : define fixes, usually defined in the init function
	hotfix.FixIt(V1, "Example 1", func(ctx context.Context, version hotfix.Version, title string) error {
		// do something
		return nil
	})
	hotfix.FixIt(V1, "Example 2", func(ctx context.Context, version hotfix.Version, title string) error {
		// do something
		return nil
	})
	hotfix.FixIt(V2, "Example 3", func(ctx context.Context, version hotfix.Version, title string) error {
		// do something
		return nil
	})
	hotfix.FixIt(V2, "Example 4", func(ctx context.Context, version hotfix.Version, title string) error {
		// do something
		return nil
	})
	hotfix.FixIt(V3, "Example 5", func(ctx context.Context, version hotfix.Version, title string) error {
		// do something
		return nil
	})
	hotfix.FixIt(V3, "Example 6", func(ctx context.Context, version hotfix.Version, title string) error {
		// do something
		return nil
	})

	// STEP 2 : apply fixes
	hotfix.SetHistoryProvider(&HistoryWrapper{
		fixed: V1,
		recorder: func(v hotfix.Version, summary []string) {
			fmt.Printf("[v%d] fixed\n", v)
			for i, s := range summary {
				fmt.Printf("%d. %s\n", i+1, s)
			}
		},
	})
	if err := hotfix.Patch(context.Background(), V1); err != nil {
		panic(err)
	}
	// Output:
	// [v2] fixed
	// 1. Example 3
	// 2. Example 4
	// [v3] fixed
	// 1. Example 5
	// 2. Example 6
}

func ExampleFixIt() {
	const (
		V3 hotfix.Version = 3
		V2 hotfix.Version = 2
		V1 hotfix.Version = 1
	)

	// STEP 1 : define fixes, usually defined in the init function
	hotfix.FixIt(V1, "Example 1", func(ctx context.Context, version hotfix.Version, title string) error {
		fmt.Printf("[v%d] %s fixed\n", version, title)
		return nil
	})
	hotfix.FixIt(V1, "Example 2", func(ctx context.Context, version hotfix.Version, title string) error {
		fmt.Printf("[v%d] %s fixed\n", version, title)
		return nil
	})
	hotfix.FixIt(V2, "Example 3", func(ctx context.Context, version hotfix.Version, title string) error {
		fmt.Printf("[v%d] %s fixed\n", version, title)
		return nil
	})
	hotfix.FixIt(V2, "Example 4", func(ctx context.Context, version hotfix.Version, title string) error {
		fmt.Printf("[v%d] %s fixed\n", version, title)
		return nil
	})
	hotfix.FixIt(V3, "Example 5", func(ctx context.Context, version hotfix.Version, title string) error {
		fmt.Printf("[v%d] %s fixed\n", version, title)
		return nil
	})
	hotfix.FixIt(V3, "Example 6", func(ctx context.Context, version hotfix.Version, title string) error {
		fmt.Printf("[v%d] %s fixed\n", version, title)
		return nil
	})

	// STEP 2 : apply fixes
	hotfix.SetHistoryProvider(&HistoryWrapper{
		fixed: V1,
	})
	if err := hotfix.Patch(context.Background(), V1); err != nil {
		panic(err)
	}
	// Output:
	// [v2] Example 3 fixed
	// [v2] Example 4 fixed
	// [v3] Example 5 fixed
	// [v3] Example 6 fixed
}
