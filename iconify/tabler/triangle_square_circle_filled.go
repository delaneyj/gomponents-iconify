package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleSquareCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="m11.132 2.504l-4 7A1 1 0 0 0 8 11h8a1 1 0 0 0 .868-1.496l-4-7a1 1 0 0 0-1.736 0zM17 13a4 4 0 1 1-3.995 4.2L13 17l.005-.2A4 4 0 0 1 17 13zm-8 0H5a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2v-4a2 2 0 0 0-2-2z"/></g>`),
		g.Group(children),
	)
}