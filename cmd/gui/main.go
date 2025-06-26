// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2015 Martin Lindhe
// SPDX-FileCopyrightText: 2016 The Ebitengine Authors

// The original project is gol (https://github.com/martinlindhe/gol) by Martin Lindhe.

package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hculpan/scorch-point-client/internal/game"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

func main() {
	g := game.NewGame(screenWidth, screenHeight)

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Game of Life (Ebitengine Demo)")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
