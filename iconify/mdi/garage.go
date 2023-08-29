package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Garage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 20h-2v-9H7v9H5V9l7-4l7 4v11M8 12h8v2H8v-2m0 3h8v2H8v-2m8 3v2H8v-2h8Z"/>`),
		g.Group(children),
	)
}