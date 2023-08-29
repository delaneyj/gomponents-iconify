package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wheel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 12a9 9 0 1 0 18 0a9 9 0 1 0-18 0"/><path d="M9 12a3 3 0 1 0 6 0a3 3 0 1 0-6 0m-6 0h6m6 0h6m-7.4-2.6L17 4.6m-6.6 10L7 19.4M7 4.6l3.4 4.8m3.2 5.2l3.4 4.8"/></g>`),
		g.Group(children),
	)
}