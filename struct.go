package go_ini_v1

// struct for configuration file
//
// 配置文件结构体
type File struct {
	sections map[string]*Section
}

// struct for section
//
// 配置文件分区结构体
type Section struct {
	value map[string]string
}

// Section: get the section
//
// 获取分区
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
//
// 根据键获取值
func (s *Section) Key(key string) string {
	return s.value[key]
}

