package bank

import "sync"

// ILoginMeta is a read-only map that allows clients to
// discover login criteria and dynamically create login prompts
type ILoginMeta interface {
	Get(key string) string
	Entries() chan *LoginMetaEntry
}

func NewLoginMeta(setup map[string]string) ILoginMeta {
	d := map[string]string{}

	for k, v := range setup {
		d[k] = v
	}
	return &LoginMeta{
		m: &sync.Mutex{},
		d: d,
	}
}

type LoginMetaEntry struct {
	Key   string
	Value string
}

type LoginMeta struct {
	m *sync.Mutex
	d map[string]string
}

func (l *LoginMeta) Get(key string) string {
	l.m.Lock()
	found := l.d[key]
	l.m.Unlock()
	return found
}

func (l *LoginMeta) Entries() chan *LoginMetaEntry {
	c := make(chan *LoginMetaEntry)
	go func() {
		l.m.Lock()
		for k, v := range l.d {
			c <- &LoginMetaEntry{k, v}
		}
		l.m.Unlock()
	}()
	return c
}
