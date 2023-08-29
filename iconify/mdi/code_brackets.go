package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CodeBrackets(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 4v2h3v12h-3v2h5V4M4 4v16h5v-2H6V6h3V4H4Z"/>`),
		g.Group(children),
	)
}