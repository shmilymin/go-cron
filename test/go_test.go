package test

import (
	"encoding/json"
	"go-cron/pkg/utils"
	"strconv"
	"testing"
)

func TestSlice(t *testing.T) {
	var keys []string
	for i := 0; i < 10; i++ {
		keys = append(keys, "aaa"+strconv.Itoa(i))
	}
	t.Log(keys)
	page := 2
	limit := 4
	idx1 := (page - 1) * limit
	idx2 := limit * page
	vs := keys[idx1:idx2]
	t.Log(vs)
	jsonByte, e := json.Marshal(vs)
	if e != nil {
		t.Error(e)
	}
	t.Log(string(jsonByte))
	var newKeys []string
	if e := json.Unmarshal(jsonByte, &newKeys); e != nil {
		t.Error(e)
	}
	t.Log(newKeys)
}

func TestDate(t *testing.T) {
	t.Log(utils.Now())
}
