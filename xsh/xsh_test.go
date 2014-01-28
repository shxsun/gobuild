package xsh

import "testing"

func TestCall(t *testing.T) {
	ret, err := Call("echo", "hello")
	if err != nil {
		t.Error(err)
	}
	t.Log(ret.Trim())
}

func TestSession(t *testing.T) {
	session := NewSession("pwd")
	err := session.Call()
	if err != nil {
		t.Error(err)
	}
}
