package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Coffin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 3L5 9l2 12h6l2-12l-2-6zm3 4v5M8 9h4"/><path d="M13 21h4l2-12l-2-6h-4"/></g>`),
		g.Group(children),
	)
}