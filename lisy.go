package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/reujab/wallpaper"
)

var zpri chan bool
var zmsg string
var zlrd string
var zlrf string

func main() {
	go func() {
		w := new(app.Window)
		w.Option(app.Title("LISY"))
		w.Option(app.Size(unit.Dp(400), unit.Dp(600)))
		if err := render(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	zpri = make(chan bool)
	go func() {
		for {
			time.Sleep(time.Second / 25)
			zpri <- true
		}
	}()
	app.Main()
}

func flis(ztop string) {
	ztag := time.Now().Format("06-01-02_15-04")
	zlrd = ""
	zlrf = ""
	fmt.Println("==========================================================")
	fmt.Println(ztop)
	err := filepath.Walk(ztop,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			zlin := strings.Replace(path, ztop, "", -1) + "," + strconv.Itoa(int(info.Size())) + "," + info.ModTime().Format("06-01-02_15-04")
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
}

func render(w *app.Window) error {
	var zbt1 widget.Clickable
	var zbt2 widget.Clickable
	var zbt3 widget.Clickable
	var zbt4 widget.Clickable
	var zibx widget.Editor
	var ztrk bool
	var znum float32
	var zpro float32
	var zops op.Ops
	th := material.NewTheme()
	go func() {
		for range zpri {
			if ztrk && zpro < 1 {
				zpro += 1.0 / 25.0 / znum
				if zpro >= 1 {
					zpro = 1
				}
				w.Invalidate()
			}
		}
	}()
	for {
		switch e := w.Event().(type) {
		case app.FrameEvent:
			gtx := app.NewContext(&zops, e)
			if zbt1.Clicked(gtx) {
				fmt.Println("List")
				zist := zibx.Text()
				zist = strings.TrimSpace(zist)
				flis(zist)
				zibx.SetText(zmsg)
			}
			if zbt2.Clicked(gtx) {
				fmt.Println("Insp")
				ztop := zibx.Text()
				zpat := ""
				w2 := new(app.Window)
				w2.Option(app.Title("Insp"))
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
				for _, path := range strings.Split(zpat, ",") {
					background, err := wallpaper.Get()
					if err != nil {
						fmt.Println(err)
					}
					fmt.Println("Current wallpaper:", background)
					err = wallpaper.SetFromFile(path)
					if err != nil {
						fmt.Println(err)
					}
					err = wallpaper.SetMode(wallpaper.Crop)
					if err != nil {
						fmt.Println(err)
						log.Fatal(err)
					}
				}
			}
			if zbt3.Clicked(gtx) {
				fmt.Println("Sync")
				zist := zibx.Text()
				zist = strings.TrimSpace(zist)
				zina := strings.Split(zist, "\n")
				flis(zina[0])
				zibx.SetText(zmsg)
				flis(zina[1])
				zibx.SetText(zmsg)
				zcmd := exec.Command("ROBOCOPY", zina[0], zina[1], "/MIR")
				out, err := zcmd.CombinedOutput()
				if err != nil {
					fmt.Println("Error: ", err.Error())
					zibx.SetText("Error: " + err.Error())
				} else {
					fmt.Println(out)
					zibx.SetText(string(out))
				}
			}
			if zbt4.Clicked(gtx) {
				fmt.Println("Ying")
				ztrk = !ztrk
				if zpro >= 1 {
					zpro = 0
				}
				zist := zibx.Text()
				zist = strings.TrimSpace(zist)
				zinu, _ := strconv.ParseFloat(zist, 32)
				znum = float32(zinu)
				znum = znum / (1 - zpro)
			}
			layout.Flex{
				Axis:    layout.Vertical,
				Spacing: layout.SpaceStart,
			}.Layout(
				gtx,
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{}.Layout(gtx,
						layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
							return layout.N.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								return layout.UniformInset(8).Layout(gtx,
									material.Button(th, &zbt1, "List").Layout,
								)
							})
						}),
						layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
							return layout.N.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								return layout.UniformInset(8).Layout(gtx,
									material.Button(th, &zbt2, "Insp").Layout,
								)
							})
						}),
						layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
							return layout.N.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								return layout.UniformInset(8).Layout(gtx,
									material.Button(th, &zbt3, "Sync").Layout,
								)
							})
						}),
						layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
							return layout.N.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								return layout.UniformInset(8).Layout(gtx,
									material.Button(th, &zbt4, "Ying").Layout,
								)
							})
						}),
					)
				}),
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					ed := material.Editor(th, &zibx, "List: Top folder.\nInsp: Images folder.\nSync: Source & target.\nYing: Limit second")
					zibx.SingleLine = false
					zibx.Alignment = text.Middle
					if ztrk && zpro < 1 {
						zrem := (1 - zpro) * znum
						zstr := fmt.Sprintf("%.1f", math.Round(float64(zrem)*10)/10)
						zibx.SetText(zstr)
					}
					margins := layout.Inset{
						Top:    unit.Dp(0),
						Right:  unit.Dp(10),
						Bottom: unit.Dp(0),
						Left:   unit.Dp(10),
					}
					border := widget.Border{
						Color:        color.NRGBA{R: 204, G: 204, B: 204, A: 255},
						CornerRadius: unit.Dp(3),
						Width:        unit.Dp(2),
					}
					return margins.Layout(gtx,
						func(gtx layout.Context) layout.Dimensions {
							return border.Layout(gtx, ed.Layout)
						},
					)
				},
				),
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					bar := material.ProgressBar(th, zpro)
					return bar.Layout(gtx)
				},
				),
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						var zimg clip.Path
						op.Offset(image.Pt(gtx.Dp(200), gtx.Dp(190))).Add(gtx.Ops)
						zimg.Begin(gtx.Ops)
						for deg := 0.0; deg <= 360; deg++ {
							rad := deg / 360 * 2 * math.Pi
							cosT := math.Cos(rad)
							sinT := math.Sin(rad)
							a := uint8(75.0 * (2 - zpro))
							b := uint8(100.0 * (2 - zpro))
							d := 20.0
							x := float64(a) * cosT
							y := -(math.Sqrt(float64(b)*float64(b)-d*d*cosT*cosT) + d*sinT) * sinT
							p := f32.Pt(float32(x), float32(y))
							zimg.LineTo(p)
						}
						zimg.Close()
						zare := clip.Outline{Path: zimg.End()}.Op()
						color := color.NRGBA{
							R: 255,
							G: uint8(239 * (1 - zpro)),
							B: uint8(174 * (1 - zpro)),
							A: 255,
						}
						paint.FillShape(gtx.Ops, color, zare)
						d := image.Point{Y: 375}
						return layout.Dimensions{Size: d}
					},
				),
			)
			e.Frame(gtx.Ops)
		case app.DestroyEvent:
			return e.Err
		}
	}
}
