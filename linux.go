package main

import (
	"os"
	"strings"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"golang.design/x/clipboard"
)

func messageBox(message_type gtk.MessageType, content string) {
	dialog := gtk.MessageDialogNew(nil, gtk.DIALOG_MODAL, message_type, gtk.BUTTONS_OK, content)
	dialog.Run()
	dialog.Destroy()
}
func save() {
	img_bytes := clipboard.Read(clipboard.FmtImage)
	if len(img_bytes) == 0 {
		glib.IdleAdd(func() {
			messageBox(gtk.MESSAGE_ERROR, "复制的不是图像")
		})
		return
	}

	glib.IdleAdd(func() {
		openFileDialog(img_bytes)
	})

}

func openFileDialog(image []byte) {
	dialog, err := gtk.FileChooserDialogNewWith2Buttons(
		"保存文件",
		nil,
		gtk.FILE_CHOOSER_ACTION_SAVE,
		"_保存", gtk.RESPONSE_ACCEPT,
		"_取消", gtk.RESPONSE_CANCEL,
	)
	if err != nil {
		messageBox(gtk.MESSAGE_ERROR, "打开文件选择器失败")
		return
	}
	filepath := ""
	defer dialog.Destroy()
	response := dialog.Run()
	if response == gtk.RESPONSE_ACCEPT {
		filepath = dialog.GetFilename()
	}
	if filepath != "" {
		if !strings.HasSuffix(filepath, ".png") {
			filepath = filepath + ".png"
		}
		err = os.WriteFile(filepath, image, 0644)
		if err != nil {
			messageBox(gtk.MESSAGE_ERROR, "保存图片失败")
		}
	}
}
