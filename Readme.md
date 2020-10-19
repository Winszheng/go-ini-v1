# 设计文档

### 需求说明

开发一个可以读ini文件的程序包。

### 任务目标

- 熟悉程序包的编写习惯（idioms）和风格（convetions）
- 熟悉 io 库操作
- 使用测试驱动的方法
- 简单 Go 程使用
- 事件通知

### 设计思路

#### 编写测试文件

测试驱动开发，编写需要的测试文件test.go

```
package main

import (
	"fmt"
	myini "github.com/Winszheng/go-ini-v1"
	"os"
)

func main() {
	filename := "demo01/my.ini"
	f, err := myini.Load(filename)
	if err!= nil {
		fmt.Println("failed to read file")
		os.Exit(1)
	}

	//典型的读取操作
	fmt.Println("App Mode: ", f.Section("").Key("app_mode"))
	fmt.Println("Data Path: ", f.Section("paths").Key("data"))
	fmt.Println("Server Protocol: ", f.Section("server").Key("protocol"))
	fmt.Println("Server Protocol: ", f.Section("server").Key("http_port"))
	fmt.Println("Server Protocol: ", f.Section("server").Key("enforce_domain"))


	var listen myini.List
	f, err = myini.Watch(filename, listen)
	if err!= nil {
		fmt.Println("failed to read file")
		os.Exit(1)
	}
	fmt.Println(filename, " has been changed, here are newest version: ")
	//典型的读取操作
	fmt.Println("App Mode: ", f.Section("").Key("app_mode"))
	fmt.Println("Data Path: ", f.Section("paths").Key("data"))
	fmt.Println("Server Protocol: ", f.Section("server").Key("protocol"))
	fmt.Println("Server Protocol: ", f.Section("server").Key("http_port"))
	fmt.Println("Server Protocol: ", f.Section("server").Key("enforce_domain"))
}
```

编写结束进行简单测试，至此，可以整体确定需要实现什么效果。

#### 数据结构

对于“数据结构”与“主要函数说明”，均可直接查阅godoc，如果你对此不感兴趣，可以直接调到后面的“如何使用”。

```
// struct for configuration file
type File struct {
   sections map[string]*Section
}

// struct for section
type Section struct {
   value map[string]string
}

// Section: get the section
func(f *File) Section(name string) *Section {
   if name == "" {
      name = "default"
   }
   if s, ok := f.sections[name]; ok {
      return s
   }
   f.sections[name] = &Section{value: map[string]string{}}
   return f.sections[name]
}

// Key: get value by key
func (s *Section) Key(key string) string {
   return s.value[key]
}
```

#### 主要函数说明

没有必要罗列全部，只写一个Load函数吧

```
// Load loads and parses from INI data sources.
// It will return error if list contains nonexistent files.
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
```

### 如何使用

首先，新建一个自己的项目



