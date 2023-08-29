package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeMinus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 20H5a1 1 0 0 1-1-1v-7.86a1 1 0 0 1 .281-.695l5.5-5.7a1 1 0 0 1 1.439 0l2.8 2.9l-1.43 1.402L10.5 6.88L6 11.54v6.455h11v1A1 1 0 0 1 16 20Zm4-6h-8v-2h8v2Z"/>`),
		g.Group(children),
	)
}