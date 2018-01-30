//------------------------------------------------------------
// Author: FreeKnight
// Usage: Main入口函数
//------------------------------------------------------------
package main
//------------------------------------------------------------
import (
	"log"
	"os"
	"os/signal"
	"syscall"
	DNS "github.com/miekg/dns"

	CONFIG "./Config"
	DNSHANDLE "./DNSHandler"
	GLOBAL "./Global"
)
//------------------------------------------------------------
func main() {
	// 加载配置
	CONFIG.LoadServerConfigFromFile(&GLOBAL.G_Config);
	// Debug输出配置
	CONFIG.DumpServerConfig(GLOBAL.G_Config);
	// 更新DomainConfig
	if(!CONFIG.UpdateDomainConfig(GLOBAL.G_Config.DomainConf)){
		CONFIG.CreateLocalTestDomainConfig(GLOBAL.G_Config.DomainConf);
	}
	// 加载DomainConfig
	if(CONFIG.LoadDomainConfig(GLOBAL.G_Config.DomainConf, GLOBAL.G_Mapv4)){
		CONFIG.DumpDomainConfig(GLOBAL.G_Mapv4);
	}
	// 开启DNS监听
	DNS.HandleFunc(".", DNSHANDLE.HandleRequest)

	// 开启线程监听UDP 53端口
	go func() {
		srv := &DNS.Server{Addr: ":53", Net: "udp"}
		err := srv.ListenAndServe()
		if err != nil {
			log.Fatal("开启UDP 53端口监听失败 %s\n", err.Error())
		}
	}()

	// 开启线程监听TCP 53端口
	go func() {
		srv := &DNS.Server{Addr: ":53", Net: "tcp"}
		err := srv.ListenAndServe()
		if err != nil {
			log.Fatal("开启TCP 53端口监听失败 %s\n", err.Error())
		}
	}()

	// 监听中断消息
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case s := <-sig:
			log.Fatalf("收到中断消息：(%d)， 程序即将退出 \n", s)
		}
	}
}
//------------------------------------------------------------