package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PentagonLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 3.03l8.53 6.198l-3.258 10.029H6.728L3.469 9.228L12 3.031Zm10.747 5.478L12 .7L1.253 8.508l4.105 12.634h13.284l4.105-12.634Z"/>`),
		g.Group(children),
	)
}