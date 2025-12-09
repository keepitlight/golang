package net

import "net"

// 私有 IPv4 地址范围 (RFC 1918)
var privateIPv4Ranges []*net.IPNet

func init() {
	for _, cidr := range []string{
		"10.0.0.0/8",     // A 类私有地址
		"172.16.0.0/12",  // B 类私有地址
		"192.168.0.0/16", // C 类私有地址
	} {
		_, block, err := net.ParseCIDR(cidr)
		if err == nil {
			privateIPv4Ranges = append(privateIPv4Ranges, block)
		}
	}
}

// IsPrivate checks if an IP address is a private address (RFC 1918 or RFC 4193 ULA).
//
// 检测 IP 地址是否为私有地址 (RFC 1918 或 RFC 4193 ULA)
func IsPrivate(ip net.IP) bool {
	// 检查 IPv4 私有地址 (RFC 1918)
	if ip.To4() != nil {
		for _, block := range privateIPv4Ranges {
			if block.Contains(ip) {
				return true
			}
		}
		// IPv4 共享地址空间 (RFC 6598) 100.64.0.0/10
		// 有时也会被视为准私有地址，但通常 RFC 1918 是最主要的判断标准。
		// 如果需要，可以添加 "100.64.0.0/10" 的判断
		return false
	}

	// 检查 IPv6 私有地址 (唯一本地地址 ULA, RFC 4193)
	// ULA 地址范围是 fc00::/7，Go 提供了 IsPrivate 方法来检查这个范围。
	// 但是，请注意：Go 的 IsPrivate 方法实际上是检查 IPv6 ULA，
	// 而不是 RFC 1918 的 IPv4 私有地址。
	// 因此，我们应该直接检查 ULA 范围
	// ULA 的前缀是 1111110 (fc00::/7)
	// 实际的分配范围是 fd00::/8，因为 fc00::/8 是保留的。
	if ip.To16() != nil {
		// 检查 IPv6 唯一本地地址 (ULA) fd00::/8
		// net.ParseIP("fd00::").Mask(net.CIDRMask(8, 128)).Equal(ip.Mask(net.CIDRMask(8, 128)))
		// 也可以手动比较字节，但使用 Contains 更可靠。
		_, ulaBlock, err := net.ParseCIDR("fd00::/8")
		if err == nil && ulaBlock.Contains(ip) {
			return true
		}
	}

	return false
}

// IsPublic checks if an IP address is a globally routable (public) address.
//
// 检测 IP 地址是否为全局路由（公共）地址。
func IsPublic(ip net.IP) bool {
	if ip == nil {
		return false
	}

	// 1. 检查是否为私有地址 (RFC 1918 / ULA)
	if IsPrivate(ip) {
		return false
	}

	// 2. 检查是否为特殊用途地址
	if ip.IsLoopback() {
		return false // 回环地址 (127.0.0.0/8, ::1)
	}
	if ip.IsMulticast() {
		return false // 组播地址 (224.0.0.0/4, ff00::/8)
	}
	if ip.IsUnspecified() {
		return false // 未指定地址 (0.0.0.0, ::)
	}
	if ip.IsLinkLocalUnicast() {
		return false // 链路本地地址 (169.254.0.0/16, fe80::/10)
	}

	// 3. 检查是否为 IPv4 保留或测试范围
	if ip.To4() != nil {
		// 检查 TEST-NET-1, 2, 3 (192.0.2.0/24, 198.51.100.0/24, 203.0.113.0/24)
		_, test1, _ := net.ParseCIDR("192.0.2.0/24")
		_, test2, _ := net.ParseCIDR("198.51.100.0/24")
		_, test3, _ := net.ParseCIDR("203.0.113.0/24")

		if test1.Contains(ip) || test2.Contains(ip) || test3.Contains(ip) {
			return false
		}

		// 检查网络地址转换 (NAT) 的保留地址 192.0.0.0/24 (RFC 6890, IANA Reserved)
		_, natReserved, _ := net.ParseCIDR("192.0.0.0/24")
		if natReserved.Contains(ip) {
			return false
		}

		// 检查 IPv4 实验性地址 (240.0.0.0/4)
		_, expBlock, _ := net.ParseCIDR("240.0.0.0/4")
		if expBlock.Contains(ip) {
			return false
		}
	}

	// 4. 排除所有已知特殊地址后，剩下的通常就是公共的、可路由的地址
	return true
}
