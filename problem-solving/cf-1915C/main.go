package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	t, _ := strconv.Atoi(strings.TrimSpace(line))

	for ; t > 0; t-- {
		line, _ = reader.ReadString('\n')
		n, _ := strconv.Atoi(strings.TrimSpace(line))
		line, _ = reader.ReadString('\n')
		nums := strings.Fields(line)

		var sum int64 = 0
		for i := 0; i < n; i++ {
			a, _ := strconv.ParseInt(nums[i], 10, 64)
			sum += a
		}

		sqrt := int64(math.Sqrt(float64(sum)))
		if sqrt*sqrt == sum {
			writer.WriteString("YES\n")
		} else {
			writer.WriteString("NO\n")
		}
	}
}
