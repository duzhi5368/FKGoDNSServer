//------------------------------------------------------------
// Author: FreeKnight
// Usage: 服务器配置文件
//------------------------------------------------------------
package Config
//------------------------------------------------------------
import (
	"flag"
	"fmt"
)
//------------------------------------------------------------
// app 配置文件数据结构
type SFKGoDNSServerConfig struct {
	DomainConf string;
	Ttl        int;
	LogFlag    bool;
};
//------------------------------------------------------------
// 加载配置文件
func LoadServerConfigFromFile(config  *SFKGoDNSServerConfig) {
	flag.IntVar(&config.Ttl, "ttl", 86400, "默认TTL保存秒数")
	flag.BoolVar(&config.LogFlag, "log", true, "是否输出Log日志")
	flag.StringVar(&config.DomainConf, "conf", "/FKDomainConfig.conf", "域名IP映射表")
	flag.Parse()
}
//------------------------------------------------------------
// 输出Log
func DumpServerConfig(config SFKGoDNSServerConfig) {
	fmt.Println("----------------------------------------------")
	fmt.Println("【服务器配置文件】")
	fmt.Println("----------------------------------------------")
	fmt.Printf("TTL保存时间 : %d\n", config.Ttl)
	fmt.Printf("是否开启日志 : %t\n", config.LogFlag)
	fmt.Printf("域名-IP配置文件: %s\n", config.DomainConf)
	fmt.Println("----------------------------------------------")
}
//------------------------------------------------------------