package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Forward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16 17l-1.425-1.4l4.6-4.6l-4.6-4.6L16 5l6 6l-6 6ZM2 19v-4q0-2.075 1.463-3.538T7 10h6.175l-3.6-3.6L11 5l6 6l-6 6l-1.425-1.4l3.6-3.6H7q-1.25 0-2.125.875T4 15v4H2Z"/>`),
		g.Group(children),
	)
}