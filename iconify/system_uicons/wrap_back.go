package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WrapBack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M6.499 6.498L3.5 9.5l3 3"/><path d="M8.5 15.5h5c2 0 3-1 3-3s-1-3-3-3h-10"/></g>`),
		g.Group(children),
	)
}