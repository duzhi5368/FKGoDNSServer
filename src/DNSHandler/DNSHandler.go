//------------------------------------------------------------
// Author: FreeKnight
// Usage: DNS请求消息处理类
//------------------------------------------------------------
package DNSHandler
//------------------------------------------------------------
import (
	"time"
	"net"
	"fmt"

	DNS "github.com/miekg/dns"
	GLOBAL "../Global"
)
//------------------------------------------------------------
// DNS请求消息处理函数
func HandleRequest(w DNS.ResponseWriter, r *DNS.Msg) {
	domain := r.Question[0].Name
	if GLOBAL.G_Config.LogFlag {
		t := time.Now()
		ip, _, _ := net.SplitHostPort(w.RemoteAddr().String())
		fmt.Printf("%d-%02d-%02d_%02d:%02d:%02d\t 请求来自：%s\t 请求访问：%s", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), ip, domain)
	}
	val, ok := GLOBAL.G_Mapv4[domain];
	if (!ok) {
		if GLOBAL.G_Config.LogFlag {
			fmt.Println("  不知道的域名，不予处理");
		}
		return;
	}

	// 模拟DNS回馈消息
	m := new(DNS.Msg)
	m.SetReply(r)
	m.Authoritative = true
	rr1 := new(DNS.A)
	rr1.Hdr = DNS.RR_Header{Name: domain, Rrtype: DNS.TypeA, Class: DNS.ClassINET,
		Ttl: uint32(GLOBAL.G_Config.Ttl)}
	rr1.A = net.ParseIP(val)
	if GLOBAL.G_Config.LogFlag {
		fmt.Printf("\n解析IP = %s\n", val);
	}
	m.Answer = []DNS.RR{rr1};
	w.WriteMsg(m)
}