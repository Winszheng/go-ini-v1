package go_ini_v1

import (
	"errors"
	"os"
	"syscall"
	"time"
)

// interface for listener
type Listener interface {
	Listen(string) error
}

// struct for listener
type List struct {

}

// handle listening
func (l *List)Listen(filename string) error {
	fileInfo, err := os.Stat(filename)
	if err!= nil {
		return err
	}
	fileSys := fileInfo.Sys().(*syscall.Stat_t)
	nanotime := fileSys.Mtim
	size := fileInfo.Size()

	for {
		fileInfo, err = os.Stat(filename)
		if err!= nil {
			return errors.New("error when listening")
		}

		time.Sleep(1*time.Second)
		if nanotime!=fileSys.Mtim || size != fileInfo.Size(){
			return nil
		}
	}
	return nil
}

// Watch: call Listen
func Watch(filename string, listen List)(*File, error){
	err := listen.Listen(filename)
	if err != nil {
		return nil, errors.New("error when call the Listen")
	}
	f, err := Load(filename)    // reload and return
	if err!= nil {
		return nil, err
	}
	return f, nil
}
