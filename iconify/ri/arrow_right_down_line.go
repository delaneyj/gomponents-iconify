package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowRightDownLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.59 16.003L5.982 7.397l1.414-1.415l8.607 8.607V7.003h2v11h-11v-2h7.585Z"/>`),
		g.Group(children),
	)
}