package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func fetchProxies(urls []string, client *http.Client) string {
	var proxies string
	for _, url := range urls {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println("\nError creating request for", url, ":", err)
			continue
		}
		req.Header.Add("User-Agent", "Mozilla/4.0 (PSP (PlayStation Portable); 2.00)")
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("\nError fetching", url, ":", err)
			continue
		}
		if resp.StatusCode != http.StatusOK {
			fmt.Println("\nError: Unexpected status code", resp.StatusCode, "for", url)
			resp.Body.Close()
			continue
		}
		body, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Println("\nError reading response from", url, ":", err)
			continue
		}
		proxies += string(body) + "\n"
	}
	return proxies
}

func main() {
	socks4Urls := []string{
		"https://raw.githubusercontent.com/jetkai/proxy-list/main/online-proxies/txt/proxies-socks4.txt",
		"https://api.proxyscrape.com/v2/?request=getproxies&protocol=socks4&timeout=10000&country=all",
		"https://raw.githubusercontent.com/jetkai/proxy-list/main/archive/txt/proxies-socks4.txt",
		"https://raw.githubusercontent.com/roosterkid/openproxylist/main/SOCKS4_RAW.txt",
		"https://raw.githubusercontent.com/monosans/proxy-list/main/proxies/socks4.txt",
		"https://raw.githubusercontent.com/TheSpeedX/PROXY-List/master/socks4.txt",
		"https://www.proxy-list.download/api/v1/get?type=socks4",
		"https://www.proxyscan.io/download?type=socks4",
		"https://api.openproxylist.xyz/socks4.txt",
		"http://proxysearcher.sourceforge.net/Proxy%20List.php?type=socks",
		"http://worm.rip/socks4.txt",
		"http://www.socks24.org/feeds/posts/default",
		"https://api.proxyscrape.com/?request=displayproxies&proxytype=socks4",
		"https://api.proxyscrape.com/?request=displayproxies&proxytype=socks4&country=all",
		"https://api.proxyscrape.com/v2/?request=displayproxies&protocol=socks4",
		"https://api.proxyscrape.com/v2/?request=getproxies&protocol=socks4",
		"https://dstat.one/proxies.php?id=4",
		"https://proxyspace.pro/socks4.txt",
		"https://raw.githubusercontent.com/B4RC0DE-TM/proxy-list/main/SOCKS4.txt",
		"https://raw.githubusercontent.com/HyperBeats/proxy-list/main/socks4.txt",
		"https://raw.githubusercontent.com/mmpx12/proxy-list/master/socks4.txt",
		"https://raw.githubusercontent.com/monosans/proxy-list/main/proxies_anonymous/socks4.txt",
		"https://raw.githubusercontent.com/saschazesiger/Free-Proxies/master/proxies/socks4.txt",
		"https://raw.githubusercontent.com/ShiftyTR/Proxy-List/master/socks4.txt",
		"https://raw.githubusercontent.com/TheSpeedX/SOCKS-List/master/socks4.txt",
		"https://www.freeproxychecker.com/result/socks4_proxies.txt",
		"https://www.my-proxy.com/free-socks-4-proxy.html",
	}
	socks5Urls := []string{
		"https://raw.githubusercontent.com/jetkai/proxy-list/main/online-proxies/txt/proxies-socks5.txt",
		"https://api.proxyscrape.com/v2/?request=getproxies&protocol=socks5&timeout=10000&country=all",
		"https://raw.githubusercontent.com/jetkai/proxy-list/main/archive/txt/proxies-socks5.txt",
		"https://raw.githubusercontent.com/roosterkid/openproxylist/main/SOCKS5_RAW.txt",
		"https://github.com/jetkai/proxy-list/blob/main/archive/txt/proxies-socks5.txt",
		"https://raw.githubusercontent.com/monosans/proxy-list/main/proxies/socks5.txt",
		"https://raw.githubusercontent.com/hookzof/socks5_list/master/proxy.txt",
		"https://raw.githubusercontent.com/TheSpeedX/PROXY-List/master/socks5.txt",
		"https://www.proxy-list.download/api/v1/get?type=socks5",
		"https://api.openproxylist.xyz/socks5.txt",
		"http://worm.rip/socks5.txt",
		"http://www.live-socks.net/feeds/posts/default",
		"http://www.socks24.org/feeds/posts/default",
		"https://api.proxyscrape.com/?request=displayproxies&proxytype=socks5",
		"https://api.proxyscrape.com/v2/?request=displayproxies&protocol=socks5",
		"https://api.proxyscrape.com/v2/?request=getproxies&protocol=socks5&timeout=10000&country=all&simplified=true",
		"https://dstat.one/proxies.php?id=5",
		"https://proxyspace.pro/socks5.txt",
		"https://raw.githubusercontent.com/B4RC0DE-TM/proxy-list/main/SOCKS5.txt",
		"https://raw.githubusercontent.com/HyperBeats/proxy-list/main/socks5.txt",
		"https://raw.githubusercontent.com/manuGMG/proxy-365/main/SOCKS5.txt",
		"https://raw.githubusercontent.com/mmpx12/proxy-list/master/socks5.txt",
		"https://raw.githubusercontent.com/monosans/proxy-list/main/proxies_anonymous/socks5.txt",
		"https://raw.githubusercontent.com/saschazesiger/Free-Proxies/master/proxies/socks5.txt",
		"https://raw.githubusercontent.com/ShiftyTR/Proxy-List/master/socks5.txt",
		"https://www.freeproxychecker.com/result/socks5_proxies.txt",
		"https://www.my-proxy.com/free-socks-5-proxy.html",
		"https://www.proxyscan.io/download?type=socks5",
	}

	client := &http.Client{}
	proxies := fetchProxies(socks4Urls, client) + fetchProxies(socks5Urls, client)

	outFile, err := os.Create("proxies/socks.txt")
	if err != nil {
		fmt.Println("\nError while creating 'socks.txt':", err)
		return
	}
	defer outFile.Close()

	_, err = outFile.WriteString(proxies)
	if err != nil {
		fmt.Println("\nError while writing to 'socks.txt':", err)
		return
	}

	fmt.Println("Proxies downloaded successfully to 'socks.txt'.")
}
