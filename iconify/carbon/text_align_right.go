package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextAlignRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6 6h14v2H6zm4 6h10v2H10zm-4 6h14v2H6zm4 6h10v2H10zM24 4h2v24h-2z"/>`),
		g.Group(children),
	)
}