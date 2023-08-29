package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lightbulb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M4.41 0a.5.5 0 0 0-.13.06l-3 1.5a.5.5 0 1 0 .44.88l3-1.5A.5.5 0 0 0 4.41 0zm1 1.5a.5.5 0 0 0-.13.06l-4 2a.5.5 0 1 0 .44.88l4-2a.5.5 0 0 0-.31-.94zm0 2a.5.5 0 0 0-.13.06l-3 1.5A.5.5 0 0 0 2.5 6h2a.506.506 0 0 0 .16-1l1.06-.56a.5.5 0 0 0-.31-.94zM2.85 7a.506.506 0 0 0 .16 1h1a.5.5 0 1 0 0-1h-1a.5.5 0 0 0-.09 0a.5.5 0 0 0-.06 0z"/>`),
		g.Group(children),
	)
}