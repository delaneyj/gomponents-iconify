package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlusMinus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 4v5H6v2h5v5h2v-5h5V9h-5V4h-2M6 18v2h12v-2H6Z"/>`),
		g.Group(children),
	)
}