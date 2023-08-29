package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableOfContents(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4 6h18v2H4zm0 6h18v2H4zm0 6h18v2H4zm0 6h18v2H4zM26 6h2v2h-2zm0 6h2v2h-2zm0 6h2v2h-2zm0 6h2v2h-2z"/>`),
		g.Group(children),
	)
}