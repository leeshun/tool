package ip

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

/*
{
  "ip": "149.28.139.3",
  "hostname": "149.28.139.3.vultr.com",
  "city": "Singapore",
  "region": "Central Singapore Community Development Council",
  "country": "SG",
  "loc": "1.2929,103.8550",
  "org": "AS20473 Choopa, LLC"
}
*/
type Info struct {
	IP       string `json:"ip"`
	Hostname string `json:"hostname"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
}

func GetIPInfo() string {
	var info Info
	resp, err := http.Get("https://ipinfo.io")
	if err != nil {
		fmt.Println("can't get ip info, the reason is", err)
		return ""
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Println("close response body err:", err)
		}
	}()
	decoder := json.NewDecoder(resp.Body)
	if err = decoder.Decode(&info); err != nil {
		fmt.Println("decode response body err", err)
	}
	b, _ := json.MarshalIndent(info, "", "  ")
	return string(b)
}

func GetIP() string {
	var ipAddress string
	resp, err := http.Get("https://ipinfo.io/ip")
	if err != nil {
		fmt.Println("can't get ip address, the reason is", err)
		return ipAddress
	}
	buf := bufio.NewReader(resp.Body)
	ipAddress, err = buf.ReadString('\n')
	if err != nil {
		fmt.Println("get ip address from response body err: ", err)
	}
	strings.TrimSpace(ipAddress)
	ipAddress = strings.Fields(ipAddress)[0]
	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Println("close response body err:", err)
		}
	}()
	return ipAddress
}

func GetLocation() string {
	var location string
	resp, err := http.Get("https://ipinfo.io/city")
	if err != nil {
		fmt.Println("can't get ip location, the reason is", err)
		return location
	}
	buf := bufio.NewReader(resp.Body)
	location, err = buf.ReadString('\n')
	if err != nil {
		fmt.Println("get ip location from response body err: ", err)
	}
	strings.TrimSpace(location)
	location = strings.Fields(location)[0]
	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Println("close response body err:", err)
		}
	}()
	return location
}
