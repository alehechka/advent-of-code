package fifteen

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func Day4Problem1(inputs []string) string {
	base := inputs[0]

	var answer int
	for {
		answer++
		if strings.HasPrefix(GetMD5Hash(fmt.Sprintf("%s%d", base, answer)), "00000") {
			return strconv.Itoa(answer)
		}
	}
}

func Day4Problem2(inputs []string) string {
	base := inputs[0]

	var answer int
	for {
		answer++
		if strings.HasPrefix(GetMD5Hash(fmt.Sprintf("%s%d", base, answer)), "000000") {
			return strconv.Itoa(answer)
		}
	}
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
