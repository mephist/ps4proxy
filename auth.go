package main

import (
	"log"
	"net"
	"net/http"
	"strings"
	"sync"
)

type AuthCache struct {
	Allowed map[string]string
	sync.Mutex
}

func (c AuthCache) CheckIP(ip string) bool {
	c.Lock()
	defer c.Unlock()
	if user, ok := c.Allowed[ip]; ok {
		log.Printf("Found IP in cache: %s [%s]\n", ip, user)
		return true
	}
	return false
}

func (c *AuthCache) AddIP(ip, reason string) {
	c.Lock()
	defer c.Unlock()
	c.Allowed[ip] = reason
	log.Printf("Whitelisted IP: %s [%s]", ip, reason)
}

func (c *AuthCache) Init() {
	c.Allowed = make(map[string]string, 666)
}

var cache AuthCache

func init() {
	cache.Init()
}

func RemoteIp(r *http.Request) string {

	remoteAddr, _, _ := net.SplitHostPort(r.RemoteAddr)
	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}
	return remoteAddr
}

func CheckRequest(r *http.Request) bool {

	ip := RemoteIp(r)

	// Check the list of allowed IPs
	if cache.CheckIP(ip) {
		return true
	} else {
		if strings.HasSuffix(r.URL.Host, "playstation.net") {
			cache.AddIP(ip, "psn")
			return true
		}
	}
	log.Printf("Rejected request: %s\n", ip)
	return false
}
