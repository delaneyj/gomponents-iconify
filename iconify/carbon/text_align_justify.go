package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextAlignJustify(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6 6h20v2H6zm0 6h20v2H6zm0 6h20v2H6zm0 6h20v2H6z"/>`),
		g.Group(children),
	)
}