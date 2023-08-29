package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExpressionlessFaceFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12Zm4.5-2.5a1 1 0 0 1 1-1H10a1 1 0 1 1 0 2H7.5a1 1 0 0 1-1-1Zm7.5-1a1 1 0 1 0 0 2h2.5a1 1 0 1 0 0-2H14ZM9 14a1 1 0 1 0 0 2h6a1 1 0 1 0 0-2H9Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}