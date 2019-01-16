package toml

import (
	"github.com/vadv/gopher-lua-libs/inspect"
	"github.com/yuin/gopher-lua"
	"testing"
)

func TestApi(t *testing.T){
	state := lua.NewState()
	Preload(state)
	inspect.Preload(state)
	if err := state.DoFile("./test/test_api.lua"); err != nil{
		t.Fatalf("execute test: %s\n", err.Error())
	}
}