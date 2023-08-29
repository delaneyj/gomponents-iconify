package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 4v12.17l-5.59-5.59L4 12l8 8l8-8l-1.41-1.41L13 16.17V4h-2z"/>`),
		g.Group(children),
	)
}