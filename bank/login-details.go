package bank

import "sync"

// ILoginDetails is a read-only map that allows clients to submit flexible login details.
// This is required as banks have a wide range of login criteria that cannot be encoded into a strict type.
//
// For successful operation, this needs to be validated at runtime by the IBank.Login() method
type ILoginDetails interface {
	Get(key string) string
	Entries() chan *LoginDetailsEntry
}

func NewLoginDetails(setup map[string]string) ILoginDetails {
	d := map[string]string{}

	for k, v := range setup {
		d[k] = v
	}

	return &LoginDetails{
		m: &sync.Mutex{},
		d: d,
	}
}

type LoginDetailsEntry struct {
	Key   string
	Value string
}

type LoginDetails struct {
	m *sync.Mutex
	d map[string]string
}

func (l *LoginDetails) Get(key string) string {
	l.m.Lock()
	found := l.d[key]
	l.m.Unlock()
	return found
}

func (l *LoginDetails) Entries() chan *LoginDetailsEntry {
	c := make(chan *LoginDetailsEntry)
	go func() {
		l.m.Lock()
		for k, v := range l.d {
			c <- &LoginDetailsEntry{k, v}
		}
		l.m.Unlock()
	}()
	return c
}
