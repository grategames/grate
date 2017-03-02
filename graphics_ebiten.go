package grate

import "github.com/hajimehoshi/ebiten"
import "github.com/hajimehoshi/ebiten/ebitenutil"

import (
	"path"
	"io/ioutil"
)

var screen *ebiten.Image
var options ebiten.DrawImageOptions
var matrix ebiten.DrawImageOptions
var err error

type EbitenImage struct {
	*ebiten.Image
	options ebiten.DrawImageOptions
	path string
}

func (img *EbitenImage) Load() {
	println("loading!")
	img.Image, _, err = ebitenutil.NewImageFromFile(img.path, ebiten.FilterNearest)
	if err != nil {
		println(err.Error())
	}
}

func (img *EbitenImage) Draw() {
	screen.DrawImage(img.Image, &(img.options))
	img.options = matrix
}

func (img *EbitenImage) Scale(x, y float64) {
	img.options.GeoM.Scale(x, y)
}

func (img *EbitenImage) Translate(x, y float64) {
	img.options.GeoM.Translate(x, y)
}

func (img *EbitenImage) Rotate(angle float64) {
	img.options.GeoM.Rotate(angle)
}

func (img *EbitenImage) Width() float64 {
	var w, _ = img.Image.Size()
	return float64(w)
}

func (img *EbitenImage) Height() float64 {
	var _, h = img.Image.Size()
	return float64(h)
}

type EbitenGraphics struct {}

var cache = make(map[string]Image)
func (EbitenGraphics) Load(dir string) {
	files, _ := ioutil.ReadDir(dir)
    for _, f := range files {
    	switch path.Ext(f.Name()) {
    		case ".jpg", ".JPG", ".png",".jpeg":
    			img := EbitenGraphics{}.Image(dir+"/"+f.Name())
    			img.Load()
    			cache[f.Name()] = img
    		default:
    	}
    }
}

func (EbitenGraphics) Image(path string) Image {
	if img, ok := cache[path]; ok {
		return img
	}
	return &EbitenImage{path:path}
}

func (EbitenGraphics) Width() float64 {
	w, _ := screen.Size()
	return float64(w)
}

func (EbitenGraphics) Height() float64 {
	_, h := screen.Size()
	return float64(h)
}

func (EbitenGraphics) Translate(x, y float64) {
	matrix.GeoM.Translate(x, y)
}
