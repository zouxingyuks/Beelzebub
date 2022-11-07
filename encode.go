package main

import (
	"Beelzebub/diy"
	"encoding/json"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"log"
	"unicode"
)

type encode struct {
	output  string //输出部分
	button  diy.Button
	windows fyne.Window
	table   *widget.Table
	uiW     float32
	uiH     float32
	file    *dialog.FileDialog
	key     binding.String
	content binding.String
}

var data = [][]string{[]string{"top left", "top right"}, []string{"bottom left", "bottom right"}}

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
	ruler = make(map[string]string, 26) //暂定最多能有100条规则
	e.table = widget.NewTable(func() (int, int) {
		return len(data), 2
	},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
		})
	return e
}

func (e *encode) addEntry(text binding.String, Holder string) *widget.Entry {
	entry := widget.NewEntryWithData(text)
	entry.PlaceHolder = Holder
	return entry
}
func (e *encode) refreshEntry(key string, content string) {
	ruler[key] = content
	log.Println("add ruler " + key + "->" + content)
	//清空输入框
	_ = e.key.Set("")
	_ = e.content.Set("")
}
func (e *encode) Entry() (c *fyne.Container) {
	c = container.NewGridWithColumns(4,
		widget.NewLabel("单字符匹配模式"),
		e.button.AddButton(sAdd, func() {
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

			if ruler[key] != "" {
				dialog.ShowCustomConfirm("确认", "是的", "不是", container.NewGridWithRows(1, widget.NewLabel("此key已经存在对应规则，是否进行替换")), func(b bool) {
					if b {
						e.refreshEntry(key, content)
					}
				}, e.windows)
				return
			}
			e.refreshEntry(key, content)

		}),
		e.addEntry(e.key, "请输入单个大写字母"),
		e.addEntry(e.content, "请输入一串小写字母"),
	)

	//输入框规模
	c.Resize(fyne.NewSize(40, 10))
	return c
}
func (e *encode) refreshTable() {
	e.table.CreateCell()
}
func (e *encode) loadUI(app fyne.App) interface{} {
	e.windows = app.NewWindow(encodeTitle)
	e.windows.Resize(fyne.NewSize(e.uiW, e.uiH))
	e.windows.SetContent(container.NewGridWithColumns(2,
		container.NewGridWithRows(2,
			e.button.AddButton(sSave, func() {
				//todo 保存文件
				e.file = dialog.NewFileSave(func(u fyne.URIWriteCloser, err error) {
					r, _ := json.Marshal(ruler)
					_, err = u.Write(r)
				}, e.windows)
				e.file.SetFileName("secret.json")
				e.file.Show()
			}),
			e.button.AddButton(sLoad, func() {
				//todo 加载文件
				e.file = dialog.NewFileOpen(func(u fyne.URIReadCloser, err error) {
					r, _ := json.Marshal(ruler)
					_, err = u.Read(r)
				}, e.windows)
				e.file.Show()
			}),
		),
		container.NewGridWithRows(2,
			//widget.NewLabel("请勿设置自相矛盾的规则"),
			e.Entry(),
			e.table,
		),
	))
	e.windows.Show()
	return e
}
