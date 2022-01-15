package git

import (
	"bufio"
	"bytes"
	"crypto/sha1"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

type Object struct {
	serializedData   []byte
	deserializedData []byte
	dataLen          int
	hash             string
	objType          string
}

func (o *Object) serialize(data io.Reader) error {
	buf, err := ioutil.ReadAll(data)
	if err != nil {
		return err
	}
	o.dataLen = len(buf)
	sd := []byte(strings.Join([]string{o.objType, fmt.Sprintf("%d", o.dataLen)}, " "))
	sd = append(sd, byte(0))
	sd = append(sd, buf...)
	o.hash = fmt.Sprintf("%x", sha1.Sum(sd))
	o.serializedData = sd
	return nil
}

func (o *Object) deserialize(data io.Reader) error {
	buf, err := ioutil.ReadAll(data)
	if err != nil {
		return err
	}
	ispace := bytes.IndexByte(buf, byte(32))
	if ispace < 0 {
		return fmt.Errorf("could not deserialize data")
	}
	o.objType = string(buf[0:ispace])

	inull := bytes.IndexByte(buf, byte(0))
	if inull < 0 {
		return fmt.Errorf("could not deserialize data")
	}
	o.dataLen, err = strconv.Atoi(string(buf[ispace+1 : inull]))
	if err != nil {
		return err
	}
	o.deserializedData = buf[inull+1:]

	return nil
}

func (o *Object) getHash() string {
	return o.hash
}

func (o *Object) getSerializedData() []byte {
	return o.serializedData
}

func (o *Object) GetObjType() string {
	return o.objType
}

func (o *Object) GetDeserializedData() string {
	return string(o.deserializedData)
}

type KVLM struct {
	KeyValues []KV
	Message   string
}

type KV struct {
	Key   string
	Value string
}

func ParseKeyValueListWithMessage(data io.Reader) *KVLM {
	s := bufio.NewScanner(data)
	curSlice := []string{}
	var key string
	messageFound := false
	mSlice := []string{}

	kvlm := KVLM{}
	kv := KV{}
	for s.Scan() {
		line := s.Text()
		if messageFound {
			mSlice = append(mSlice, line)
			continue
		}
		if strings.HasPrefix(line, " ") {
			curSlice = append(curSlice, line)
			continue
		} else if !strings.Contains(line, " ") {
			kv.Value = strings.Join(curSlice, "\n")
			fmt.Println(strings.Join(curSlice, "\n"))
			kvlm.KeyValues = append(kvlm.KeyValues, kv)
			messageFound = true
			continue
		} else if len(curSlice) > 0 {
			fmt.Println(key, strings.Join(curSlice, "\n"))
			kv.Value = strings.Join(curSlice, "\n")
			kvlm.KeyValues = append(kvlm.KeyValues, kv)
			curSlice = []string{}
		}
		lslice := strings.Split(line, " ")
		if len(lslice) > 1 {
			key = lslice[0]
			kv.Key = lslice[0]
			newSlice := []string{}
			for i := 1; i < len(lslice); i++ {
				newSlice = append(newSlice, lslice[i])
			}
			curSlice = append(curSlice, strings.Join(newSlice, " "))
		}
	}
	kvlm.Message = strings.Join(mSlice, "\n")
	return &kvlm
}
