package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Strategy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m14 22l-6-3.5v-7l3.5-2.05v6.05h5V9.45L20 11.5v7L14 22ZM5.5 11L2 9V5l3.5-2L9 5v4l-3.5 2Zm7.5 3V2h9l-2 3l2 3h-7v6h-2Z"/>`),
		g.Group(children),
	)
}