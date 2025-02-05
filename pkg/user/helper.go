package user

import "sync/atomic"

var counter int32 = 0

func GenerateID() int {

	return int(atomic.AddInt32(&counter, 1))

}
