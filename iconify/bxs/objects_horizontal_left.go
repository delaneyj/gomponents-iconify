package bxs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ObjectsHorizontalLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h2v20H2z"/><rect width="16" height="6" x="6" y="13" fill="currentColor" rx="1"/><rect width="12" height="6" x="6" y="5" fill="currentColor" rx="1"/>`),
		g.Group(children),
	)
}