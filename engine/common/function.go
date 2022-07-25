package common

import (
	"bytes"
	"encoding/gob"
	"os"
	"time"
)

// 获取当前路径
func GetProgrammePath() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return path, err
	}
	return path, err
}

func ThreeUnary(isTrue bool, a interface{}, b interface{}) interface{} {
	if isTrue {
		return a
	}
	return b
}

// 返回毫秒级时间戳
func UnixMilli() int64 {
	return time.Now().UnixNano() / 1e6
}

// 使用gob进行序列化与反序列化,性能不如json,部分类型还需要手动注册,不建议使用
func Serialization(arg interface{}) ([]byte, error) {
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	err := enc.Encode(arg)
	if err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}

func Deserialization(arg []byte) (interface{}, error) {
	var res interface{}
	buff := bytes.NewBuffer(arg)
	dec := gob.NewDecoder(buff)
	err := dec.Decode(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
