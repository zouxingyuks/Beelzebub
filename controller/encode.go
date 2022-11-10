package controller

import (
	"Beelzebub/algorithm"
	"Beelzebub/diy"
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"log"
	"os"
	"unicode"
)

type encode struct {
	output  string //输出部分
	button  diy.Button
	add     *widget.Button
	save    *widget.Button
	load    *widget.Button
	windows fyne.Window
	table   *widget.Table
	uiW     float32
	uiH     float32
	File    *dialog.FileDialog
	key     binding.String
	content binding.String
}

var (
	data [26][2]string
)

func newEncode() (e *encode) {
	e = &encode{
		output:  "",
		windows: nil,
		uiW:     800,
		uiH:     400,
	}
	e.key = binding.NewString()
	e.content = binding.NewString()
	e.button.New(20)
	algorithm.Ruler = make(map[string]string, 26) //暂定最多能有100条规则
	return e
}

func (e *encode) addEntry(text binding.String, Holder string) *widget.Entry {
	entry := widget.NewEntryWithData(text)
	entry.PlaceHolder = Holder
	return entry
}
func (e *encode) refresh(key string, content string) {
	algorithm.Ruler[key] = content
	data[key[0]-'A'][0] = key
	data[key[0]-'A'][1] = content
	log.Println("add ruler " + key + "->" + content)
	//清空输入框
	_ = e.key.Set("")
	_ = e.content.Set("")
	e.table.Refresh()

}
func (e *encode) Entry() (c *fyne.Container) {
	c = container.NewGridWithColumns(2,

		e.addEntry(e.key, "请输入单个大写字母"),
		e.addEntry(e.content, "请输入一串小写字母"),
	)
	return c
}
func (e *encode) SaveRuler() *widget.Button {

	e.save = widget.NewButton("保存", func() {
		e.File = dialog.NewFileSave(func(u fyne.URIWriteCloser, err error) {
			r, _ := json.Marshal(algorithm.Ruler)
			_, err = u.Write(r)
			_ = u.Close()
		}, e.windows)
		e.File.SetFileName("secret.json")
		e.File.Show()
	})
	return e.save

}
func (e *encode) LoadRuler() *widget.Button {

	e.load = widget.NewButton("加载", func() {
		e.File = dialog.NewFileOpen(func(u fyne.URIReadCloser, err error) {
			fileReader, err := os.Open(u.URI().Path())
			if err != nil {
				fmt.Println("读取文件失败")
			}
			defer func(fileReader *os.File) {
				_ = fileReader.Close()
			}(fileReader)
			_ = json.NewDecoder(fileReader).Decode(&algorithm.Ruler)
			for key, value := range algorithm.Ruler {
				data[key[0]-'A'][0] = key
				data[key[0]-'A'][1] = value
			}
			e.table.Refresh()
		}, e.windows)
		e.File.Show()
	})

	return e.load
}

// 按钮-添加规则
func (e *encode) addButton() *widget.Button {
	e.add = widget.NewButton("添加", func() {
		key, _ := e.key.Get()
		content, _ := e.content.Get()
		if key == "" || content == "" {
			diy.EmptyWarning(e.windows)
			return
		}
		//key 检测
		if len(key) > 1 || !unicode.IsUpper(rune(key[0])) {
			dialog.ShowCustom("警告", "懂了懂了", container.NewGridWithRows(1, widget.NewLabel("单字符匹配模式下 key 仅可以输入单个大写字母")), e.windows)
			return
		}
		//content 检测
		for _, v := range content {
			if !unicode.IsLetter(v) {
				dialog.ShowCustom("警告", "懂了懂了", container.NewGridWithRows(1, widget.NewLabel("单字符匹配模式下 content 仅可以输入字母")), e.windows)
				return
			}
		}

		if algorithm.Ruler[key] != "" {
			dialog.ShowCustomConfirm("确认", "是的", "不是", container.NewGridWithRows(1, widget.NewLabel("此key已经存在对应规则，是否进行替换")), func(b bool) {
				if b {
					e.refresh(key, content)
				}
			}, e.windows)
			return
		}
		e.refresh(key, content)
	})
	e.add.Resize(fyne.NewSize(20, 10))
	return e.add
}

func (e *encode) rulersTable() *widget.Table {
	e.table = widget.NewTable(
		func() (int, int) { return len(data), 2 },
		func() fyne.CanvasObject {
			return widget.NewLabel("01234567890123456789")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			if data[i.Row][i.Col] != "" {
				o.(*widget.Label).SetText(data[i.Row][i.Col])
			} else {
				o.(*widget.Label).SetText("")
			}
		})
	return e.table
}
func (e *encode) loadUI(app fyne.App) interface{} {
	e.windows = app.NewWindow("开始设置加密规则")
	e.windows.Resize(fyne.NewSize(e.uiW, e.uiH))
	e.windows.SetContent(
		container.NewGridWithRows(2,
			container.NewGridWithColumns(2,
				container.New(layout.NewFormLayout(),
					//保存文件
					e.SaveRuler(),
					//加载文件
					e.addButton(),
					e.LoadRuler(),
					widget.NewLabel("删除"),
				),
				//widget.NewLabel("请勿设置自相矛盾的规则"),
				e.Entry(),
			),
			e.rulersTable(),
		),
	)
	e.windows.Show()
	return e
}
