//------------------------------------------------------------
// Author: FreeKnight
// Usage: 全局变量
//------------------------------------------------------------
package Global
//------------------------------------------------------------
import (
	CONFIG "../Config"
)
//------------------------------------------------------------
var G_Mapv4 	map[string]string = make(map[string]string);		// 域名-IP映射表
var G_Config  	CONFIG.SFKGoDNSServerConfig;				// 服务器配置
//------------------------------------------------------------
