package keyboard

import (
	"bufio"
	"strconv"
	"strings"
)

func GetFloat() (float64,float64) {
	reader:=bufio.NewReader(os.stdin)
	input,err:=reader.ReadString('\n')
	if err!=nil {
		return 0,err
	}

	input=strings.TrimSpace(input)
	number,err:=strconv.ParseFloat(input,64)
	if err!=nil {
		return 0,err
	}
	return number,nil
}
