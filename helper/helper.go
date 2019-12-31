package helper

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"reflect"
)

type MapData map[string]interface{}

func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

func Tab(tab interface{}) (tabName string) {
	if tn := reflect.ValueOf(tab).MethodByName("TableName"); tn.IsValid() {
		tabName = tn.Call(nil)[0].String()
	}
	return
}
func TabA(tab interface{}) string { return fmt.Sprintf("%s as a", Tab(tab)) }
func TabB(tab interface{}) string { return fmt.Sprintf("%s as b", Tab(tab)) }
func TabC(tab interface{}) string { return fmt.Sprintf("%s as c", Tab(tab)) }
func TabD(tab interface{}) string { return fmt.Sprintf("%s as d", Tab(tab)) }
func TabE(tab interface{}) string { return fmt.Sprintf("%s as e", Tab(tab)) }
