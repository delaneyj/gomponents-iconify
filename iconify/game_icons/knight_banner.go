package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KnightBanner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m256 23.57l-16.1 48.86h32.2zM73 90.43v15.97h366V90.43zm48 33.97V479l135-105l135 105V124.4zm87 37h96l-32 80l80-32v96l-80-32l32 80h-96l32-80l-80 32v-96l80 32zm48 235.4l-23 17.9v73.7h46v-73.7z"/>`),
		g.Group(children),
	)
}