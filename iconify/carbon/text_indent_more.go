package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextIndentMore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M20 6h10v2H20zm0 6h7v2h-7zm0 6h10v2H20zm0 6h7v2h-7zM14 4h2v24h-2zM3.85 22.5l-1.41-1.38L8.73 15L2.44 8.88L3.85 7.5l7.71 7.5l-7.71 7.5z"/>`),
		g.Group(children),
	)
}