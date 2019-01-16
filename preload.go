package libs

import (
	cert_util "github.com/vadv/gopher-lua-libs/cert_util"
	chef "github.com/vadv/gopher-lua-libs/chef"
	cmd "github.com/vadv/gopher-lua-libs/cmd"
	crypto "github.com/vadv/gopher-lua-libs/crypto"
	db "github.com/vadv/gopher-lua-libs/db"
	filepath "github.com/vadv/gopher-lua-libs/filepath"
	goos "github.com/vadv/gopher-lua-libs/goos"
	http "github.com/vadv/gopher-lua-libs/http"
	humanize "github.com/vadv/gopher-lua-libs/humanize"
	inspect "github.com/vadv/gopher-lua-libs/inspect"
	ioutil "github.com/vadv/gopher-lua-libs/ioutil"
	json "github.com/vadv/gopher-lua-libs/json"
	plugin "github.com/vadv/gopher-lua-libs/plugin"
	regexp "github.com/vadv/gopher-lua-libs/regexp"
	runtime "github.com/vadv/gopher-lua-libs/runtime"
	storage "github.com/vadv/gopher-lua-libs/storage"
	strings "github.com/vadv/gopher-lua-libs/strings"
	tac "github.com/vadv/gopher-lua-libs/tac"
	tcp "github.com/vadv/gopher-lua-libs/tcp"
	telegram "github.com/vadv/gopher-lua-libs/telegram"
	template "github.com/vadv/gopher-lua-libs/template"
	time "github.com/vadv/gopher-lua-libs/time"
	xmlpath "github.com/vadv/gopher-lua-libs/xmlpath"
	yaml "github.com/vadv/gopher-lua-libs/yaml"
	zabbix "github.com/vadv/gopher-lua-libs/zabbix"

	lua "github.com/yuin/gopher-lua"
)

// Preload(): preload indicated gopher lua packages
// If options is nil, all packages will be preload
func Preload(L *lua.LState, options []string) {
	if options == nil {
		time.Preload(L)
		strings.Preload(L)
		filepath.Preload(L)
		ioutil.Preload(L)
		http.Preload(L)
		regexp.Preload(L)
		tac.Preload(L)
		inspect.Preload(L)
		yaml.Preload(L)
		plugin.Preload(L)
		cmd.Preload(L)
		json.Preload(L)
		tcp.Preload(L)
		xmlpath.Preload(L)
		db.Preload(L)
		cert_util.Preload(L)
		runtime.Preload(L)
		telegram.Preload(L)
		zabbix.Preload(L)
		crypto.Preload(L)
		goos.Preload(L)
		storage.Preload(L)
		humanize.Preload(L)
		chef.Preload(L)
		template.Preload(L)
	}
	for _, op := range options {
		switch op {
		case "time":
			time.Preload(L)
		case "strings":
			strings.Preload(L)
		case "filepath":
			filepath.Preload(L)
		case "ioutil":
			ioutil.Preload(L)
		case "http":
			http.Preload(L)
		case "regexp":
			regexp.Preload(L)
		case "tac":
			tac.Preload(L)
		case "inspect":
			inspect.Preload(L)
		case "yaml":
			yaml.Preload(L)
		case "plugin":
			plugin.Preload(L)
		case "cmd":
			cmd.Preload(L)
		case "json":
			json.Preload(L)
		case "tcp":
			tcp.Preload(L)
		case "xmlpath":
			xmlpath.Preload(L)
		case "db":
			db.Preload(L)
		case "cert_util":
			cert_util.Preload(L)
		case "runtime":
			runtime.Preload(L)
		case "telegram":
			telegram.Preload(L)
		case "zabbix":
			zabbix.Preload(L)
		case "crypto":
			crypto.Preload(L)
		case "goos":
			goos.Preload(L)
		case "storage":
			storage.Preload(L)
		case "humanize":
			humanize.Preload(L)
		case "chef":
			chef.Preload(L)
		case "template":
			template.Preload(L)
		}
	}
}
