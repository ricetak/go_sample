package main

//WALK関連のライブラリ
import (
	"github.com/lxn/win"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

//その他標準ライブラリ
import (
    "fmt"
    "time"
     "os"
    "log"
)

const (
	SIZE_W = 400
	SIZE_H = 140
)

func main() {

	//----------------------------------------------------------
	// 現在時刻の取得
	//----------------------------------------------------------
	t := time.Now()
    const layout = "2006/01/02 15:04:05"
	current_time := " (" + t.Format(layout) + ")"
	fmt.Println(current_time)

	//----------------------------------------------------------
	// ウィンドウ生成
	//----------------------------------------------------------
	mw := new(MyMainWindow)

	MainWindow{
		Visible: false,
		AssignTo: &mw.MainWindow,
		//MinSize: Size{200, 100},
		Title: "Sticky note" + current_time,
		Font: Font{PointSize:11},
		//Background: SolidColorBrush{Color: walk.RGB(33, 230, 33)},
		Layout: VBox {
			 MarginsZero:true,
			 SpacingZero:true,
			 Margins: Margins{0, 0, 0, 0},
		},

		Children: []Widget {
		 	TextEdit {
			 	AssignTo: &mw.textArea,
			 	VScroll: true,
			 	//Background: SolidColorBrush{Color: walk.RGB(33, 230, 33)},
		 	},

		 	Composite{
				Layout: HBox{
					MarginsZero:true,
					SpacingZero:true,
					Margins: Margins{0, 2, 0, 2},
				},

				Children: []Widget{
					HSpacer{},
					PushButton{
						Text: "閉じる",
						OnClicked: func() {
							// ログファイルに保存
							file, err := os.OpenFile("./memo_log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
							if err != nil {
								log.Fatal(err)
							}
							defer file.Close()

							fmt.Fprintln(file, 	"---------------------------------------------------------")
							fmt.Fprintln(file, 	t.Format(layout))
							fmt.Fprintln(file, 	"---------------------------------------------------------")
							fmt.Fprintln(file, 	mw.textArea.Text())
							fmt.Fprintln(file, 	"")

							// ウィンドウをクローズ
							mw.Close()
						},
					},
					HSpacer{},
				},
			},
		 },
	}.Create()

	//----------------------------------------------------------
	// ウィンドウスタイルの設定
	//----------------------------------------------------------
	window_style := win.WS_CAPTION | win.WS_SIZEBOX | win.WS_MAXIMIZEBOX | win.WS_MINIMIZEBOX /* | win.WS_SYSMENU*/
	win.SetWindowLong(mw.Handle(), win.GWL_STYLE, int32(window_style))

	//----------------------------------------------------------
    // ウィンドウ位置の設定
    //----------------------------------------------------------
	xScreen := win.GetSystemMetrics(win.SM_CXSCREEN);
	yScreen := win.GetSystemMetrics(win.SM_CYSCREEN);

	win.SetWindowPos(
		mw.Handle(),
		win.HWND_TOPMOST,	// 最前面ウィンドウ
		(xScreen - SIZE_W)/2,
		(yScreen - SIZE_H)/2,
		SIZE_W,
		SIZE_H,
		/*win.SWP_SHOWWINDOW,*/ win.SWP_FRAMECHANGED, /*win.SWP_DRAWFRAME,*/
    )

	//----------------------------------------------------------
	// ウィンドウ表示
	//----------------------------------------------------------
	win.ShowWindow(mw.Handle(), win.SW_SHOW)

	mw.Run()
}

// ウィンドウの定義
type MyMainWindow struct {
	*walk.MainWindow
	textArea *walk.TextEdit
	LinkLabel *walk.LinkLabel
}


