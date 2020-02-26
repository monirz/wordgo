package main

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "moc.h"
import "C"
import (
	"strings"
	"unsafe"

	"github.com/therecipe/qt"
	std_core "github.com/therecipe/qt/core"
)

func cGoFreePacked(ptr unsafe.Pointer) { std_core.NewQByteArrayFromPointer(ptr).DestroyQByteArray() }
func cGoUnpackString(s C.struct_Moc_PackedString) string {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_Moc_PackedString) []byte {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		gs := C.GoString(s.data)
		return *(*[]byte)(unsafe.Pointer(&gs))
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
}
func unpackStringList(s string) []string {
	if len(s) == 0 {
		return make([]string, 0)
	}
	return strings.Split(s, "¡¦!")
}

type QmlBridge_ITF interface {
	std_core.QObject_ITF
	QmlBridge_PTR() *QmlBridge
}

func (ptr *QmlBridge) QmlBridge_PTR() *QmlBridge {
	return ptr
}

func (ptr *QmlBridge) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QmlBridge) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQmlBridge(ptr QmlBridge_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlBridge_PTR().Pointer()
	}
	return nil
}

func NewQmlBridgeFromPointer(ptr unsafe.Pointer) (n *QmlBridge) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(QmlBridge)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *QmlBridge:
			n = deduced

		case *std_core.QObject:
			n = &QmlBridge{QObject: *deduced}

		default:
			n = new(QmlBridge)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackQmlBridge9911b0_Constructor
func callbackQmlBridge9911b0_Constructor(ptr unsafe.Pointer) {
	this := NewQmlBridgeFromPointer(ptr)
	qt.Register(ptr, this)
	this.ConnectAdd(this.add)
}

//export callbackQmlBridge9911b0_SendToQml
func callbackQmlBridge9911b0_SendToQml(ptr unsafe.Pointer, source C.struct_Moc_PackedString, action C.struct_Moc_PackedString, data C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "sendToQml"); signal != nil {
		(*(*func(string, string, string))(signal))(cGoUnpackString(source), cGoUnpackString(action), cGoUnpackString(data))
	}

}

func (ptr *QmlBridge) ConnectSendToQml(f func(source string, action string, data string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "sendToQml") {
			C.QmlBridge9911b0_ConnectSendToQml(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "sendToQml")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "sendToQml"); signal != nil {
			f := func(source string, action string, data string) {
				(*(*func(string, string, string))(signal))(source, action, data)
				f(source, action, data)
			}
			qt.ConnectSignal(ptr.Pointer(), "sendToQml", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sendToQml", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectSendToQml() {
	if ptr.Pointer() != nil {
		C.QmlBridge9911b0_DisconnectSendToQml(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "sendToQml")
	}
}

func (ptr *QmlBridge) SendToQml(source string, action string, data string) {
	if ptr.Pointer() != nil {
		var sourceC *C.char
		if source != "" {
			sourceC = C.CString(source)
			defer C.free(unsafe.Pointer(sourceC))
		}
		var actionC *C.char
		if action != "" {
			actionC = C.CString(action)
			defer C.free(unsafe.Pointer(actionC))
		}
		var dataC *C.char
		if data != "" {
			dataC = C.CString(data)
			defer C.free(unsafe.Pointer(dataC))
		}
		C.QmlBridge9911b0_SendToQml(ptr.Pointer(), C.struct_Moc_PackedString{data: sourceC, len: C.longlong(len(source))}, C.struct_Moc_PackedString{data: actionC, len: C.longlong(len(action))}, C.struct_Moc_PackedString{data: dataC, len: C.longlong(len(data))})
	}
}

//export callbackQmlBridge9911b0_SendToGo
func callbackQmlBridge9911b0_SendToGo(ptr unsafe.Pointer, data C.struct_Moc_PackedString) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "sendToGo"); signal != nil {
		tempVal := (*(*func(string) []string)(signal))(cGoUnpackString(data))
		return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "¡¦!")), len: C.longlong(len(strings.Join(tempVal, "¡¦!")))}
	}
	tempVal := make([]string, 0)
	return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "¡¦!")), len: C.longlong(len(strings.Join(tempVal, "¡¦!")))}
}

func (ptr *QmlBridge) ConnectSendToGo(f func(data string) []string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "sendToGo"); signal != nil {
			f := func(data string) []string {
				(*(*func(string) []string)(signal))(data)
				return f(data)
			}
			qt.ConnectSignal(ptr.Pointer(), "sendToGo", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sendToGo", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectSendToGo() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "sendToGo")
	}
}

func (ptr *QmlBridge) SendToGo(data string) []string {
	if ptr.Pointer() != nil {
		var dataC *C.char
		if data != "" {
			dataC = C.CString(data)
			defer C.free(unsafe.Pointer(dataC))
		}
		return unpackStringList(cGoUnpackString(C.QmlBridge9911b0_SendToGo(ptr.Pointer(), C.struct_Moc_PackedString{data: dataC, len: C.longlong(len(data))})))
	}
	return make([]string, 0)
}

//export callbackQmlBridge9911b0_Find
func callbackQmlBridge9911b0_Find(ptr unsafe.Pointer, id C.struct_Moc_PackedString) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "find"); signal != nil {
		tempVal := (*(*func(string) []string)(signal))(cGoUnpackString(id))
		return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "¡¦!")), len: C.longlong(len(strings.Join(tempVal, "¡¦!")))}
	}
	tempVal := make([]string, 0)
	return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "¡¦!")), len: C.longlong(len(strings.Join(tempVal, "¡¦!")))}
}

func (ptr *QmlBridge) ConnectFind(f func(id string) []string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "find"); signal != nil {
			f := func(id string) []string {
				(*(*func(string) []string)(signal))(id)
				return f(id)
			}
			qt.ConnectSignal(ptr.Pointer(), "find", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "find", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectFind() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "find")
	}
}

func (ptr *QmlBridge) Find(id string) []string {
	if ptr.Pointer() != nil {
		var idC *C.char
		if id != "" {
			idC = C.CString(id)
			defer C.free(unsafe.Pointer(idC))
		}
		return unpackStringList(cGoUnpackString(C.QmlBridge9911b0_Find(ptr.Pointer(), C.struct_Moc_PackedString{data: idC, len: C.longlong(len(id))})))
	}
	return make([]string, 0)
}

//export callbackQmlBridge9911b0_Add
func callbackQmlBridge9911b0_Add(ptr unsafe.Pointer, obj C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "add"); signal != nil {
		(*(*func([]*std_core.QVariant))(signal))(func(l C.struct_Moc_PackedList) []*std_core.QVariant {
			out := make([]*std_core.QVariant, int(l.len))
			tmpList := NewQmlBridgeFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__add_obj_atList(i)
			}
			return out
		}(obj))
	}

}

func (ptr *QmlBridge) ConnectAdd(f func(obj []*std_core.QVariant)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "add") {
			C.QmlBridge9911b0_ConnectAdd(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "add")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "add"); signal != nil {
			f := func(obj []*std_core.QVariant) {
				(*(*func([]*std_core.QVariant))(signal))(obj)
				f(obj)
			}
			qt.ConnectSignal(ptr.Pointer(), "add", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "add", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectAdd() {
	if ptr.Pointer() != nil {
		C.QmlBridge9911b0_DisconnectAdd(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "add")
	}
}

func (ptr *QmlBridge) Add(obj []*std_core.QVariant) {
	if ptr.Pointer() != nil {
		C.QmlBridge9911b0_Add(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQmlBridgeFromPointer(NewQmlBridgeFromPointer(nil).__add_obj_newList())
			for _, v := range obj {
				tmpList.__add_obj_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackQmlBridge9911b0_Wrds
func callbackQmlBridge9911b0_Wrds(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "wrds"); signal != nil {
		tempVal := (*(*func() []string)(signal))()
		return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "¡¦!")), len: C.longlong(len(strings.Join(tempVal, "¡¦!")))}
	}
	tempVal := NewQmlBridgeFromPointer(ptr).WrdsDefault()
	return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "¡¦!")), len: C.longlong(len(strings.Join(tempVal, "¡¦!")))}
}

func (ptr *QmlBridge) ConnectWrds(f func() []string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "wrds"); signal != nil {
			f := func() []string {
				(*(*func() []string)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "wrds", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "wrds", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectWrds() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "wrds")
	}
}

func (ptr *QmlBridge) Wrds() []string {
	if ptr.Pointer() != nil {
		return unpackStringList(cGoUnpackString(C.QmlBridge9911b0_Wrds(ptr.Pointer())))
	}
	return make([]string, 0)
}

func (ptr *QmlBridge) WrdsDefault() []string {
	if ptr.Pointer() != nil {
		return unpackStringList(cGoUnpackString(C.QmlBridge9911b0_WrdsDefault(ptr.Pointer())))
	}
	return make([]string, 0)
}

//export callbackQmlBridge9911b0_SetWrds
func callbackQmlBridge9911b0_SetWrds(ptr unsafe.Pointer, wrds C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setWrds"); signal != nil {
		(*(*func([]string))(signal))(unpackStringList(cGoUnpackString(wrds)))
	} else {
		NewQmlBridgeFromPointer(ptr).SetWrdsDefault(unpackStringList(cGoUnpackString(wrds)))
	}
}

func (ptr *QmlBridge) ConnectSetWrds(f func(wrds []string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setWrds"); signal != nil {
			f := func(wrds []string) {
				(*(*func([]string))(signal))(wrds)
				f(wrds)
			}
			qt.ConnectSignal(ptr.Pointer(), "setWrds", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setWrds", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectSetWrds() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setWrds")
	}
}

func (ptr *QmlBridge) SetWrds(wrds []string) {
	if ptr.Pointer() != nil {
		wrdsC := C.CString(strings.Join(wrds, "¡¦!"))
		defer C.free(unsafe.Pointer(wrdsC))
		C.QmlBridge9911b0_SetWrds(ptr.Pointer(), C.struct_Moc_PackedString{data: wrdsC, len: C.longlong(len(strings.Join(wrds, "¡¦!")))})
	}
}

func (ptr *QmlBridge) SetWrdsDefault(wrds []string) {
	if ptr.Pointer() != nil {
		wrdsC := C.CString(strings.Join(wrds, "¡¦!"))
		defer C.free(unsafe.Pointer(wrdsC))
		C.QmlBridge9911b0_SetWrdsDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: wrdsC, len: C.longlong(len(strings.Join(wrds, "¡¦!")))})
	}
}

//export callbackQmlBridge9911b0_WrdsChanged
func callbackQmlBridge9911b0_WrdsChanged(ptr unsafe.Pointer, wrds C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "wrdsChanged"); signal != nil {
		(*(*func([]string))(signal))(unpackStringList(cGoUnpackString(wrds)))
	}

}

func (ptr *QmlBridge) ConnectWrdsChanged(f func(wrds []string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "wrdsChanged") {
			C.QmlBridge9911b0_ConnectWrdsChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "wrdsChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "wrdsChanged"); signal != nil {
			f := func(wrds []string) {
				(*(*func([]string))(signal))(wrds)
				f(wrds)
			}
			qt.ConnectSignal(ptr.Pointer(), "wrdsChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "wrdsChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectWrdsChanged() {
	if ptr.Pointer() != nil {
		C.QmlBridge9911b0_DisconnectWrdsChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "wrdsChanged")
	}
}

func (ptr *QmlBridge) WrdsChanged(wrds []string) {
	if ptr.Pointer() != nil {
		wrdsC := C.CString(strings.Join(wrds, "¡¦!"))
		defer C.free(unsafe.Pointer(wrdsC))
		C.QmlBridge9911b0_WrdsChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: wrdsC, len: C.longlong(len(strings.Join(wrds, "¡¦!")))})
	}
}

//export callbackQmlBridge9911b0_Word
func callbackQmlBridge9911b0_Word(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "word"); signal != nil {
		tempVal := (*(*func() string)(signal))()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewQmlBridgeFromPointer(ptr).WordDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QmlBridge) ConnectWord(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "word"); signal != nil {
			f := func() string {
				(*(*func() string)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "word", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "word", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectWord() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "word")
	}
}

func (ptr *QmlBridge) Word() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QmlBridge9911b0_Word(ptr.Pointer()))
	}
	return ""
}

func (ptr *QmlBridge) WordDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QmlBridge9911b0_WordDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackQmlBridge9911b0_SetWord
func callbackQmlBridge9911b0_SetWord(ptr unsafe.Pointer, word C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setWord"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(word))
	} else {
		NewQmlBridgeFromPointer(ptr).SetWordDefault(cGoUnpackString(word))
	}
}

func (ptr *QmlBridge) ConnectSetWord(f func(word string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setWord"); signal != nil {
			f := func(word string) {
				(*(*func(string))(signal))(word)
				f(word)
			}
			qt.ConnectSignal(ptr.Pointer(), "setWord", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setWord", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectSetWord() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setWord")
	}
}

func (ptr *QmlBridge) SetWord(word string) {
	if ptr.Pointer() != nil {
		var wordC *C.char
		if word != "" {
			wordC = C.CString(word)
			defer C.free(unsafe.Pointer(wordC))
		}
		C.QmlBridge9911b0_SetWord(ptr.Pointer(), C.struct_Moc_PackedString{data: wordC, len: C.longlong(len(word))})
	}
}

func (ptr *QmlBridge) SetWordDefault(word string) {
	if ptr.Pointer() != nil {
		var wordC *C.char
		if word != "" {
			wordC = C.CString(word)
			defer C.free(unsafe.Pointer(wordC))
		}
		C.QmlBridge9911b0_SetWordDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: wordC, len: C.longlong(len(word))})
	}
}

//export callbackQmlBridge9911b0_WordChanged
func callbackQmlBridge9911b0_WordChanged(ptr unsafe.Pointer, word C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "wordChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(word))
	}

}

func (ptr *QmlBridge) ConnectWordChanged(f func(word string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "wordChanged") {
			C.QmlBridge9911b0_ConnectWordChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "wordChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "wordChanged"); signal != nil {
			f := func(word string) {
				(*(*func(string))(signal))(word)
				f(word)
			}
			qt.ConnectSignal(ptr.Pointer(), "wordChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "wordChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectWordChanged() {
	if ptr.Pointer() != nil {
		C.QmlBridge9911b0_DisconnectWordChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "wordChanged")
	}
}

func (ptr *QmlBridge) WordChanged(word string) {
	if ptr.Pointer() != nil {
		var wordC *C.char
		if word != "" {
			wordC = C.CString(word)
			defer C.free(unsafe.Pointer(wordC))
		}
		C.QmlBridge9911b0_WordChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: wordC, len: C.longlong(len(word))})
	}
}

func QmlBridge_QRegisterMetaType() int {
	return int(int32(C.QmlBridge9911b0_QmlBridge9911b0_QRegisterMetaType()))
}

func (ptr *QmlBridge) QRegisterMetaType() int {
	return int(int32(C.QmlBridge9911b0_QmlBridge9911b0_QRegisterMetaType()))
}

func QmlBridge_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.QmlBridge9911b0_QmlBridge9911b0_QRegisterMetaType2(typeNameC)))
}

func (ptr *QmlBridge) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.QmlBridge9911b0_QmlBridge9911b0_QRegisterMetaType2(typeNameC)))
}

func QmlBridge_QmlRegisterType() int {
	return int(int32(C.QmlBridge9911b0_QmlBridge9911b0_QmlRegisterType()))
}

func (ptr *QmlBridge) QmlRegisterType() int {
	return int(int32(C.QmlBridge9911b0_QmlBridge9911b0_QmlRegisterType()))
}

func QmlBridge_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.QmlBridge9911b0_QmlBridge9911b0_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *QmlBridge) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.QmlBridge9911b0_QmlBridge9911b0_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *QmlBridge) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.QmlBridge9911b0___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QmlBridge) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge9911b0___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QmlBridge) __children_newList() unsafe.Pointer {
	return C.QmlBridge9911b0___children_newList(ptr.Pointer())
}

func (ptr *QmlBridge) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.QmlBridge9911b0___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QmlBridge) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge9911b0___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *QmlBridge) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QmlBridge9911b0___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QmlBridge) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.QmlBridge9911b0___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QmlBridge) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge9911b0___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QmlBridge) __findChildren_newList() unsafe.Pointer {
	return C.QmlBridge9911b0___findChildren_newList(ptr.Pointer())
}

func (ptr *QmlBridge) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.QmlBridge9911b0___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QmlBridge) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge9911b0___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QmlBridge) __findChildren_newList3() unsafe.Pointer {
	return C.QmlBridge9911b0___findChildren_newList3(ptr.Pointer())
}

func (ptr *QmlBridge) __add_obj_atList(i int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.QmlBridge9911b0___add_obj_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QmlBridge) __add_obj_setList(i std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge9911b0___add_obj_setList(ptr.Pointer(), std_core.PointerFromQVariant(i))
	}
}

func (ptr *QmlBridge) __add_obj_newList() unsafe.Pointer {
	return C.QmlBridge9911b0___add_obj_newList(ptr.Pointer())
}

func NewQmlBridge(parent std_core.QObject_ITF) *QmlBridge {
	QmlBridge_QRegisterMetaType()
	tmpValue := NewQmlBridgeFromPointer(C.QmlBridge9911b0_NewQmlBridge(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQmlBridge9911b0_DestroyQmlBridge
func callbackQmlBridge9911b0_DestroyQmlBridge(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QmlBridge"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQmlBridgeFromPointer(ptr).DestroyQmlBridgeDefault()
	}
}

func (ptr *QmlBridge) ConnectDestroyQmlBridge(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QmlBridge"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QmlBridge", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QmlBridge", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectDestroyQmlBridge() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QmlBridge")
	}
}

func (ptr *QmlBridge) DestroyQmlBridge() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QmlBridge9911b0_DestroyQmlBridge(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QmlBridge) DestroyQmlBridgeDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QmlBridge9911b0_DestroyQmlBridgeDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQmlBridge9911b0_ChildEvent
func callbackQmlBridge9911b0_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewQmlBridgeFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QmlBridge) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge9911b0_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackQmlBridge9911b0_ConnectNotify
func callbackQmlBridge9911b0_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQmlBridgeFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QmlBridge) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge9911b0_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQmlBridge9911b0_CustomEvent
func callbackQmlBridge9911b0_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewQmlBridgeFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *QmlBridge) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge9911b0_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackQmlBridge9911b0_DeleteLater
func callbackQmlBridge9911b0_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQmlBridgeFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QmlBridge) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QmlBridge9911b0_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQmlBridge9911b0_Destroyed
func callbackQmlBridge9911b0_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbackQmlBridge9911b0_DisconnectNotify
func callbackQmlBridge9911b0_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQmlBridgeFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QmlBridge) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge9911b0_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQmlBridge9911b0_Event
func callbackQmlBridge9911b0_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQmlBridgeFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *QmlBridge) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QmlBridge9911b0_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQmlBridge9911b0_EventFilter
func callbackQmlBridge9911b0_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQmlBridgeFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *QmlBridge) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QmlBridge9911b0_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQmlBridge9911b0_ObjectNameChanged
func callbackQmlBridge9911b0_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQmlBridge9911b0_TimerEvent
func callbackQmlBridge9911b0_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewQmlBridgeFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QmlBridge) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge9911b0_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

func init() {
	qt.ItfMap["main.QmlBridge_ITF"] = QmlBridge{}
	qt.FuncMap["main.NewQmlBridge"] = NewQmlBridge
}
