package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pick(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m13 8l-9.383 9.418a2.091 2.091 0 0 0 0 2.967a2.11 2.11 0 0 0 2.976 0L16 11"/><path d="M9 3h4.586a1 1 0 0 1 .707.293l6.414 6.414a1 1 0 0 1 .293.707V15a2 2 0 1 1-4 0v-3l-5-5H9a2 2 0 1 1 0-4z"/></g>`),
		g.Group(children),
	)
}