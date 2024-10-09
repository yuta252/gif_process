package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/gif"
	"image/png"
	"os"
)

func main() {
	file, err := os.Open("input/marugame.gif")
	if err != nil {
		fmt.Println("Error opening GIF:", err)
		return
	}
	defer file.Close()

	gifImage, err := gif.DecodeAll(file)
	if err != nil {
		fmt.Println("Error decoding GIF:", err)
		return
	}
	// ベースとなるキャンバスを用意（最初のフレームを基準にサイズを決定）
	var canvas *image.Paletted
	bounds := gifImage.Image[0].Bounds()

	// 各フレームを重ねていく処理
	for i, frame := range gifImage.Image {
		if i == 0 {
			// 最初のフレームをベースとしてキャンバスを作成
			canvas = image.NewPaletted(bounds, frame.Palette)
		}

		// 現在のフレームをキャンバスに重ねる
		draw.Draw(canvas, bounds, frame, image.Point{}, draw.Over)

		// 各フレームをPNGとして保存
		outputFilename := fmt.Sprintf("output/frame_%d.png", i)
		outputFile, err := os.Create(outputFilename)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer outputFile.Close()

		// PNG形式でフレームを書き出す
		err = png.Encode(outputFile, canvas)
		if err != nil {
			fmt.Println("Error encoding PNG:", err)
			return
		}

		fmt.Printf("Saved frame %d as %s\n", i, outputFilename)
	}
}
