package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignPiFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6 8a1 1 0 0 1 1-1h10a1 1 0 1 1 0 2h-1v5.03c0 .536.434.97.97.97H17a1 1 0 1 1 0 2h-.03A2.97 2.97 0 0 1 14 14.03V9h-4v7a1 1 0 1 1-2 0V9H7a1 1 0 0 1-1-1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}