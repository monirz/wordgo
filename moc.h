

#pragma once

#ifndef GO_MOC_9911b0_H
#define GO_MOC_9911b0_H

#include <stdint.h>

#ifdef __cplusplus
class QmlBridge9911b0;
void QmlBridge9911b0_QmlBridge9911b0_QRegisterMetaTypes();
extern "C" {
#endif

struct Moc_PackedString { char* data; long long len; void* ptr; };
struct Moc_PackedList { void* data; long long len; };
void QmlBridge9911b0_ConnectSendToQml(void* ptr, long long t);
void QmlBridge9911b0_DisconnectSendToQml(void* ptr);
void QmlBridge9911b0_SendToQml(void* ptr, struct Moc_PackedString source, struct Moc_PackedString action, struct Moc_PackedString data);
struct Moc_PackedString QmlBridge9911b0_SendToGo(void* ptr, struct Moc_PackedString data);
struct Moc_PackedString QmlBridge9911b0_Find(void* ptr, struct Moc_PackedString id);
void QmlBridge9911b0_ConnectAdd(void* ptr, long long t);
void QmlBridge9911b0_DisconnectAdd(void* ptr);
void QmlBridge9911b0_Add(void* ptr, void* obj);
struct Moc_PackedString QmlBridge9911b0_Wrds(void* ptr);
struct Moc_PackedString QmlBridge9911b0_WrdsDefault(void* ptr);
void QmlBridge9911b0_SetWrds(void* ptr, struct Moc_PackedString wrds);
void QmlBridge9911b0_SetWrdsDefault(void* ptr, struct Moc_PackedString wrds);
void QmlBridge9911b0_ConnectWrdsChanged(void* ptr, long long t);
void QmlBridge9911b0_DisconnectWrdsChanged(void* ptr);
void QmlBridge9911b0_WrdsChanged(void* ptr, struct Moc_PackedString wrds);
struct Moc_PackedString QmlBridge9911b0_Word(void* ptr);
struct Moc_PackedString QmlBridge9911b0_WordDefault(void* ptr);
void QmlBridge9911b0_SetWord(void* ptr, struct Moc_PackedString word);
void QmlBridge9911b0_SetWordDefault(void* ptr, struct Moc_PackedString word);
void QmlBridge9911b0_ConnectWordChanged(void* ptr, long long t);
void QmlBridge9911b0_DisconnectWordChanged(void* ptr);
void QmlBridge9911b0_WordChanged(void* ptr, struct Moc_PackedString word);
int QmlBridge9911b0_QmlBridge9911b0_QRegisterMetaType();
int QmlBridge9911b0_QmlBridge9911b0_QRegisterMetaType2(char* typeName);
int QmlBridge9911b0_QmlBridge9911b0_QmlRegisterType();
int QmlBridge9911b0_QmlBridge9911b0_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName);
void* QmlBridge9911b0___children_atList(void* ptr, int i);
void QmlBridge9911b0___children_setList(void* ptr, void* i);
void* QmlBridge9911b0___children_newList(void* ptr);
void* QmlBridge9911b0___dynamicPropertyNames_atList(void* ptr, int i);
void QmlBridge9911b0___dynamicPropertyNames_setList(void* ptr, void* i);
void* QmlBridge9911b0___dynamicPropertyNames_newList(void* ptr);
void* QmlBridge9911b0___findChildren_atList(void* ptr, int i);
void QmlBridge9911b0___findChildren_setList(void* ptr, void* i);
void* QmlBridge9911b0___findChildren_newList(void* ptr);
void* QmlBridge9911b0___findChildren_atList3(void* ptr, int i);
void QmlBridge9911b0___findChildren_setList3(void* ptr, void* i);
void* QmlBridge9911b0___findChildren_newList3(void* ptr);
void* QmlBridge9911b0___add_obj_atList(void* ptr, int i);
void QmlBridge9911b0___add_obj_setList(void* ptr, void* i);
void* QmlBridge9911b0___add_obj_newList(void* ptr);
void* QmlBridge9911b0_NewQmlBridge(void* parent);
void QmlBridge9911b0_DestroyQmlBridge(void* ptr);
void QmlBridge9911b0_DestroyQmlBridgeDefault(void* ptr);
void QmlBridge9911b0_ChildEventDefault(void* ptr, void* event);
void QmlBridge9911b0_ConnectNotifyDefault(void* ptr, void* sign);
void QmlBridge9911b0_CustomEventDefault(void* ptr, void* event);
void QmlBridge9911b0_DeleteLaterDefault(void* ptr);
void QmlBridge9911b0_DisconnectNotifyDefault(void* ptr, void* sign);
char QmlBridge9911b0_EventDefault(void* ptr, void* e);
char QmlBridge9911b0_EventFilterDefault(void* ptr, void* watched, void* event);
;
void QmlBridge9911b0_TimerEventDefault(void* ptr, void* event);

#ifdef __cplusplus
}
#endif

#endif