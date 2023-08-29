package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Alert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 21L11.5 2.81L22 21H1m19.27-1L11.5 4.81L2.73 20h17.54M11 14v-4h1v4h-1m0 2h1v2h-1v-2Z"/>`),
		g.Group(children),
	)
}