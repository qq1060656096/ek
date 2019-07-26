package ek

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type UserEventTest struct {
	Uid int64 `json:"uid"`
	Name string `json:"name"`
}


func TestNew(t *testing.T) {
	ue := UserEventTest{
		Uid: 1,
		Name: "小明",
	}
	ip := "127.0.0.1"
	var opTime int64 = 1564045436// 2019-07-25 17:03:56
	e := NewEventRaw("USER_REGISTER", "1", ue, ip, opTime)
	b, err := json.Marshal(e)
	fmt.Println(string(b))
	if err != nil {
		t.Error("event.New json.Marshal error:" , err)
	}
	assert.Equal(t, string(b), e.String())
}