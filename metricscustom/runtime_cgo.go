// +build cgo
// +build !appengine

package metricscustom

import "runtime"

func numCgoCall() int64 {
	return runtime.NumCgoCall()
}
