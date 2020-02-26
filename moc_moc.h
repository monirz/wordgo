/****************************************************************************
** Meta object code from reading C++ file 'moc.cpp'
**
** Created by: The Qt Meta Object Compiler version 67 (Qt 5.13.2)
**
** WARNING! All changes made in this file will be lost!
*****************************************************************************/

#include <memory>
#include <QtCore/qbytearray.h>
#include <QtCore/qmetatype.h>
#include <QtCore/QList>
#if !defined(Q_MOC_OUTPUT_REVISION)
#error "The header file 'moc.cpp' doesn't include <QObject>."
#elif Q_MOC_OUTPUT_REVISION != 67
#error "This file was generated using the moc from 5.13.2. It"
#error "cannot be used with the include files from this version of Qt."
#error "(The moc has changed too much.)"
#endif

QT_BEGIN_MOC_NAMESPACE
QT_WARNING_PUSH
QT_WARNING_DISABLE_DEPRECATED
struct qt_meta_stringdata_QmlBridge9911b0_t {
    QByteArrayData data[15];
    char stringdata0[105];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_QmlBridge9911b0_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_QmlBridge9911b0_t qt_meta_stringdata_QmlBridge9911b0 = {
    {
QT_MOC_LITERAL(0, 0, 15), // "QmlBridge9911b0"
QT_MOC_LITERAL(1, 16, 9), // "sendToQml"
QT_MOC_LITERAL(2, 26, 0), // ""
QT_MOC_LITERAL(3, 27, 6), // "source"
QT_MOC_LITERAL(4, 34, 6), // "action"
QT_MOC_LITERAL(5, 41, 4), // "data"
QT_MOC_LITERAL(6, 46, 3), // "add"
QT_MOC_LITERAL(7, 50, 3), // "obj"
QT_MOC_LITERAL(8, 54, 11), // "wrdsChanged"
QT_MOC_LITERAL(9, 66, 4), // "wrds"
QT_MOC_LITERAL(10, 71, 11), // "wordChanged"
QT_MOC_LITERAL(11, 83, 4), // "word"
QT_MOC_LITERAL(12, 88, 8), // "sendToGo"
QT_MOC_LITERAL(13, 97, 4), // "find"
QT_MOC_LITERAL(14, 102, 2) // "id"

    },
    "QmlBridge9911b0\0sendToQml\0\0source\0"
    "action\0data\0add\0obj\0wrdsChanged\0wrds\0"
    "wordChanged\0word\0sendToGo\0find\0id"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_QmlBridge9911b0[] = {

 // content:
       8,       // revision
       0,       // classname
       0,    0, // classinfo
       6,   14, // methods
       2,   66, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       4,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    3,   44,    2, 0x06 /* Public */,
       6,    1,   51,    2, 0x06 /* Public */,
       8,    1,   54,    2, 0x06 /* Public */,
      10,    1,   57,    2, 0x06 /* Public */,

 // slots: name, argc, parameters, tag, flags
      12,    1,   60,    2, 0x0a /* Public */,
      13,    1,   63,    2, 0x0a /* Public */,

 // signals: parameters
    QMetaType::Void, QMetaType::QString, QMetaType::QString, QMetaType::QString,    3,    4,    5,
    QMetaType::Void, QMetaType::QVariantList,    7,
    QMetaType::Void, QMetaType::QStringList,    9,
    QMetaType::Void, QMetaType::QString,   11,

 // slots: parameters
    QMetaType::QStringList, QMetaType::QString,    5,
    QMetaType::QStringList, QMetaType::QString,   14,

 // properties: name, type, flags
       9, QMetaType::QStringList, 0x00495103,
      11, QMetaType::QString, 0x00495103,

 // properties: notify_signal_id
       2,
       3,

       0        // eod
};

void QmlBridge9911b0::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        auto *_t = static_cast<QmlBridge9911b0 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->sendToQml((*reinterpret_cast< QString(*)>(_a[1])),(*reinterpret_cast< QString(*)>(_a[2])),(*reinterpret_cast< QString(*)>(_a[3]))); break;
        case 1: _t->add((*reinterpret_cast< QList<QVariant>(*)>(_a[1]))); break;
        case 2: _t->wrdsChanged((*reinterpret_cast< QStringList(*)>(_a[1]))); break;
        case 3: _t->wordChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 4: { QStringList _r = _t->sendToGo((*reinterpret_cast< QString(*)>(_a[1])));
            if (_a[0]) *reinterpret_cast< QStringList*>(_a[0]) = std::move(_r); }  break;
        case 5: { QStringList _r = _t->find((*reinterpret_cast< QString(*)>(_a[1])));
            if (_a[0]) *reinterpret_cast< QStringList*>(_a[0]) = std::move(_r); }  break;
        default: ;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            using _t = void (QmlBridge9911b0::*)(QString , QString , QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&QmlBridge9911b0::sendToQml)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (QmlBridge9911b0::*)(QList<QVariant> );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&QmlBridge9911b0::add)) {
                *result = 1;
                return;
            }
        }
        {
            using _t = void (QmlBridge9911b0::*)(QStringList );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&QmlBridge9911b0::wrdsChanged)) {
                *result = 2;
                return;
            }
        }
        {
            using _t = void (QmlBridge9911b0::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&QmlBridge9911b0::wordChanged)) {
                *result = 3;
                return;
            }
        }
    }
#ifndef QT_NO_PROPERTIES
    else if (_c == QMetaObject::ReadProperty) {
        auto *_t = static_cast<QmlBridge9911b0 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: *reinterpret_cast< QStringList*>(_v) = _t->wrds(); break;
        case 1: *reinterpret_cast< QString*>(_v) = _t->word(); break;
        default: break;
        }
    } else if (_c == QMetaObject::WriteProperty) {
        auto *_t = static_cast<QmlBridge9911b0 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: _t->setWrds(*reinterpret_cast< QStringList*>(_v)); break;
        case 1: _t->setWord(*reinterpret_cast< QString*>(_v)); break;
        default: break;
        }
    } else if (_c == QMetaObject::ResetProperty) {
    }
#endif // QT_NO_PROPERTIES
}

QT_INIT_METAOBJECT const QMetaObject QmlBridge9911b0::staticMetaObject = { {
    &QObject::staticMetaObject,
    qt_meta_stringdata_QmlBridge9911b0.data,
    qt_meta_data_QmlBridge9911b0,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *QmlBridge9911b0::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *QmlBridge9911b0::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_QmlBridge9911b0.stringdata0))
        return static_cast<void*>(this);
    return QObject::qt_metacast(_clname);
}

int QmlBridge9911b0::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QObject::qt_metacall(_c, _id, _a);
    if (_id < 0)
        return _id;
    if (_c == QMetaObject::InvokeMetaMethod) {
        if (_id < 6)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 6;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 6)
            *reinterpret_cast<int*>(_a[0]) = -1;
        _id -= 6;
    }
#ifndef QT_NO_PROPERTIES
    else if (_c == QMetaObject::ReadProperty || _c == QMetaObject::WriteProperty
            || _c == QMetaObject::ResetProperty || _c == QMetaObject::RegisterPropertyMetaType) {
        qt_static_metacall(this, _c, _id, _a);
        _id -= 2;
    } else if (_c == QMetaObject::QueryPropertyDesignable) {
        _id -= 2;
    } else if (_c == QMetaObject::QueryPropertyScriptable) {
        _id -= 2;
    } else if (_c == QMetaObject::QueryPropertyStored) {
        _id -= 2;
    } else if (_c == QMetaObject::QueryPropertyEditable) {
        _id -= 2;
    } else if (_c == QMetaObject::QueryPropertyUser) {
        _id -= 2;
    }
#endif // QT_NO_PROPERTIES
    return _id;
}

// SIGNAL 0
void QmlBridge9911b0::sendToQml(QString _t1, QString _t2, QString _t3)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))), const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t2))), const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t3))) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}

// SIGNAL 1
void QmlBridge9911b0::add(QList<QVariant> _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 1, _a);
}

// SIGNAL 2
void QmlBridge9911b0::wrdsChanged(QStringList _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 2, _a);
}

// SIGNAL 3
void QmlBridge9911b0::wordChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 3, _a);
}
QT_WARNING_POP
QT_END_MOC_NAMESPACE
