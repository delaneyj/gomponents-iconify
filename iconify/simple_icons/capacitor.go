package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Capacitor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m24 3.7l-5.766 5.766l5.725 5.736l-3.713 3.712L5.073 3.742L8.786.03l5.736 5.726L20.284 0L24 3.7zM.029 8.785l3.713-3.713l15.173 15.173l-3.713 3.714l-5.732-5.726L3.7 24L0 20.285l5.754-5.764L.029 8.785z"/>`),
		g.Group(children),
	)
}