

#define protected public
#define private public

#include "moc.h"
#include "_cgo_export.h"

#include <QByteArray>
#include <QChildEvent>
#include <QEvent>
#include <QGraphicsObject>
#include <QGraphicsWidget>
#include <QLayout>
#include <QMetaMethod>
#include <QMetaObject>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDeviceWindow>
#include <QPdfWriter>
#include <QString>
#include <QTimerEvent>
#include <QVariant>
#include <QWidget>
#include <QWindow>

#ifdef QT_QML_LIB
	#include <QQmlEngine>
#endif


class QmlBridge9911b0: public QObject
{
Q_OBJECT
Q_PROPERTY(QStringList wrds READ wrds WRITE setWrds NOTIFY wrdsChanged)
Q_PROPERTY(QString word READ word WRITE setWord NOTIFY wordChanged)
public:
	QmlBridge9911b0(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");QmlBridge9911b0_QmlBridge9911b0_QRegisterMetaType();QmlBridge9911b0_QmlBridge9911b0_QRegisterMetaTypes();callbackQmlBridge9911b0_Constructor(this);};
	void Signal_SendToQml(QString source, QString action, QString data) { QByteArray* t828d33 = new QByteArray(source.toUtf8()); Moc_PackedString sourcePacked = { const_cast<char*>(t828d33->prepend("WHITESPACE").constData()+10), t828d33->size()-10, t828d33 };QByteArray* t34eb4c = new QByteArray(action.toUtf8()); Moc_PackedString actionPacked = { const_cast<char*>(t34eb4c->prepend("WHITESPACE").constData()+10), t34eb4c->size()-10, t34eb4c };QByteArray* ta17c9a = new QByteArray(data.toUtf8()); Moc_PackedString dataPacked = { const_cast<char*>(ta17c9a->prepend("WHITESPACE").constData()+10), ta17c9a->size()-10, ta17c9a };callbackQmlBridge9911b0_SendToQml(this, sourcePacked, actionPacked, dataPacked); };
	void Signal_Add(QList<QVariant> obj) { callbackQmlBridge9911b0_Add(this, ({ QList<QVariant>* tmpValue9b5c0b = new QList<QVariant>(obj); Moc_PackedList { tmpValue9b5c0b, tmpValue9b5c0b->size() }; })); };
	QStringList wrds() { return ({ Moc_PackedString tempVal = callbackQmlBridge9911b0_Wrds(this); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("¡¦!", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	void setWrds(QStringList wrds) { QByteArray* tee2095 = new QByteArray(wrds.join("¡¦!").toUtf8()); Moc_PackedString wrdsPacked = { const_cast<char*>(tee2095->prepend("WHITESPACE").constData()+10), tee2095->size()-10, tee2095 };callbackQmlBridge9911b0_SetWrds(this, wrdsPacked); };
	void Signal_WrdsChanged(QStringList wrds) { QByteArray* tee2095 = new QByteArray(wrds.join("¡¦!").toUtf8()); Moc_PackedString wrdsPacked = { const_cast<char*>(tee2095->prepend("WHITESPACE").constData()+10), tee2095->size()-10, tee2095 };callbackQmlBridge9911b0_WrdsChanged(this, wrdsPacked); };
	QString word() { return ({ Moc_PackedString tempVal = callbackQmlBridge9911b0_Word(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setWord(QString word) { QByteArray* t3cbcd9 = new QByteArray(word.toUtf8()); Moc_PackedString wordPacked = { const_cast<char*>(t3cbcd9->prepend("WHITESPACE").constData()+10), t3cbcd9->size()-10, t3cbcd9 };callbackQmlBridge9911b0_SetWord(this, wordPacked); };
	void Signal_WordChanged(QString word) { QByteArray* t3cbcd9 = new QByteArray(word.toUtf8()); Moc_PackedString wordPacked = { const_cast<char*>(t3cbcd9->prepend("WHITESPACE").constData()+10), t3cbcd9->size()-10, t3cbcd9 };callbackQmlBridge9911b0_WordChanged(this, wordPacked); };
	 ~QmlBridge9911b0() { callbackQmlBridge9911b0_DestroyQmlBridge(this); };
	void childEvent(QChildEvent * event) { callbackQmlBridge9911b0_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQmlBridge9911b0_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQmlBridge9911b0_CustomEvent(this, event); };
	void deleteLater() { callbackQmlBridge9911b0_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQmlBridge9911b0_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQmlBridge9911b0_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQmlBridge9911b0_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQmlBridge9911b0_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQmlBridge9911b0_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQmlBridge9911b0_TimerEvent(this, event); };
	QStringList wrdsDefault() { return _wrds; };
	void setWrdsDefault(QStringList p) { if (p != _wrds) { _wrds = p; wrdsChanged(_wrds); } };
	QString wordDefault() { return _word; };
	void setWordDefault(QString p) { if (p != _word) { _word = p; wordChanged(_word); } };
signals:
	void sendToQml(QString source, QString action, QString data);
	void add(QList<QVariant> obj);
	void wrdsChanged(QStringList wrds);
	void wordChanged(QString word);
public slots:
	QStringList sendToGo(QString data) { QByteArray* ta17c9a = new QByteArray(data.toUtf8()); Moc_PackedString dataPacked = { const_cast<char*>(ta17c9a->prepend("WHITESPACE").constData()+10), ta17c9a->size()-10, ta17c9a };return ({ Moc_PackedString tempVal = callbackQmlBridge9911b0_SendToGo(this, dataPacked); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("¡¦!", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	QStringList find(QString id) { QByteArray* t87ea5d = new QByteArray(id.toUtf8()); Moc_PackedString idPacked = { const_cast<char*>(t87ea5d->prepend("WHITESPACE").constData()+10), t87ea5d->size()-10, t87ea5d };return ({ Moc_PackedString tempVal = callbackQmlBridge9911b0_Find(this, idPacked); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("¡¦!", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
private:
	QStringList _wrds;
	QString _word;
};

Q_DECLARE_METATYPE(QmlBridge9911b0*)


void QmlBridge9911b0_QmlBridge9911b0_QRegisterMetaTypes() {
	qRegisterMetaType<QStringList>();
	qRegisterMetaType<QString>();
}

void QmlBridge9911b0_ConnectSendToQml(void* ptr, long long t)
{
	QObject::connect(static_cast<QmlBridge9911b0*>(ptr), static_cast<void (QmlBridge9911b0::*)(QString, QString, QString)>(&QmlBridge9911b0::sendToQml), static_cast<QmlBridge9911b0*>(ptr), static_cast<void (QmlBridge9911b0::*)(QString, QString, QString)>(&QmlBridge9911b0::Signal_SendToQml), static_cast<Qt::ConnectionType>(t));
}

void QmlBridge9911b0_DisconnectSendToQml(void* ptr)
{
	QObject::disconnect(static_cast<QmlBridge9911b0*>(ptr), static_cast<void (QmlBridge9911b0::*)(QString, QString, QString)>(&QmlBridge9911b0::sendToQml), static_cast<QmlBridge9911b0*>(ptr), static_cast<void (QmlBridge9911b0::*)(QString, QString, QString)>(&QmlBridge9911b0::Signal_SendToQml));
}

void QmlBridge9911b0_SendToQml(void* ptr, struct Moc_PackedString source, struct Moc_PackedString action, struct Moc_PackedString data)
{
	static_cast<QmlBridge9911b0*>(ptr)->sendToQml(QString::fromUtf8(source.data, source.len), QString::fromUtf8(action.data, action.len), QString::fromUtf8(data.data, data.len));
}

struct Moc_PackedString QmlBridge9911b0_SendToGo(void* ptr, struct Moc_PackedString data)
{
	QStringList returnArg;
	QMetaObject::invokeMethod(static_cast<QmlBridge9911b0*>(ptr), "sendToGo", Q_RETURN_ARG(QStringList, returnArg), Q_ARG(QString, QString::fromUtf8(data.data, data.len)));
	return ({ QByteArray* t8e5b69 = new QByteArray(returnArg.join("¡¦!").toUtf8()); Moc_PackedString { const_cast<char*>(t8e5b69->prepend("WHITESPACE").constData()+10), t8e5b69->size()-10, t8e5b69 }; });
}

struct Moc_PackedString QmlBridge9911b0_Find(void* ptr, struct Moc_PackedString id)
{
	QStringList returnArg;
	QMetaObject::invokeMethod(static_cast<QmlBridge9911b0*>(ptr), "find", Q_RETURN_ARG(QStringList, returnArg), Q_ARG(QString, QString::fromUtf8(id.data, id.len)));
	return ({ QByteArray* t8e5b69 = new QByteArray(returnArg.join("¡¦!").toUtf8()); Moc_PackedString { const_cast<char*>(t8e5b69->prepend("WHITESPACE").constData()+10), t8e5b69->size()-10, t8e5b69 }; });
}

void QmlBridge9911b0_ConnectAdd(void* ptr, long long t)
{
	QObject::connect(static_cast<QmlBridge9911b0*>(ptr), static_cast<void (QmlBridge9911b0::*)(QList<QVariant>)>(&QmlBridge9911b0::add), static_cast<QmlBridge9911b0*>(ptr), static_cast<void (QmlBridge9911b0::*)(QList<QVariant>)>(&QmlBridge9911b0::Signal_Add), static_cast<Qt::ConnectionType>(t));
}

void QmlBridge9911b0_DisconnectAdd(void* ptr)
{
	QObject::disconnect(static_cast<QmlBridge9911b0*>(ptr), static_cast<void (QmlBridge9911b0::*)(QList<QVariant>)>(&QmlBridge9911b0::add), static_cast<QmlBridge9911b0*>(ptr), static_cast<void (QmlBridge9911b0::*)(QList<QVariant>)>(&QmlBridge9911b0::Signal_Add));
}

void QmlBridge9911b0_Add(void* ptr, void* obj)
{
	static_cast<QmlBridge9911b0*>(ptr)->add(({ QList<QVariant>* tmpP = static_cast<QList<QVariant>*>(obj); QList<QVariant> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

struct Moc_PackedString QmlBridge9911b0_Wrds(void* ptr)
{
	return ({ QByteArray* td5d06f = new QByteArray(static_cast<QmlBridge9911b0*>(ptr)->wrds().join("¡¦!").toUtf8()); Moc_PackedString { const_cast<char*>(td5d06f->prepend("WHITESPACE").constData()+10), td5d06f->size()-10, td5d06f }; });
}

struct Moc_PackedString QmlBridge9911b0_WrdsDefault(void* ptr)
{
	return ({ QByteArray* t114fe5 = new QByteArray(static_cast<QmlBridge9911b0*>(ptr)->wrdsDefault().join("¡¦!").toUtf8()); Moc_PackedString { const_cast<char*>(t114fe5->prepend("WHITESPACE").constData()+10), t114fe5->size()-10, t114fe5 }; });
}

void QmlBridge9911b0_SetWrds(void* ptr, struct Moc_PackedString wrds)
{
	static_cast<QmlBridge9911b0*>(ptr)->setWrds(QString::fromUtf8(wrds.data, wrds.len).split("¡¦!", QString::SkipEmptyParts));
}

void QmlBridge9911b0_SetWrdsDefault(void* ptr, struct Moc_PackedString wrds)
{
	static_cast<QmlBridge9911b0*>(ptr)->setWrdsDefault(QString::fromUtf8(wrds.data, wrds.len).split("¡¦!", QString::SkipEmptyParts));
}

void QmlBridge9911b0_ConnectWrdsChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QmlBridge9911b0*>(ptr), static_cast<void (QmlBridge9911b0::*)(QStringList)>(&QmlBridge9911b0::wrdsChanged), static_cast<QmlBridge9911b0*>(ptr), static_cast<void (QmlBridge9911b0::*)(QStringList)>(&QmlBridge9911b0::Signal_WrdsChanged), static_cast<Qt::ConnectionType>(t));
}

void QmlBridge9911b0_DisconnectWrdsChanged(void* ptr)
{
	QObject::disconnect(static_cast<QmlBridge9911b0*>(ptr), static_cast<void (QmlBridge9911b0::*)(QStringList)>(&QmlBridge9911b0::wrdsChanged), static_cast<QmlBridge9911b0*>(ptr), static_cast<void (QmlBridge9911b0::*)(QStringList)>(&QmlBridge9911b0::Signal_WrdsChanged));
}

void QmlBridge9911b0_WrdsChanged(void* ptr, struct Moc_PackedString wrds)
{
	static_cast<QmlBridge9911b0*>(ptr)->wrdsChanged(QString::fromUtf8(wrds.data, wrds.len).split("¡¦!", QString::SkipEmptyParts));
}

struct Moc_PackedString QmlBridge9911b0_Word(void* ptr)
{
	return ({ QByteArray* t6415ee = new QByteArray(static_cast<QmlBridge9911b0*>(ptr)->word().toUtf8()); Moc_PackedString { const_cast<char*>(t6415ee->prepend("WHITESPACE").constData()+10), t6415ee->size()-10, t6415ee }; });
}

struct Moc_PackedString QmlBridge9911b0_WordDefault(void* ptr)
{
	return ({ QByteArray* t3a80f1 = new QByteArray(static_cast<QmlBridge9911b0*>(ptr)->wordDefault().toUtf8()); Moc_PackedString { const_cast<char*>(t3a80f1->prepend("WHITESPACE").constData()+10), t3a80f1->size()-10, t3a80f1 }; });
}

void QmlBridge9911b0_SetWord(void* ptr, struct Moc_PackedString word)
{
	static_cast<QmlBridge9911b0*>(ptr)->setWord(QString::fromUtf8(word.data, word.len));
}

void QmlBridge9911b0_SetWordDefault(void* ptr, struct Moc_PackedString word)
{
	static_cast<QmlBridge9911b0*>(ptr)->setWordDefault(QString::fromUtf8(word.data, word.len));
}

void QmlBridge9911b0_ConnectWordChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QmlBridge9911b0*>(ptr), static_cast<void (QmlBridge9911b0::*)(QString)>(&QmlBridge9911b0::wordChanged), static_cast<QmlBridge9911b0*>(ptr), static_cast<void (QmlBridge9911b0::*)(QString)>(&QmlBridge9911b0::Signal_WordChanged), static_cast<Qt::ConnectionType>(t));
}

void QmlBridge9911b0_DisconnectWordChanged(void* ptr)
{
	QObject::disconnect(static_cast<QmlBridge9911b0*>(ptr), static_cast<void (QmlBridge9911b0::*)(QString)>(&QmlBridge9911b0::wordChanged), static_cast<QmlBridge9911b0*>(ptr), static_cast<void (QmlBridge9911b0::*)(QString)>(&QmlBridge9911b0::Signal_WordChanged));
}

void QmlBridge9911b0_WordChanged(void* ptr, struct Moc_PackedString word)
{
	static_cast<QmlBridge9911b0*>(ptr)->wordChanged(QString::fromUtf8(word.data, word.len));
}

int QmlBridge9911b0_QmlBridge9911b0_QRegisterMetaType()
{
	return qRegisterMetaType<QmlBridge9911b0*>();
}

int QmlBridge9911b0_QmlBridge9911b0_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<QmlBridge9911b0*>(const_cast<const char*>(typeName));
}

int QmlBridge9911b0_QmlBridge9911b0_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<QmlBridge9911b0>();
#else
	return 0;
#endif
}

int QmlBridge9911b0_QmlBridge9911b0_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<QmlBridge9911b0>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* QmlBridge9911b0___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QmlBridge9911b0___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QmlBridge9911b0___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QmlBridge9911b0___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QmlBridge9911b0___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QmlBridge9911b0___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QmlBridge9911b0___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QmlBridge9911b0___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QmlBridge9911b0___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QmlBridge9911b0___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QmlBridge9911b0___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QmlBridge9911b0___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QmlBridge9911b0___add_obj_atList(void* ptr, int i)
{
	return new QVariant(({QVariant tmp = static_cast<QList<QVariant>*>(ptr)->at(i); if (i == static_cast<QList<QVariant>*>(ptr)->size()-1) { static_cast<QList<QVariant>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QmlBridge9911b0___add_obj_setList(void* ptr, void* i)
{
	static_cast<QList<QVariant>*>(ptr)->append(*static_cast<QVariant*>(i));
}

void* QmlBridge9911b0___add_obj_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QVariant>();
}

void* QmlBridge9911b0_NewQmlBridge(void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new QmlBridge9911b0(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new QmlBridge9911b0(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new QmlBridge9911b0(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new QmlBridge9911b0(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new QmlBridge9911b0(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new QmlBridge9911b0(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new QmlBridge9911b0(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new QmlBridge9911b0(static_cast<QWindow*>(parent));
	} else {
		return new QmlBridge9911b0(static_cast<QObject*>(parent));
	}
}

void QmlBridge9911b0_DestroyQmlBridge(void* ptr)
{
	static_cast<QmlBridge9911b0*>(ptr)->~QmlBridge9911b0();
}

void QmlBridge9911b0_DestroyQmlBridgeDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void QmlBridge9911b0_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QmlBridge9911b0*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void QmlBridge9911b0_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QmlBridge9911b0*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QmlBridge9911b0_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QmlBridge9911b0*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void QmlBridge9911b0_DeleteLaterDefault(void* ptr)
{
	static_cast<QmlBridge9911b0*>(ptr)->QObject::deleteLater();
}

void QmlBridge9911b0_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QmlBridge9911b0*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QmlBridge9911b0_EventDefault(void* ptr, void* e)
{
	return static_cast<QmlBridge9911b0*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char QmlBridge9911b0_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QmlBridge9911b0*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}



void QmlBridge9911b0_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QmlBridge9911b0*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}

#include "moc_moc.h"
