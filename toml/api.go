package toml

import (
	"encoding/json"
	LJson "github.com/vadv/gopher-lua-libs/json"
	"fmt"
	"github.com/pelletier/go-toml"
	"github.com/yuin/gopher-lua"
)


// transfer toml to json
// then transfer json to lua table
func Decode(L *lua.LState) int{
	str := L.CheckString(1)
	tree, err := toml.LoadBytes([]byte(str))
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
	}

	treeMap := tree.ToMap()
	j, err := json.Marshal(treeMap)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
	}

	value, err := LJson.ValueDecode(L, j)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}
	L.Push(value)
	return 1

}

func fromTOML(L *lua.LState, val interface{}) lua.LValue{
	fmt.Println("llll")
	fmt.Println(val)
	fmt.Println("lll")
	switch converted := val.(type){
	case bool:
		return lua.LBool(converted)
	case float64:
		return lua.LNumber(converted)
	case int:
		return lua.LNumber(converted)
	case int64:
		return lua.LNumber(converted)
	case string:
		return lua.LString(converted)
	case []interface{}:
		arr := L.CreateTable(len(converted), 0)
		for _, item := range converted{
			arr.Append(fromTOML(L, item))
		}
		return arr
	case map[interface{}]interface{}:
		tb := L.CreateTable(0, len(converted))
		for key, item := range converted{
			tb.RawSetH(fromTOML(L, key), fromTOML(L, item))
		}
		return tb
	case interface{}:
		fmt.Println("converted")
		if v, ok := converted.(bool); ok {
			return lua.LBool(v)
		}
		if v, ok := converted.(float64); ok {
			return lua.LNumber(v)
		}
		if v, ok := converted.(string); ok {
			return lua.LString(v)
		}
	}
	return lua.LNil
}