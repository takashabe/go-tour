package main

import "fmt"

type IPAddr [4]byte

// expected: IPAddr{1,  2,  3,  4} => "1.2.3.4"
func (i IPAddr) String() string {
	// TODO: 配列ループさせて値を取り出して結合
	return fmt.Sprintf("")
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
