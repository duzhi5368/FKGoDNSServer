//------------------------------------------------------------
// Author: FreeKnight
// Usage: 域名-IP配置文件
//------------------------------------------------------------
package Config
//------------------------------------------------------------
import (
	"bufio"
	"os"
	"strings"
	"fmt"

	COMMON "../Common"
)
//------------------------------------------------------------
// 从服务器更新拉取DomainConfig文件
func UpdateDomainConfig(config string) bool{
	return false;	// TEMP CODE
}
//------------------------------------------------------------
// 创建本地临时版DomainConfig文件
func CreateLocalTestDomainConfig(config string){

}
//------------------------------------------------------------
// 读取加载DomainConfig文件
func LoadDomainConfig(config string, mapv4 map[string]string)(bool){
	configPath := COMMON.GetCurrentDirectory() + config;
	inputFile, inputError := os.Open(configPath);
	if(inputError != nil){
		fmt.Printf("错误: 加载文件 %s 失败：%s\n", configPath, inputError);
		return false;
	} else{
		fmt.Printf("加载文件 %s 完毕...\n", configPath);
	}
	defer inputFile.Close();				// 延迟释放

	scanner := bufio.NewScanner(inputFile);
	for scanner.Scan() {					// 逐行扫描
		line := scanner.Text()
		if (! strings.HasPrefix(line, "#")) {		// # 为注释行
			fields := strings.Fields(line)
			if (len(fields) == 2){
				mapv4[fields[0]] = fields[1];
			}
		}
	}
	return true;
}
//------------------------------------------------------------
// 输出Log
func DumpDomainConfig(mapv4 map[string]string){
	fmt.Println("----------------------------------------------")
	fmt.Println("【域名IP表】")
	fmt.Println("----------------------------------------------")
	for domain, ip := range mapv4 {
		fmt.Print(domain);
		fmt.Println(" - " + ip);
	}
	fmt.Println("----------------------------------------------")
}
//------------------------------------------------------------