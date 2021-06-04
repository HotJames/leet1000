package defang_ip

import "strings"

func DeFangIPAddr(address string) string {

	return strings.Replace(address, ".", "[.]", -1)

}
