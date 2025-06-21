package hotfix

import (
	"context"
	"sync"
)

type Version int // Hotfix version number

const (
	Omitted Version = 0 // omitted version
)

type Hotfix func(ctx context.Context, version Version, title string) error

type call struct {
	title  string
	hotfix Hotfix
}

type fixes struct {
	history History
	calls   map[Version][]*call
	mutex   sync.Mutex
	once    sync.Once
}

func (f *fixes) append(ver Version, title string, hotfixes ...Hotfix) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if f.calls == nil {
		f.calls = make(map[Version][]*call)
	}
	for _, hotfix := range hotfixes {
		var cs []*call
		if v, ok := f.calls[ver]; ok {
			cs = v
		}
		f.calls[ver] = append(cs, &call{
			title:  title,
			hotfix: hotfix,
		})
	}
}

func (f *fixes) do(ctx context.Context, version Version) (err error) {
	f.once.Do(func() {
		if version <= Omitted {
			return
		}
		last := Omitted
		if f.history != nil {
			// check history
			last = f.history.Fixed()
		}
		if len(f.calls) < 1 {
			return
		}
		var mv Version
		for v := range f.calls {
			if v > mv {
				mv = v
			}
		}
		if version > mv {
			return
		}
		for i := version; i <= mv; i++ {
			if i <= last {
				continue
			}
			var summary []string
			if cs, ok := f.calls[i]; ok {
				for _, c := range cs {
					summary = append(summary, c.title)
					if err = c.hotfix(ctx, i, c.title); err != nil {
						return
					}
				}
			}
			if f.history != nil {
				f.history.Record(i, summary)
			}
		}
	})
	return
}
