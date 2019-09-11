// +build go1.5

package metricscustom

import "runtime"

func gcCPUFraction(memStats *runtime.MemStats) float64 {
	return memStats.GCCPUFraction
}
