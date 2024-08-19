package js_replace

import (
	"bytes"
	"log"
)

func findToReplaceBlock(js []byte, findAfter string) (bool, []byte) {
	index := bytes.Index(js, []byte(findAfter))
	if index == -1 {
		log.Println("find js replace block not found:", findAfter)
		return false, nil
	}

	start := false
	openBraces := 0
	var result bytes.Buffer

	for _, char := range js[index+len(findAfter):] {
		if char == '{' {
			start = true
			openBraces++
		} else if char == '}' {
			openBraces--
		}
		if !start {
			continue
		}
		result.WriteByte(char)

		if openBraces == 0 {
			break
		}
	}
	log.Println("find js replace block", findAfter, result.String())
	return true, result.Bytes()
}

// TODO
func replaceJs(js string) string {
	return js
}
