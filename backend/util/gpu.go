package util

import (
	"log"
	"strings"
)

type GpuBrand string

const (
	APPLE   GpuBrand = "Apple Silicon"
	NVIDIA  GpuBrand = "NVIDIA"
	AMD     GpuBrand = "AMD"
	INTEL   GpuBrand = "Intel"
	UNKNOWN GpuBrand = "Unknown"
)

func GetGpu() GpuBrand {
	cmd := gpuCmd()
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	return parseGPUBrand(string(out))
}

func parseGPUBrand(output string) GpuBrand {
	output = strings.ToLower(output)

	if strings.Contains(output, "apple") {
		return APPLE
	} else if strings.Contains(output, "nvidia") {
		return NVIDIA
	} else if strings.Contains(output, "amd") {
		return AMD
	} else if strings.Contains(output, "intel") {
		return INTEL
	} else {
		return UNKNOWN
	}
}
