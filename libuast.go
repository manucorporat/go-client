package bblfsh

import "C"
import "reflect"
import "unsafe"

func init() {
	c.create_go_node_api()
}

func readAttribute(ptr unsafe.Pointer, attribute string) reflect.Value {
	obj := *((*interface{})(ptr))
	value := reflect.ValueOf(obj)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	return value
}

func readString(ptr unsafe.Pointer, attribute string) string {
	return readAttribute(ptr, attribute).String()
}

func readLen(ptr unsafe.Pointer, attribute string) int {
	return readAttribute(ptr, attribute).Len()
}

func readIndex(ptr unsafe.Pointer, attribute string, index int) reflect.Value {
	return readAttribute(ptr, attribute).Index(index)
}

func getInternalType(ptr unsafe.Pointer) string {
	return readString(ptr, "InternalType")
}

func getPropertiesSize(prt unsafe.Pointer) int {
	return readAttribute(ptr, attribute).Len()
}

func getToken(ptr unsafe.Pointer) string {
	return readString(ptr, "Token")
}

func getChildrenSizeToken(ptr unsafe.Pointer) int {
	return readLen(ptr, "Children")
}

func getChild(ptr unsafe.Pointer, index int) interface{} {
	return readIndex(ptr, "Children", index).Pointer()
}

func getRolesSizeToken(ptr unsafe.Pointer) int {
	return readLen(ptr, "Roles")
}
func getRole(ptr unsafe.Pointer, index int) uint16 {
	return uint16(readIndex(ptr, "Roles", index).Uint())
}
