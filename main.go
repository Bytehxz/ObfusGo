package main

import (
	"fmt"
	"ObfusGo/getProxy"
)

func main() {
	fmt.Println("Getting proxies...")

	proxies := getProxy.GetProxies()

	fmt.Println("Proxies list:")
	for _, proxy := range proxies {
		fmt.Println(proxy)
	}
}

