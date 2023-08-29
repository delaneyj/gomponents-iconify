package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeMinus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 17v2h-8v-2h8M12 3l10 9h-4a6.005 6.005 0 0 0-5.66 8H5v-8H2l10-9Z"/>`),
		g.Group(children),
	)
}