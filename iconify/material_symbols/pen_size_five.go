package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PenSizeFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.875 18.125Q5 17.25 5 16t.875-2.125l8-8Q14.75 5 16 5t2.125.875Q19 6.75 19 8t-.875 2.125l-8 8Q9.25 19 8 19t-2.125-.875Z"/>`),
		g.Group(children),
	)
}