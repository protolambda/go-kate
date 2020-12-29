package kate

import (
	"fmt"
	"strings"
)

func debugBigPtrs(msg string, values []*Big) {
	var out strings.Builder
	out.WriteString("---")
	out.WriteString(msg)
	out.WriteString("---\n")
	for i := range values {
		out.WriteString(fmt.Sprintf("#%4d: %s\n", i, bigStr(values[i])))
	}
	fmt.Println(out.String())
}

func debugG1s(msg string, values []G1) {
	var out strings.Builder
	out.WriteString("---")
	out.WriteString(msg)
	out.WriteString("---\n")
	for i := range values {
		out.WriteString(fmt.Sprintf("#%4d: %s\n", i, strG1(&values[i])))
	}
	fmt.Println(out.String())
}

func debugBigs(msg string, values []Big) {
	var out strings.Builder
	out.WriteString("---")
	out.WriteString(msg)
	out.WriteString("---\n")
	for i := range values {
		out.WriteString(fmt.Sprintf("#%4d: %s\n", i, bigStr(&values[i])))
	}
	fmt.Println(out.String())
}

func debugBigsOffsetStride(msg string, values []Big, offset uint64, stride uint64) {
	var out strings.Builder
	out.WriteString("---")
	out.WriteString(msg)
	out.WriteString("---\n")
	j := uint64(0)
	for i := offset; i < uint64(len(values)); i += stride {
		out.WriteString(fmt.Sprintf("#%4d: %s\n", j, bigStr(&values[i])))
		j++
	}
	fmt.Println(out.String())
}