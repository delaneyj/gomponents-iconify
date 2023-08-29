package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatStrike(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 7h4V5H7v2h4v3h2V7Zm-2 12v-5h2v5h-2Zm-6-6h14v-2H5v2Z"/>`),
		g.Group(children),
	)
}