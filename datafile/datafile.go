package datafile

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/juju/fslock"
)

var lock sync.Mutex

// Save saves a representation of v to the file at path.
func Save(path string, v interface{}) error {
	lock.Lock()
	defer lock.Unlock()
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	r, err := Marshal(v)
	if err != nil {
		return err
	}
	_, err = io.Copy(f, r)
	return err
}

// Marshal is a function that marshals the object into an
// io.Reader.
// By default, it uses the JSON marshaller.
var Marshal = func(v interface{}) (io.Reader, error) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

// Unmarshal is a function that unmarshals the data from the
// reader into the specified value.
// By default, it uses the JSON unmarshaller.
var Unmarshal = func(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}

// Load loads the file at path into v.
// Use os.IsNotExist() to see if the returned error is due
// to the file being missing.
func Load(path string, v interface{}) error {
	lock.Lock()
	defer lock.Unlock()
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return Unmarshal(f, v)
}

// LoadWithLock loads the file at path into v.
// Use os.IsNotExist() to see if the returned error is due
// to the file being missing.
func LoadWithLock(path string, v interface{}) error {
	lock := fslock.New(fmt.Sprintf("%s.lock", path))
	lockErr := lock.TryLock()
	if lockErr != nil {
		return lockErr
	}
	defer lock.Unlock()
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return Unmarshal(f, v)
}

// SaveWithLock saves a representation of v to the file at path.
func SaveWithLock(path string, v interface{}) error {
	lock := fslock.New(fmt.Sprintf("%s.lock", path))
	lockErr := lock.TryLock()
	if lockErr != nil {
		return lockErr
	}
	defer lock.Unlock()
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	r, err := Marshal(v)
	if err != nil {
		return err
	}
	_, err = io.Copy(f, r)
	return err
}
