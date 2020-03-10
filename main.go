package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"runtime"
	"strings"

	"github.com/monirz/gotri"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/quickcontrols2"
	"github.com/therecipe/qt/widgets"
	"os"
)

var (
	qmlObjects = make(map[string]*core.QObject)

	qmlBridge          *QmlBridge
	manipulatedFromQml *widgets.QWidget
)

var t *gotri.Trie

func main() {

	var (
		f   *os.File
		err error
	)

	if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {

		filePath := "/usr/share/dictionary.txt"

		_, err := os.Stat(filePath)

		if err != nil {
			err = errors.New("Error lookign file path")
			log.Fatal(err)
		}

		if os.IsNotExist(err) {
			filePath = "dictionary.txt"
		}

		f, err = os.Open(filePath)
		if err != nil {
			err = errors.New("dictionary.txt file is not found in /usr/share or in the current directory")
			log.Fatal(err)
		}

	} else {
		f, err = os.Open("dictionary.txt")
		if err != nil {
			err = errors.New("dictionary.txt file is not found in the current directory")
			log.Fatal(err)
		}

	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	count := 0
	t = gotri.New()

	for scanner.Scan() {

		splitted := strings.Split(scanner.Text(), "|")
		count++

		// if count == 5000 {
		// 	break
		// }

		t.Add(splitted[0], splitted[1])
	}

	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	gui.NewQGuiApplication(len(os.Args), os.Args)
	quickcontrols2.QQuickStyle_SetStyle("Material")

	bridge := initQmlBridge()

	// create the qml application engine
	engine := qml.NewQQmlApplicationEngine(nil)
	engine.RootContext().SetContextProperty("qmlBridge", bridge)
	engine.Load(core.NewQUrl3("qrc:/qml/main.qml", 0))

	gui.QGuiApplication_Exec()
}

type QmlBridge struct {
	core.QObject
	core.QAbstractListModel

	// _ func() `constructor:"init"`

	Wr []string                          `property:"wrds"`
	_  string                            `property:"word"`
	_  func(source, action, data string) `signal:"sendToQml"`
	_  func(data string) []string        `slot:"sendToGo"`
	_  func(id string) []string          `slot:"find"`

	_ func(obj []*core.QVariant) `signal:"add,auto"`
}

func (wm *QmlBridge) find(key string) string {

	if len(qmlBridge.Wr) > 0 {
		return qmlBridge.Wr[0]
	}

	return "No value"

}

func initQmlBridge() *QmlBridge {

	qmlBridge = NewQmlBridge(nil)

	qmlBridge.ConnectSendToGo(func(data string) []string {

		wordList := t.GetSuggestion(data, 10)

		qmlBridge.Wr = wordList
		qmlBridge.SetWrds(wordList)
		// qmlBridge.SetWord(data)
		return wordList

	})

	qmlBridge.ConnectRowCount(qmlBridge.rowCount)
	qmlBridge.ConnectData(qmlBridge.data)

	qmlBridge.ConnectFind(func(data string) []string {

		word, ok := t.Search(data)

		wordList := strings.Split(word, ",")

		if !ok {
			return wordList
		}
		return wordList

	})

	return qmlBridge
}

func (qm *QmlBridge) init() {

}

func (qm *QmlBridge) rowCount(*core.QModelIndex) int {
	return len(qm.Wr)
}

func (qm *QmlBridge) data(index *core.QModelIndex, role int) *core.QVariant {
	if role != int(core.Qt__DisplayRole) {
		return core.NewQVariant()
	}

	item := qm.Wr[index.Row()]
	return core.NewQVariant1(fmt.Sprintf("%v", item))
}

//---------Add
func (qm *QmlBridge) add(item []*core.QVariant) {
	// qm.BeginInsertRows(core.NewQModelIndex(), len(qm.modelData), len(qm.modelData))
	// qm.modelData = append(qm.modelData, ListItem{item[0].ToString(), item[1].ToString()})
	qm.EndInsertRows()
}
