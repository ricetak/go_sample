package main

// WALK
import (
	"github.com/lxn/win"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

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

	t := time.Now()
    const layout = "2006/01/02 15:04:05"
	current_time := " (" + t.Format(layout) + ")"
	fmt.Println(current_time)

	//----------------------------------------------------------
	// Create Window
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
						Text: "Close",
						OnClicked: func() {
							// logging
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

							mw.Close()
						},
					},
					HSpacer{},
				},
			},
		 },
	}.Create()

	//----------------------------------------------------------
	// Set Window Style
	//----------------------------------------------------------
	window_style := win.WS_CAPTION | win.WS_SIZEBOX | win.WS_MAXIMIZEBOX | win.WS_MINIMIZEBOX /* | win.WS_SYSMENU*/
	win.SetWindowLong(mw.Handle(), win.GWL_STYLE, int32(window_style))

	//----------------------------------------------------------
    // Set Window Position
    //----------------------------------------------------------
	xScreen := win.GetSystemMetrics(win.SM_CXSCREEN);
	yScreen := win.GetSystemMetrics(win.SM_CYSCREEN);

	win.SetWindowPos(
		mw.Handle(),
		win.HWND_TOPMOST,	// TOPMOST
		(xScreen - SIZE_W)/2,
		(yScreen - SIZE_H)/2,
		SIZE_W,
		SIZE_H,
		/*win.SWP_SHOWWINDOW,*/ win.SWP_FRAMECHANGED, /*win.SWP_DRAWFRAME,*/
    )

	//----------------------------------------------------------
	// Show
	//----------------------------------------------------------
	win.ShowWindow(mw.Handle(), win.SW_SHOW)

	mw.Run()
}

// Window
type MyMainWindow struct {
	*walk.MainWindow
	textArea *walk.TextEdit
	LinkLabel *walk.LinkLabel
}


