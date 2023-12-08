package main

import "fmt"

type IPAddr [4]byte

// String method to implement fmt.Stringer interface for IPAddr.
func (ip IPAddr) String() string {
	// Use fmt.Sprintf to format the IP address as a dotted quad.
	// This will convert each byte in the IPAddr to a string and join them with dots.
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
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
