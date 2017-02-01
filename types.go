package envStore

import (
	"sync"
)

type dictionary map[string]string

type Environment struct {
	data     dictionary
	mutex    sync.Mutex
	useMutex bool
	ignoreEmptyLines bool
}

type Config struct {
	FromFilePaths []string
	FromStrings   []string
	FromSystem    bool
	UseMutex      bool
	IgnoreEmptyLines bool
}
