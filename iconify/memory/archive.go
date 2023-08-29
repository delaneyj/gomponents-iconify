package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Archive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M2 2h18v6h-1v12H3V8H2V2m15 16V8H5v10h12M4 4v2h14V4H4m3 6h8v2H7v-2Z"/>`),
		g.Group(children),
	)
}