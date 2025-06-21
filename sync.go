package golang

type RWLocker interface {
	RLock()
	RUnlock()
}
