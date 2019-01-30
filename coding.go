package main

import (
    "bufio"
    "os"
	"fmt"
	"strings"
	"strconv"
	"path/filepath"
)

type cmdfile struct {
	cmd string
	name string
	params string
	rem string
	content string
}

var _root string

func main() {
	GetInput()
 }

func ExportError(error string) {
	fmt.Println(error)
 }

func GetInput() {
	reader := bufio.NewReader(os.Stdin)
	PrintWelcome()
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
		var fp = GetRootDoc() + "__templates/" + codeType + "/" + cmd + "_" + name + ".txt"
		var cf, err = Getfile(fp, cmd, name)
		if err != nil {
			fp = GetRootDoc() + "__templates/" + codeType + "/" + cmd+ "_" + ".txt"
			fmt.Println(err);
			fmt.Println("使用默认模板'" + fp + "'...")
			cf, err = Getfile(fp, cmd, name)
			if err != nil {
				fmt.Println("抱歉,没有找到相应的模板~")
			}
		}

		if err == nil {
			var opt = Output(cf, cmd, name, ript)
			fmt.Println("=== ↓↓↓↓↓ === 代码已经生成,并已自动复制到粘贴板中,请收好 === ↓↓↓↓↓ ====")
			fmt.Println("")
			fmt.Println(opt)
			fmt.Println("")
			fmt.Println("=== ↑↑↑↑↑ === 代码已经生成,并已自动复制到粘贴板中,请收好 === ↑↑↑↑↑ ===")
		}
	}
	return true
 }

func GetRootDoc() string {
	if strings.Count(_root, "") > 1 {
		return _root
	}
	// 先尝试使用'./'为根目录,如果没有发现'coding.go'文件,则使用运行目录
	var result = "./"
	file, err := os.Open(result + "coding.go");
	if (file == nil) || (err != nil) {
		result, err := filepath.Abs(filepath.Dir(os.Args[0]))
		fmt.Println("result:", result)
		if err != nil {
			fmt.Println("未找到正确的目录!")
		}
		_root = result + "/"
	}
	file.Close()
	_root = result
	return _root;
}

func PrintWelcome() {
	fp := GetRootDoc() + "welcome.txt"
	txt, err := ReadAll(fp)
	if err != nil {
		return
	}
	lines := strings.Split(txt, "\n")
	for index := 0; index < len(lines); index++ {
		line := lines[index]
		fmt.Println(line)
	}
}

func Getfile(fp string, cmd string, name string) (cmdfile, error) {
	var cf cmdfile
	cf.cmd = cmd
	cf.name = name
	txt, err := ReadAll(fp)
	if err != nil {
		return cf, err
	}
	lines := strings.Split(txt, "\n")

	var lastKey string
	for index := 0; index < len(lines); index++ {
		line := lines[index]
		if strings.Index(line, "#") == 0 {
			lastKey = line
			continue
		}
		if strings.Index(lastKey, "params") > 0 {
			cf.params = line
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
	return cf, nil
 }

func ReadAll(fp string) (string, error) {
	var result string
	// 取文件全部内容
	file, err := os.Open(fp);
	if err != nil {
		return result, err
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
	return result, nil
 }

func Output(cf cmdfile, cmd string, name string, input string) string {
	var result = cf.content
	// 替换类名
	result = strings.Replace(result, "<.name.>", name, -1)
	// 获取输入格式
	ipts := strings.Split(input, ",")
	fmts := strings.Split(cf.params, ",")
	for index := 0; index < len(ipts); index++ {
		if len(fmts) > index {
			result = strings.Replace(result, "<." + fmts[index] + ".>", ipts[index], -1)
		}
	}
	return result
 }
