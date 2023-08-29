package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EyeglassTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8 4H6L3 14v2.5M16 4h2l3 10v2.5M10 16h4"/><path d="M14 16.5a3.5 3.5 0 1 0 7 0a3.5 3.5 0 1 0-7 0m-11 0a3.5 3.5 0 1 0 7 0a3.5 3.5 0 1 0-7 0"/></g>`),
		g.Group(children),
	)
}