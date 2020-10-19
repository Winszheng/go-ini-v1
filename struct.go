package go_ini_v1

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

