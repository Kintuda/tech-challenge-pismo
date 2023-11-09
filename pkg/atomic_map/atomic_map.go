package atomicmap

// import "sync"

// type AtomicMap struct {
// 	Lock sync.Mutex
// 	Map  map[string]string
// }

// func NewAtomicMap() *AtomicMap {
// 	return &AtomicMap{}
// }

// func (a *AtomicMap) Insert(id, data string) error {
// 	a.Lock.Lock()
// 	a.Map[id] = data

// 	defer a.Lock.Unlock()
// 	retur nil
// }
