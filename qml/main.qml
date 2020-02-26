import QtQuick 2.6
import QtQuick.Controls 2.0

ApplicationWindow {
    width: 1024
    height: 600
    visible: true
    color: "#EFEBE7"
    id: root

     property variant items: ["four", "five", "foo", "bar", "baz"]
    property variant items2: []
    property variant resultItems: ["earth", "mars", "saturn"]
    property bool ok: false

   Connections {
       target: qmlBridge
   }

    Column{

        x: 10

        width: parent.width
        //        height: 420

        Rectangle{
            id: textRec
            height: 60
            color: "#efebe7"
            width: parent.width
            //            border.width: 1
            //            border.color: "#B7B4B1"


            TextField {
                id: textField
                x: 0
                y: 10

                leftPadding: 30
                topPadding: 5
                width: parent.width * 0.5
                font.family: "Arial"
                //                padding: 6
                renderType: Text.QtRendering
                horizontalAlignment: Text.AlignLeft
                focus: true
                selectByMouse: true
                //                anchors.margins: 10
                placeholderText: "Type a word..."
                rightPadding: 30

                Image {
                    id: searchIcon
                    anchors {left: parent.left; margins: 1 }
                    smooth: true
                    fillMode: Image.PreserveAspectFit
                    source: "search.png"
                    height: parent.height - 2
                     width: parent.height - 2

                }


                background: Rectangle {
                    color: "#FEFDEB"
                    border.color:   textField.activeFocus ? textField.selectionStart == 0 && textField.selectionEnd == 0 ? "#000000" : "#FF7855" : "transparent"
                    radius: 3
                }

                onTextChanged: {
                    // console.log("text changed")
                    listView.model = qmlBridge.sendToGo(text)
                    popup.open()
                }

                Keys.forwardTo: [listView.currentItem, listView]

                Popup {
                    id: popup
                    padding: 1
                    y: parent.height
                    width: parent.width
                    contentHeight: listView.contentHeight

                    background: Rectangle {
                        implicitWidth: 200
                        implicitHeight: 200
                        border.color: "#444"
                        color: "#EFEBE7"
                    }

                    ListView {
                        id: listView
                        currentIndex: 0
                        anchors.fill: parent
                        model: items2

                        delegate: ItemDelegate {
                            id: control
                            text: modelData
                            font.family: "Arial"
                            width: parent.width
                            highlighted: index === listView.currentIndex

                            background: Rectangle {
                                color: control.highlighted ? "#FA774E" : "#eeeeee"
                            }
                        }

                        Keys.onReturnPressed: {
                            textField.text = listView.model[listView.currentIndex]
                            rezsultList.model = qmlBridge.find(listView.model[listView.currentIndex])
                            wordArea.text =  listView.model[listView.currentIndex]
                            popup.close()
                        }


                    }
                }
            }
        }


        Row{
            spacing: 10
            height: 500

            Rectangle{
                color: "#FEFDEB"
                border.width: 1
                border.color: "#B7B4B1"
                width: 360
                height: parent.height
                Text {
                    id: wordArea
                    font.family: "Arial"
                    x: 10
                    y: 5
                    text: ""
                }

            }

            Rectangle{
                id: wordMeaning
                color: "#FEFDEB"
                border.width: 1
                ListView{
                    id: rezsultList
                    width: 180; height: 200


                    topMargin: 10
                    leftMargin: 10
                    model: []
                    delegate: Text {
                        text: modelData

                        bottomPadding: 20
                    }
                }

                border.color: "#B7B4B1"
                width: root.width * 0.6
                height: parent.height

                Text{
                    id: resultRec
                    font.family: "Arial"
                    leftPadding: 10
                    topPadding: 5
                }

                MouseArea {
                    anchors.fill: parent
                    propagateComposedEvents: true
                    onPressed: {
                        wordMeaning.focus = true                    }
                }
            }
        }
    }
}
