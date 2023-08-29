package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleChevronsRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m9 9l3 3l-3 3m4-6l3 3l-3 3"/><path d="M3 12a9 9 0 1 0 0-.265V12z"/></g>`),
		g.Group(children),
	)
}