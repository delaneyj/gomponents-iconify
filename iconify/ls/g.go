package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func G(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M15 697h75c36 70 111 118 195 118c103 0 189-74 205-172c1-4 1-30 1-68c-50 57-124 93-207 93C132 668 0 548 0 398s132-270 284-270c82 0 156 36 207 92c-1-54-1-92-1-92h72v484c0 14-1 29-3 42c-22 129-136 227-273 227c-122 0-231-76-271-184zm268-95c114 0 207-91 207-204s-93-204-207-204c-115 0-215 91-215 204s100 204 215 204z"/>`),
		g.Group(children),
	)
}