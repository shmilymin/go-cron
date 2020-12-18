package test

import (
	"encoding/json"
	"go-cron/pkg/utils"
	"log"
	"strconv"
	"testing"
)

func TestSlice(t *testing.T) {
	var keys []string
	for i := 0; i < 10; i++ {
		keys = append(keys, "aaa"+strconv.Itoa(i))
	}
	tmp := keys[1:2]
	log.Println(tmp)
	tmp[0] = "aaa888"
	t.Log(keys)
	page := 1
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

type Atest struct {
	N1 int
	N2 int
}

func TestNum(t *testing.T) {
	var i int
	t.Log(i)
}
