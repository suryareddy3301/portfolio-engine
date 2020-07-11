package main

import "testing"

func Test_server(t *testing.T) {
	server := NewServer()
	info := server.BasicInfo[0]
	if server == nil {
		t.Fatal("server is nil", info)
	}

}
