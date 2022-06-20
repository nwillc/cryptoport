package gjson

import (
	"encoding/json"
	"github.com/nwillc/genfuncs"
)

func Marshal[V any](v *V) *genfuncs.Result[[]byte] {
	return genfuncs.NewResultError(json.Marshal(v))
}

func Unmarshal[R any](j []byte) *genfuncs.Result[*R] {
	var result R
	return genfuncs.NewResultError(&result, json.Unmarshal(j, &result))
}
