package display

import (
	"github.com/HongJaison/go-admin/modules/utils"
	"github.com/HongJaison/go-admin/template/types"
	"strconv"
)

type FileSize struct {
	types.BaseDisplayFnGenerator
}

func init() {
	types.RegisterDisplayFnGenerator("filesize", new(FileSize))
}

func (f *FileSize) Get(args ...interface{}) types.FieldFilterFn {
	return func(value types.FieldModel) interface{} {
		size, _ := strconv.ParseUint(value.Value, 10, 64)
		return utils.FileSize(size)
	}
}
