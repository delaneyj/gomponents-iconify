package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MarkdownMarkSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M14.846 12.923H1.154A1.154 1.154 0 0 1 0 11.77V4.231a1.154 1.154 0 0 1 1.154-1.154h13.692A1.154 1.154 0 0 1 16 4.23v7.538a1.154 1.154 0 0 1-1.154 1.154Zm-11-2.308v-3l1.539 1.923l1.538-1.923v3h1.539v-5.23H6.923L5.385 7.308L3.846 5.385H2.308v5.23h1.538ZM14.154 8h-1.539V5.385h-1.538V8H9.538l2.308 2.692L14.154 8Z"/>`),
		g.Group(children),
	)
}