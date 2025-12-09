package net

import (
	"net"
	"testing"
)

// --- 测试 IsPrivate 函数 ---

// TestIsPrivateIPv4 验证 IPv4 私有地址和非私有地址的判断。
func TestIsPrivateIPv4(t *testing.T) {
	tests := []struct {
		ip       string
		expected bool
	}{
		// RFC 1918 私有地址
		{"10.0.0.1", true},       // A 类私有
		{"10.255.255.255", true}, // A 类私有
		{"172.16.0.1", true},     // B 类私有起始
		{"172.31.255.255", true}, // B 类私有结束
		{"192.168.1.1", true},    // C 类私有

		// 公共地址
		{"8.8.8.8", false},   // Google DNS
		{"1.1.1.1", false},   // Cloudflare DNS
		{"180.1.1.1", false}, // 随机公共 IP

		// 特殊地址 (不应被 IsPrivate 标记)
		{"127.0.0.1", false},   // 回环
		{"169.254.1.1", false}, // 链路本地

		// 运营商级 NAT (CGNAT) 地址 - 通常不视为 RFC 1918 私有
		{"100.64.0.1", false},

		// 无效地址
		{"", false},
	}

	for _, test := range tests {
		ip := net.ParseIP(test.ip)
		if ip == nil && test.ip != "" {
			// 如果是有效IP，但解析失败，则跳过
			continue
		}

		// 调用 IsPrivate 函数
		result := IsPrivate(ip)

		if result != test.expected {
			t.Errorf("IsPrivate(%s): 期望 %v, 得到 %v", test.ip, test.expected, result)
		}
	}
}

// TestIsPrivateIPv6 验证 IPv6 ULA 地址和非 ULA 地址的判断。
func TestIsPrivateIPv6(t *testing.T) {
	tests := []struct {
		ip       string
		expected bool
	}{
		// ULA 私有地址 (fd00::/8)
		{"fd00::1", true},
		{"fd00:1234:5678:90ab::", true},

		// ULA 保留地址 (fc00::/7，但 Go 中通常只检查 fd00::/8)
		{"fc00::1", false}, // 理论上是 ULA 的一部分，但在 fd00::/8 之外，通常不标记为 Private

		// 公共 IPv6 (全球单播地址)
		{"2001:db8::1", false},          // 文档示例地址
		{"2001:4860:4860::8888", false}, // Google DNS

		// 特殊地址 (不应被 IsPrivate 标记)
		{"::1", false},     // 回环
		{"fe80::1", false}, // 链路本地
	}

	for _, test := range tests {
		ip := net.ParseIP(test.ip)
		if ip == nil {
			continue
		}

		result := IsPrivate(ip)

		if result != test.expected {
			t.Errorf("IsPrivate(%s): 期望 %v, 得到 %v", test.ip, test.expected, result)
		}
	}
}

// --- 测试 IsPublic 函数 ---

// TestIsPublic 验证公共、私有和特殊地址的分类。
func TestIsPublic(t *testing.T) {
	tests := []struct {
		ip       string
		expected bool // 期望 IsPublic 返回的值
	}{
		// 公共地址 (期望 Public = true)
		{"8.8.8.8", true},
		{"203.111.55.1", true},
		{"2001:db8::10", true}, // 假设为公共，排除保留

		// 私有地址 (期望 Public = false)
		{"192.168.1.1", false},
		{"172.16.1.1", false},
		{"10.0.0.1", false},
		{"fd00::1", false}, // IPv6 ULA

		// 特殊地址 (期望 Public = false)
		{"127.0.0.1", false},   // 回环
		{"::1", false},         // 回环
		{"169.254.1.1", false}, // 链路本地
		{"fe80::1", false},     // 链路本地
		{"0.0.0.0", false},     // 未指定
		{"::", false},          // 未指定
		{"224.0.0.1", false},   // 组播
		{"ff00::1", false},     // 组播

		// IPv4 测试/保留地址 (期望 Public = false)
		{"192.0.2.1", false},   // TEST-NET-1
		{"203.0.113.1", false}, // TEST-NET-3
		{"240.0.0.1", false},   // 实验性
	}

	for _, test := range tests {
		ip := net.ParseIP(test.ip)
		if ip == nil {
			continue
		}

		result := IsPublic(ip)

		if result != test.expected {
			t.Errorf("IsPublic(%s): 期望 %v, 得到 %v", test.ip, test.expected, result)
		}
	}
}

// --- 测试辅助方法 (可选) ---

// TestInitPrivateRanges 验证 init 函数是否正确初始化了私有地址范围。
func TestInitPrivateRanges(t *testing.T) {
	if len(privateIPv4Ranges) != 3 {
		t.Errorf("init() 期望初始化 3 个 IPv4 私有地址范围，得到 %d 个", len(privateIPv4Ranges))
	}

	// 简单检查一个范围是否正确 (例如 10.0.0.0/8)
	tenNet := net.ParseIP("10.0.0.1")
	if !privateIPv4Ranges[0].Contains(tenNet) {
		t.Errorf("初始化失败：10.0.0.0/8 范围未被正确包含")
	}
}
