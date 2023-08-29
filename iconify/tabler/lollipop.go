package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lollipop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 10a7 7 0 1 0 14 0a7 7 0 1 0-14 0"/><path d="M21 10a3.5 3.5 0 0 0-7 0m0 0a3.5 3.5 0 0 1-7 0m7 7a3.5 3.5 0 0 0 0-7m0-7a3.5 3.5 0 0 0 0 7M3 21l6-6"/></g>`),
		g.Group(children),
	)
}