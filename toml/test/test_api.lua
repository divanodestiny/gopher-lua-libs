local toml = require("toml")
local inspect = require("inspect")

local text = [[
[table]
  a = 1
  b = false
  s = "hello"
]]
local result, err = toml.decode(text)
if err then
    error(err)
end
inspect(result, {newline="", indent=""})
inspect(result.table)
if not (result["table"]["a"] == 1 and
        result["table"]["b"] == false and
        result["table"]["s"] == "hello") then
    error("decode error")
end

print("done: toml.decode()")