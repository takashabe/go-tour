package main

import "fmt"

type IPAddr [4]byte

// expected: IPAddr{1,  2,  3,  4} => "1.2.3.4"
func (addr IPAddr) String() string {
	var ret string
	for _, v := range addr {
		ret = ret + fmt.Sprint(v) + "."
	}
	ret = ret[:len(ret)-1]
	ret = "\"" + ret + "\""
	return ret
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
