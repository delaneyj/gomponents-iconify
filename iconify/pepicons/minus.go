package pepicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Minus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M5 10h10"/><path fill="currentColor" d="M5 11a1 1 0 1 1 0-2h10a1 1 0 1 1 0 2H5Z"/></g>`),
		g.Group(children),
	)
}