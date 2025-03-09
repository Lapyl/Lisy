package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"github.com/tiagomelo/go-clipboard/clipboard"
	"golang.org/x/sys/windows"
)

var zcha chan string
var zind int
var zpat string
var zpaa []string
var mod = windows.NewLazyDLL("user32.dll")

type (
	HANDLE uintptr
	HWND   HANDLE
)
type ImageResult struct {
	Error  error
	Format string
	Image  image.Image
}

func main() {
	go func() {
		fmt.Println("A")
		w1 := new(app.Window)
		w1.Option(app.Title("LISY"))
		w1.Option(app.Size(unit.Dp(600), unit.Dp(600)))
		if err := fromMain(w1); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func getWindow(funcName string) uintptr {
	proc := mod.NewProc(funcName)
	hwnd, _, _ := proc.Call()
	return hwnd
}

func lqsTrim(s string) string {
	s = strings.ReplaceAll(s, "\n", "")
	s = strings.ReplaceAll(s, "\"", "")
	s = strings.TrimSpace(s)
	return s
}

func runListShow(ztop string) {
	zpat = ""
	filepath.Walk(ztop,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			} else {
				if info.IsDir() {
					fmt.Println("Dir: " + path)
				} else {
					if strings.HasSuffix(path, ".jpg") || strings.HasSuffix(path, ".png") {
						fmt.Println("File: " + path)
						zpat = zpat + "," + path
					}
				}
			}
			return nil
		})
	zpaa = strings.Split(zpat, ",")
	fmt.Println(len(zpaa))
	go func() {
		w2 := new(app.Window)
		w2.Option(app.Title("LisyImages"))
		if err := runShow(w2); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	zcha = make(chan string, 1)
	zind = 0
	go func() {
		for {
			if zind < len(zpaa)-2 {
				time.Sleep(time.Second * 2)
				zcha <- "LisyImages"
			}
		}
	}()
	// app.Main()
}

func runShow(w *app.Window) error {
	var ops op.Ops
	th := material.NewTheme()
	var openBtn widget.Clickable
	for {
		switch e := w.Event().(type) {
		case app.FrameEvent:
			select {
			case ztit := <-zcha:
				if zind > len(zpaa)-1 {
					// w.Close()
					return nil
				} else {
					zind = zind + 1
					fmt.Println(zind)
					gtx := app.NewContext(&ops, e)
					w.Option(app.Title(ztit))
					file, err := os.OpenFile(strings.Split(zpat, ",")[zind], 0, 0)
					if err != nil {
						fmt.Println("failed opening image file: %w", err)
					}
					defer file.Close()
					imgData, format, err := image.Decode(file)
					if err != nil {
						fmt.Println("failed decoding image data: %w", err)
					}
					img := ImageResult{Image: imgData, Format: format}
					layout.Flex{Axis: layout.Vertical}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return material.Button(th, &openBtn, zpaa[zind]).Layout(gtx)
						}),
						layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
							return widget.Image{
								Src: paint.NewImageOp(img.Image),
								Fit: widget.Contain,
							}.Layout(gtx)
						}),
					)
					e.Frame(gtx.Ops)
				}
			}
		}
	}
}

func fromMain(w *app.Window) error {
	var xsou widget.Clickable
	var xdes widget.Clickable
	var xlst widget.Clickable
	var xsyn widget.Clickable
	var xsho widget.Clickable
	var ysou widget.Editor
	var ydes widget.Editor
	var ymsg widget.Editor
	var zsou string
	var zdes string
	zmsg := "LIST prepares csv files of lists of folders and files in your specified source folder."
	zmsg = zmsg + "\nSYNC synchronizes your specified backup folder with your specified source folder."
	zmsg = zmsg + "\nSHOW presents a slideshow of images from your specified source folder and then freeze this window."
	var zops op.Ops
	th := material.NewTheme()
	for {
		switch e := w.Event().(type) {
		case app.FrameEvent:
			gtx := app.NewContext(&zops, e)
			if xsou.Clicked(gtx) {
				zold, err := clipboard.New().PasteText()
				if err != nil {
					log.Fatal(err)
				}
				exec.Command("cmd", "/C", "start", "explorer.exe").Start()
				if hwnd := getWindow("GetForegroundWindow"); hwnd != 0 {
					for ysou.Text() == "" {
						znew, _ := clipboard.New().PasteText()
						if znew != zold {
							zsou = lqsTrim(strings.Replace(znew, zold, "", 0))
							if strings.Contains(zsou, ":\\") {
								ysou.SetText(zsou)
							}
						}
					}
				}
			}
			if xdes.Clicked(gtx) {
				zold, err := clipboard.New().PasteText()
				if err != nil {
					log.Fatal(err)
				}
				exec.Command("cmd", "/C", "start", "explorer.exe").Start()
				if hwnd := getWindow("GetForegroundWindow"); hwnd != 0 {
					for ydes.Text() == "" {
						znew, _ := clipboard.New().PasteText()
						if znew != zold {
							zdes = lqsTrim(strings.Replace(znew, zold, "", 0))
							if strings.Contains(zdes, ":\\") {
								ydes.SetText(zdes)
							}
						}
					}
				}
			}
			if xlst.Clicked(gtx) {
				zlin := ""
				zlrd := ""
				zlrf := ""
				ztag := time.Now().Format("06-01-02_15-04")
				ztop := strings.TrimSpace(ysou.Text())
				err := filepath.Walk(ztop,
					func(path string, info os.FileInfo, err error) error {
						if err != nil {
							return err
						}
						zlin = strings.Replace(path, ztop, "", -1) + "," + strconv.Itoa(int(info.Size())) + "," + info.ModTime().Format("06-01-02_15-04")
						fmt.Println(zlin)
						if info.IsDir() {
							zlrd = zlrd + zlin + "\r\n"
						} else {
							zlrf = zlrf + zlin + "\r\n"
						}
						return nil
					})
				if err != nil {
					log.Println(err)
				} else {
					os.WriteFile(ztop+"\\ListDir_"+ztag+".csv", []byte("RelPath,Bytes,ModYmdhm\r\n"+zlrd), 0666)
					os.WriteFile(ztop+"\\ListFil_"+ztag+".csv", []byte("RelPath,Bytes,ModYmdhm\r\n"+zlrf), 0666)
				}
				zmsg = "See: " + ztop + "\\ListDir(Fil)_" + ztag + ".csv"
				fmt.Println(zmsg)
				zmsg = zmsg + "\n" + zlrf
			}
			if xsho.Clicked(gtx) {
				runListShow(strings.TrimSpace(ysou.Text()))
			}
			if xsyn.Clicked(gtx) {
				zsou = strings.TrimSpace(ysou.Text())
				zdes = strings.TrimSpace(ydes.Text())
				if len(ydes.Text()) > 0 {
					zcmd := exec.Command("ROBOCOPY", zsou, zdes, "/MIR")
					zout, err := zcmd.CombinedOutput()
					if err != nil {
						zmsg = "Error: " + string(err.Error())
					} else {
						zmsg = string(zout)
					}
					fmt.Println(zmsg)
				}
				ymsg.SetText(zmsg)
			}
			layout.Flex{
				Axis:    layout.Vertical,
				Spacing: layout.SpaceStart,
			}.Layout(
				gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return material.Button(th, &xsou, "^").Layout(gtx)
						}),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return material.Label(th, unit.Sp(10), "Source folder\nfor List Sync View\n..............................").Layout(gtx)
						}),
						layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
							ed := material.Editor(th, &ysou, "")
							ysou.SingleLine = false
							ysou.Alignment = text.Start
							margins := layout.Inset{
								Top:    unit.Dp(1),
								Bottom: unit.Dp(1),
								Left:   unit.Dp(10),
								Right:  unit.Dp(10),
							}
							border := widget.Border{
								Color:        color.NRGBA{R: 204, G: 204, B: 204, A: 255},
								CornerRadius: unit.Dp(5),
								Width:        unit.Dp(1),
							}
							return margins.Layout(gtx,
								func(gtx layout.Context) layout.Dimensions {
									return border.Layout(gtx, ed.Layout)
								},
							)
						}),
					)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return material.Button(th, &xdes, "^").Layout(gtx)
						}),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return material.Label(th, unit.Sp(10), "Backup folder\nfor Sync\n..............................").Layout(gtx)
						}),
						layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
							ed := material.Editor(th, &ydes, "")
							ydes.SingleLine = false
							ydes.Alignment = text.Start
							margins := layout.Inset{
								Top:    unit.Dp(1),
								Bottom: unit.Dp(1),
								Left:   unit.Dp(10),
								Right:  unit.Dp(10),
							}
							border := widget.Border{
								Color:        color.NRGBA{R: 204, G: 204, B: 204, A: 255},
								CornerRadius: unit.Dp(5),
								Width:        unit.Dp(1),
							}
							return margins.Layout(gtx,
								func(gtx layout.Context) layout.Dimensions {
									return border.Layout(gtx, ed.Layout)
								},
							)
						}),
					)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return material.Label(th, unit.Sp(10), "To input a folder: click ^, select folder in File Explorer, right click, click Copy as path, close File Explorer.").Layout(gtx)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{}.Layout(gtx,
						layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
							return material.Button(th, &xlst, "LIST").Layout(gtx)
						}),
						layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
							return material.Button(th, &xsyn, "SYNC").Layout(gtx)
						}),
						layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
							return material.Button(th, &xsho, "SHOW").Layout(gtx)
						}),
					)
				}),
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					ed := material.Editor(th, &ymsg, zmsg)
					ymsg.SingleLine = false
					ymsg.Alignment = text.Start
					margins := layout.Inset{
						Top:    unit.Dp(1),
						Bottom: unit.Dp(1),
						Left:   unit.Dp(10),
						Right:  unit.Dp(10),
					}
					border := widget.Border{
						Color:        color.NRGBA{R: 204, G: 204, B: 204, A: 255},
						CornerRadius: unit.Dp(5),
						Width:        unit.Dp(1),
					}
					return margins.Layout(gtx,
						func(gtx layout.Context) layout.Dimensions {
							return border.Layout(gtx, ed.Layout)
						},
					)
				}),
			)
			e.Frame(gtx.Ops)
		case app.DestroyEvent:
			return e.Err
		}
	}
}
