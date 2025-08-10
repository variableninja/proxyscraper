package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	urls := []string{
		"https://raw.githubusercontent.com/jetkai/proxy-list/main/online-proxies/txt/proxies-https.txt",
		"https://raw.githubusercontent.com/jetkai/proxy-list/main/online-proxies/txt/proxies-http.txt",
		"https://api.proxyscrape.com/v2/?request=getproxies&protocol=http&timeout=10000&country=all",
		"https://raw.githubusercontent.com/jetkai/proxy-list/main/archive/txt/proxies-https.txt",
		"https://raw.githubusercontent.com/jetkai/proxy-list/main/archive/txt/proxies-http.txt",
		"https://raw.githubusercontent.com/roosterkid/openproxylist/main/HTTPS_RAW.txt",
		"https://raw.githubusercontent.com/monosans/proxy-list/main/proxies/http.txt",
		"https://raw.githubusercontent.com/TheSpeedX/PROXY-List/master/http.txt",
		"https://www.proxy-list.download/api/v1/get?type=http",
		"https://www.proxy-list.download/api/v1/get?type=https",
		"https://www.proxyscan.io/download?type=http",
		"https://www.proxyscan.io/download?type=https",
		"https://api.openproxylist.xyz/http.txt",
		"http://ab57.ru/downloads/proxyold.txt",
		"http://alexa.lr2b.com/proxylist.txt",
		"http://rootjazz.com/proxies/proxies.txt",
		"http://spys.me/proxy.txt",
		"http://worm.rip/http.txt",
		"http://www.httptunnel.ge/ProxyListForFree.aspx",
		"http://www.proxyserverlist24.top/feeds/posts/default",
		"https://api.proxyscrape.com/?request=displayproxies&proxytype=http",
		"https://api.proxyscrape.com/?request=getproxies&proxytype=http&timeout=6000&country=all&ssl=yes&anonymity=all",
		"https://dstat.one/proxies.php?id=2",
		"https://dstat.one/proxies.php?id=3",
		"https://free-proxy-list.net/",
		"https://free-proxy-list.net/anonymous-proxy.html",
		"https://free-proxy-list.net/uk-proxy.html",
		"https://www.proxy-list.download/api/v1/get?type=http",
		"https://www.proxy-list.download/api/v1/get?type=https",
		"https://www.proxyscan.io/download?type=http",
		"https://www.proxyscan.io/download?type=https",
	}

	proxies := ""
	client := &http.Client{}
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

	outFile, err := os.Create("proxies/http.txt")
	if err != nil {
		fmt.Println("\nError while creating 'http.txt':", err)
		return
	}
	defer outFile.Close()

	_, err = outFile.WriteString(proxies)
	if err != nil {
		fmt.Println("\nError while writing to 'http.txt':", err)
		return
	}

	fmt.Println("Proxies downloaded successfully to 'http.txt'.")
}
