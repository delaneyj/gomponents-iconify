package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GridCircles(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7.5" cy="7.5" r="2"/><circle cx="13.5" cy="7.5" r="2"/><circle cx="7.5" cy="13.5" r="2"/><circle cx="13.5" cy="13.5" r="2"/></g>`),
		g.Group(children),
	)
}