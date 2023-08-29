package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DogSide(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19 3l-4 4l3 3l1-1l1 1l2-2l-3-3V3M3 7L2 8l3 3v3l-1 1v6h2v-3l2-3h7v6h2V11l-3-3l-1 1H5L3 7Z"/>`),
		g.Group(children),
	)
}