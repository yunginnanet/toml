package toml

import (
	"bufio"
	"bytes"
	"strings"
)

func removeQuotes(data []byte) string {
	xerox := bufio.NewScanner(bytes.NewReader(data))
	str := &strings.Builder{}
	for xerox.Scan() {
		if !strings.Contains(xerox.Text(), `"`) {
			str.Write(xerox.Bytes())
			str.Write([]byte{'\n'})
			continue
		}
		split := strings.Split(xerox.Text(), "=")
		str.WriteString(split[0])
		str.WriteString("=")
		str.WriteString(strings.ReplaceAll(split[1], `"`, ""))
		str.Write([]byte{'\n'})
	}
	return str.String()
}
