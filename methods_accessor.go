package envStore

import (
	"bytes"
	"errors"
	"strconv"
	"strings"
)

var (
	KeyNotFoundErr = errors.New("environment variable not found")
)

func (e *Environment) Get(key string) (value string, err error) {
	e.readLockIfNecessary()
	defer e.readUnlockIfNecessary()

	value, ok := e.data[strings.ToUpper(key)]
	if !ok {
		err = KeyNotFoundErr
	}

	return
}

func (e *Environment) GetByteSlice(key string) (value []byte, err error) {
	rawValue, err := e.Get(key)
	if err != nil {
		return
	}

	value = []byte(rawValue)
	return
}

func (e *Environment) GetInt(key string) (value int, err error) {
	rawValue, err := e.Get(key)
	if err != nil {
		return
	}

	value, err = strconv.Atoi(rawValue)
	return
}

func (e *Environment) GetUint(key string) (value uint, err error) {
	rawValue, err := e.Get(key)
	if err != nil {
		return
	}

	value64, err := strconv.ParseUint(rawValue, 10, 0)
	value = uint(value64)
	return
}

func (e *Environment) GetFloat(key string) (value float64, err error) {
	rawValue, err := e.Get(key)
	if err != nil {
		return
	}

	value, err = strconv.ParseFloat(rawValue, 64)
	return
}

func (e *Environment) GetBool(key string) (value bool, err error) {
	rawValue, err := e.Get(key)
	if err != nil {
		return
	}

	value, err = strconv.ParseBool(rawValue)
	return
}

func mustPanic(err error, key string) {
	msg := err.Error()

	if err == KeyNotFoundErr {
		msg += " [" + strings.ToUpper(key) + "]"
	}

	panic(msg)
}

func (e *Environment) MustGet(key string) (value string) {
	value, err := e.Get(key)
	if err != nil {
		mustPanic(err, key)
	}

	return
}

func (e *Environment) MustGetByteSlice(key string) (value []byte) {
	value, err := e.GetByteSlice(key)
	if err != nil {
		mustPanic(err, key)
	}

	return
}

func (e *Environment) MustGetInt(key string) (value int) {
	value, err := e.GetInt(key)
	if err != nil {
		mustPanic(err, key)
	}

	return
}

func (e *Environment) MustGetUint(key string) (value uint) {
	value, err := e.GetUint(key)
	if err != nil {
		mustPanic(err, key)
	}

	return
}

func (e *Environment) MustGetFloat(key string) (value float64) {
	value, err := e.GetFloat(key)
	if err != nil {
		mustPanic(err, key)
	}

	return
}

func (e *Environment) MustGetBool(key string) (value bool) {
	value, err := e.GetBool(key)
	if err != nil {
		mustPanic(err, key)
	}

	return
}

func (e *Environment) Set(key, value string) {
	e.writeLockIfNecessary()
	defer e.writeUnlockIfNecessary()

	e.data[strings.ToUpper(key)] = value
}

func (e *Environment) Unset(key string) {
	e.writeLockIfNecessary()
	defer e.writeUnlockIfNecessary()

	delete(e.data, strings.ToUpper(key))
}

func (e *Environment) Clear() {
	e.writeLockIfNecessary()
	defer e.writeUnlockIfNecessary()

	e.data = make(dictionary)
}

func (e *Environment) Contains(key string) bool {
	e.readLockIfNecessary()
	defer e.readUnlockIfNecessary()

	_, ok := e.data[strings.ToUpper(key)]
	return ok
}

func (e *Environment) Count() int {
	e.readLockIfNecessary()
	defer e.readUnlockIfNecessary()

	return len(e.data)
}

func (e *Environment) Keys() []string {
	e.readLockIfNecessary()
	defer e.readUnlockIfNecessary()

	i := 0
	keys := make([]string, len(e.data))
	for k := range e.data {
		keys[i] = k
		i++
	}

	return keys
}

func (e *Environment) Values() []string {
	e.readLockIfNecessary()
	defer e.readUnlockIfNecessary()

	i := 0
	values := make([]string, len(e.data))
	for _, v := range e.data {
		values[i] = v
		i++
	}

	return values
}

func (e *Environment) Pairs() [][]string {
	e.readLockIfNecessary()
	defer e.readUnlockIfNecessary()

	i := 0
	pairs := make([][]string, len(e.data))
	for k, v := range e.data {
		pairs[i] = make([]string, 2)
		pairs[i][0], pairs[i][1] = k, v
		i++
	}

	return pairs
}

func (e *Environment) String() string {
	e.readLockIfNecessary()
	defer e.readUnlockIfNecessary()

	var buffer bytes.Buffer
	var passedFirstIteration bool

	for k, v := range e.data {
		if passedFirstIteration {
			buffer.WriteRune('\n')
		} else {
			passedFirstIteration = true
		}
		buffer.WriteString(k)
		buffer.WriteRune('=')
		buffer.WriteString(v)
	}

	return buffer.String()
}
