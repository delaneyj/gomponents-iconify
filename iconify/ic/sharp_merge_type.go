package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpMergeType(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 20.41L18.41 19L15 15.59L13.59 17L17 20.41zM7.5 8H11v5.59L5.59 19L7 20.41l6-6V8h3.5L12 3.5L7.5 8z"/>`),
		g.Group(children),
	)
}