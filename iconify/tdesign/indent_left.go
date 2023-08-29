package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IndentLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 4v2h20V4H2Zm6 7v2h14v-2H8Zm-6 7h20v2H2v-2Zm-.414-6l3.182 3.182l1.414-1.414L4.414 12l1.768-1.768l-1.414-1.414L1.586 12Z"/>`),
		g.Group(children),
	)
}