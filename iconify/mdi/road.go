package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Road(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 16h2v4h-2m0-10h2v4h-2m0-10h2v4h-2M4 22h16V2H4v20Z"/>`),
		g.Group(children),
	)
}