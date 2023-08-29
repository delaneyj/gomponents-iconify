package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Axe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m13 9l7.383 7.418c.823.82.823 2.148 0 2.967a2.11 2.11 0 0 1-2.976 0L10 12"/><path d="m6.66 15.66l-3.32-3.32a1.25 1.25 0 0 1 .42-2.044L7 9l6-6l3 3l-6 6l-1.296 3.24a1.25 1.25 0 0 1-2.044.42z"/></g>`),
		g.Group(children),
	)
}