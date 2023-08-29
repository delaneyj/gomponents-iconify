package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bowl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 15c0 3.9-3.1 7-7 7H9c-3.9 0-7-3.1-7-7v-3h20v3Z"/>`),
		g.Group(children),
	)
}