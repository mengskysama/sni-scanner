package scanner

import (
	"errors"
	"net"
	"strconv"
	"strings"
)

func inetA2n(ip string) (int64, error) {
	bits := strings.Split(ip, ".")
	if len(bits) != 4 {
		return 0, errors.New("invalid ip format")
	}
	b0, _ := strconv.Atoi(bits[0])
	b1, _ := strconv.Atoi(bits[1])
	b2, _ := strconv.Atoi(bits[2])
	b3, _ := strconv.Atoi(bits[3])

	var sum int64

	sum += int64(b0) << 24
	sum += int64(b1) << 16
	sum += int64(b2) << 8
	sum += int64(b3)

	return sum, nil
}

func inetN2a(ipInt int64) net.IP {
	var bytes [4]byte
	bytes[0] = byte(ipInt & 0xFF)
	bytes[1] = byte((ipInt >> 8) & 0xFF)
	bytes[2] = byte((ipInt >> 16) & 0xFF)
	bytes[3] = byte((ipInt >> 24) & 0xFF)
	return net.IPv4(bytes[3], bytes[2], bytes[1], bytes[0])
}
