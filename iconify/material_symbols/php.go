package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Php(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 15V9h3.5q.6 0 1.05.45T8 10.5v1q0 .6-.45 1.05T6.5 13h-2v2H3Zm6.5 0V9H11v2h2V9h1.5v6H13v-2.5h-2V15H9.5Zm7 0V9H20q.6 0 1.05.45t.45 1.05v1q0 .6-.45 1.05T20 13h-2v2h-1.5Zm-12-3.5h2v-1h-2v1Zm13.5 0h2v-1h-2v1Z"/>`),
		g.Group(children),
	)
}