package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PageNextOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 3H5a2 2 0 0 0-2 2v4h2V5h17v14H5v-4H3v4a2 2 0 0 0 2 2h17a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2M7 15v-2H0v-2h7V9l4 3l-4 3m13-2h-7v-2h7v2m0-4h-7V7h7v2m-3 8h-4v-2h4v2Z"/>`),
		g.Group(children),
	)
}