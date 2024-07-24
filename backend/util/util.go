package util

import "fmt"

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
)

func HumanReadableSize(size int64) string {
	var unit string
	var value float64

	switch {
	case size >= EB:
		unit = "EB"
		value = float64(size) / EB
	case size >= PB:
		unit = "PB"
		value = float64(size) / PB
	case size >= TB:
		unit = "TB"
		value = float64(size) / TB
	case size >= GB:
		unit = "GB"
		value = float64(size) / GB
	case size >= MB:
		unit = "MB"
		value = float64(size) / MB
	case size >= KB:
		unit = "KB"
		value = float64(size) / KB
	default:
		unit = "bytes"
		value = float64(size)
	}

	return fmt.Sprintf("%.2f %s", value, unit)
}
