package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bookmark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M17.396 1.84L6.076 25.987l7.34-4.566l1.187 8.564l11.32-24.146l-8.527-3.997zm1.735 7.394a1.125 1.125 0 0 1 .956-2.037a1.124 1.124 0 1 1-.955 2.036z"/>`),
		g.Group(children),
	)
}