package test

import (
	"algo/algorithm/util"
	"bufio"
	"os"
	"testing"
)

func TestScanInputs(t *testing.T){
	in := bufio.NewReader(os.Stdin)
	for i := 0; ;i++{
		line,err := in.ReadSlice('\n')
		if line[0] == '\n' ||  err != nil{
			break
		}
		util.DD(line)
	}
}
