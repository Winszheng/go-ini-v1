package go_ini_v1

import (
	"bufio"
	"errors"
	"io"
	"os"
	"runtime"
	"strings"
)


var annotationSymbol = '#'

// init: Linux: '#' while Windows: ';'

// init: Linux系统: '#', Windows系统: ';'
func Init() {
	if runtime.GOOS == "windows" {
		annotationSymbol = ';'
	}
}

// Load loads and parses from INI data sources.
// It will return error if list contains nonexistent files.
//
// 加载配置文件并绑定到结构体变量
func Load(filename string) (*File, error) {
	Init()

	f := File{sections: map[string]*Section{}}

	var count int = 0
	fin, err := os.Open(filename)
	if err != nil {
		panic(err)
		return nil, errors.New("error when open file")
	}
	defer fin.Close()

	/*create a Reader*/
	rd := bufio.NewReader(fin)

	/*read the file and stop when meet err or EOF*/
	curSection := "default"
	f.sections[curSection] = &Section{value: map[string]string{}}

	for {
		line, err := rd.ReadString('\n')
		if err != nil || err == io.EOF {
			break
		}
		count++

		if line[0] == '['{
			// is section
			index := strings.Index(line, "]")
			curSection = strings.TrimSpace(line[1:index])
			 f.Section(curSection)
			 f.sections[curSection] = &Section{value: map[string]string{}}

		}else if line[0] == uint8(annotationSymbol){
			// comment line
			continue
		}else{
			// key: value
			index := strings.Index(line, "=")
			if index < 0 {
				continue
			}

			key := strings.TrimSpace(line[:index])
			if len(key) == 0 {
				continue
			}

			val := strings.TrimSpace(line[index+1:])
			if len(val) == 0 {
				continue
			}

			f.sections[curSection].value[key] = val
		}
	}
	return &f, nil
}


