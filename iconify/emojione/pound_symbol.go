package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PoundSymbol(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#4fd1d9"/><path fill="#fff" d="m49.4 28l1.6-6h-6.1l1.6-6h-6.1l-1.6 6h-8.1l1.6-6h-6.1l-1.6 6h-6.1l-1.6 6H23l-2.2 8h-6.1L13 42h6.1l-1.6 6h6.1l1.6-6h8.1l-1.6 6h6.1l1.6-6h6.1l1.6-6H41l2.2-8h6.2M35 36h-8.1l2.2-8h8.1L35 36"/>`),
		g.Group(children),
	)
}