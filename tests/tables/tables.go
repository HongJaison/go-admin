package tables

import "github.com/HongJaison/go-admin/plugins/admin/modules/table"

var Generators = map[string]table.Generator{
	"posts":    GetPostsTable,
	"authors":  GetAuthorsTable,
	"external": GetExternalTable,
}
