package logutil

import (
	"net"
	"os"
	"strconv"
	"strings"
)

func DefaultGateway() string {
	// 读取 /proc/net/route 文件内容
	data, err := os.ReadFile("/proc/net/route")
	if err != nil {
		panic(err)
	}

	// 将文件内容按行分割
	lines := strings.Split(string(data), "\n")

	// 遍历每一行
	for _, line := range lines {
		// 分割行，以空格为分隔符
		fields := strings.Fields(line)
		if len(fields) < 3 {
			continue
		}
		// 提取默认网关的信息（Destination 为 00000000 表示默认网关）
		if fields[1] == "00000000" {
			// 将十六进制的 IP 地址转换为 IPv4 地址字符串
			ipHex := fields[2]
			ip := make(net.IP, 4)
			for i := 0; i < 4; i++ {
				val, _ := strconv.ParseInt(ipHex[2*i:2*(i+1)], 16, 0)
				ip[i] = byte(val)
			}
			// 将网关地址进行反转
			ip = reverseIP(ip)
			// 打印默认网关的 IP 地址

			return ip.String()
		}
	}

	panic("Default gateway not found")
}

// reverseIP 反转 IPv4 地址
func reverseIP(ip net.IP) net.IP {
	for i := 0; i < len(ip)/2; i++ {
		j := len(ip) - i - 1
		ip[i], ip[j] = ip[j], ip[i]
	}
	return ip
}
