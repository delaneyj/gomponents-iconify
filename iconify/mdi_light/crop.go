package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 6h7a3 3 0 0 1 3 3v7h-1V9a2 2 0 0 0-2-2H8V6m0 13a3 3 0 0 1-3-3V7H1V6h4V2h1v14a2 2 0 0 0 2 2h13v1h-3v4h-1v-4H8Z"/>`),
		g.Group(children),
	)
}