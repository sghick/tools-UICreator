package main

import (
    "bufio"
    "os"
	"fmt"
	"strings"
	"strconv"
)

type cmdfile struct {
	cmd string
	name string
	format string
	rem string
	content string
}

func main() {
	GetInput()
 }

func ExportError(error string) {
	fmt.Println(error)
 }

func GetInput() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("-----------------------------------------")
	fmt.Println("-----------------------------------------")
    fmt.Println("** 欢迎使用代码机器,不会偷懒的猿不是好猴子~")
    fmt.Println("** 小提示:\n**   查询可用模板: --s .\n**   退出: --exit")
	fmt.Println("-----------------------------------------")
	codeTypes :=[] string {"oc", "javascript"}
	for index := 0; index < len(codeTypes); index++ {
		fmt.Println(index, ": " + codeTypes[index])
	}
	fmt.Println("-----------------------------------------")
	fmt.Print("请从上面列表中选择一门语言类型,序号或者名字都看你心情:")
	codeType := ""
	for {
		input, _ := reader.ReadString('\n')
		// 兼容Windows And Mac
		input = strings.Replace(input, "\n", "", -1)
		idx, err := strconv.Atoi(input)
		if (err != nil) {
			for index := 0; index < len(codeTypes); index++ {
				if (input == codeTypes[index]) {
					codeType = input
					break
				}
			}
			if len(codeType) != 0 {
				break
			} else {
				fmt.Println("认真点,还不支持这种语言,要么自己加,要么换一个!")
			}
		} else {
			if (len(codeTypes) > idx) {
				codeType = codeTypes[idx]
				break
			} else {
				fmt.Println("认真点,没有这个号~~再给你一次机会!")
			}
		}
	}
	fmt.Println("您选择的语言是:" + codeType)

	for {
        fmt.Print("-> ")
        input, _ := reader.ReadString('\n')
        // 兼容Windows And Mac
		input = strings.Replace(input, "\n", "", -1)
        if BeginRun(input, codeType) == false {
			break
		}
	}
 }

func BeginRun(ipt string, codeType string) bool {
	// 如果以"--"开头的,直接解析
	if strings.Index(ipt, "--") == 0 {
		if strings.Index(ipt, "--e") == 0 {
			return false
		}
		if strings.Index(ipt, "--h") == 0 {

		}
		if strings.Index(ipt, "--s") == 0 {

		}
		return true
	}
	if strings.Index(ipt, "-") == 0 {	
		var ipts = strings.Split(ipt, " ")
		var cmd = ""
		var name = ""
		var ript = ""
		// 截取操作命令符
		if len(ipts) >= 1 {
			cmd = ipts[0][1:len(ipts[0])]
		}
		// 截取模板名字
		if len(ipts) >= 2 {
			name = ipts[1]
		}
		// 截取内容区域
		if len(ipts) >= 3 {
			ript = ipts[2]
		}
		var root = "./__templates/" 
		var fp = root + codeType + "/" + cmd + "_" + name + ".txt"
		var cf = Getfile(fp, cmd, name)
		var opt = Output(cf, cmd, name, ript)

		fmt.Println("=== ↓↓↓↓↓ === 代码已经生成,并已自动复制到粘贴板中,请收好 === ↓↓↓↓↓ ====")
		fmt.Println("")
		fmt.Println(opt)
		fmt.Println("")
		fmt.Println("=== ↑↑↑↑↑ === 代码已经生成,并已自动复制到粘贴板中,请收好 === ↑↑↑↑↑ ===")
	}
	return true
 }

func Getfile(fp string, cmd string, name string) cmdfile {
	var cf cmdfile
	cf.cmd = cmd
	cf.name = name
	txt := ReadAll(fp)
	lines := strings.Split(txt, "\n")

	var lastKey string
	for index := 0; index < len(lines); index++ {
		line := lines[index]
		if strings.Index(line, "#") == 0 {
			lastKey = line
			continue
		}
		if strings.Index(lastKey, "format") > 0 {
			cf.format = line
			lastKey = ""
		}
		if strings.Index(lastKey, "rem") > 0 {
			cf.rem = line
			lastKey = ""
		}
		if strings.Index(lastKey, "content") > 0 {
			cf.content = cf.content + line + "\n"
		}
	}
	return cf
 }

func ReadAll(fp string) string {
	var result string
	// 取文件全部内容
	file, error := os.Open(fp);
	if error != nil {
		fmt.Println(error);
	}
	buf2 := make([]byte, 1024);
	ix := 0;
	for {
		//ReadAt从指定的偏移量开始读取，不会改变文件偏移量
		len, _ := file.ReadAt(buf2, int64(ix));
		ix = ix + len;
		if len == 0 {
			break;
		}
		result = string(buf2)
	}
	file.Close()
	return result
 }

func Output(cf cmdfile, cmd string, name string, input string) string {
	var result = cf.content
	// 替换类名
	result = strings.Replace(result, "<.name.>", name, -1)
	// 获取输入格式
	ipts := strings.Split(input, ",")
	fmts := strings.Split(cf.format, ",")
	for index := 0; index < len(ipts); index++ {
		if len(fmts) > index {
			result = strings.Replace(result, "<." + fmts[index] + ".>", ipts[index], -1)
		}
	}
	return result
 }
