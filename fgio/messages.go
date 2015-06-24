
package fgio

import (
	"strconv"
	"fmt"
)



type MessageFrame struct {
	Node string ` json:"path" `
	Name string ` json:"name" `
	Type string ` json:"type" `
	Index int ` json:"index" `
	NChildren int ` json:"nChildren" `
	RawValue interface{} ` json:"value" `
}

// Returns RawValue as string regardless of Type
func (me MessageFrame) StrValue() string {

	switch me.Type {
	case "string":
		return me.RawValue.(string)

	case "double":
		return fmt.Sprintf("%.f", me.RawValue)

	case "float":
		return strconv.FormatFloat(me.RawValue.(float64), 'f', 20, 64)
	}

	return "#### OOOPS ##########"
}
