//------------------------------------------------------------
// Author: FreeKnight
// Usage: 基本函数库
//------------------------------------------------------------
package Common
//------------------------------------------------------------
import (
	"path/filepath"
	"os"
	"log"
	"strings"
)
//------------------------------------------------------------
// 获取当前程序工作目录
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
//------------------------------------------------------------