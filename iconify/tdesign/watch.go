package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Watch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.117 0h9.766l.503 4.025A3 3 0 0 1 20 7v2h1v6h-1v2a3 3 0 0 1-2.614 2.975L16.883 24H7.117l-.503-4.025A3 3 0 0 1 4 17V7a3 3 0 0 1 2.614-2.975L7.117 0Zm1.516 4h6.734l-.25-2H8.883l-.25 2ZM7 18h10a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1H7a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1Zm1.633 2l.25 2h6.234l.25-2H8.633Z"/>`),
		g.Group(children),
	)
}