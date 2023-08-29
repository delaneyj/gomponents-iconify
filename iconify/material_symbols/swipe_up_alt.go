package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwipeUpAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 20q-2.075 0-3.538-1.463T7 15q0-1.825 1.137-3.188T11 10.1V5.825L9.4 7.4L8 6l4-4l4 4l-1.4 1.425l-1.6-1.6V10.1q1.725.35 2.863 1.713T17 15q0 2.075-1.463 3.538T12 20Z"/>`),
		g.Group(children),
	)
}