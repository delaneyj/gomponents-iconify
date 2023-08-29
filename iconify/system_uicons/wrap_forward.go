package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WrapForward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m14.5 12.5l2.998-3.002l-3-3"/><path d="M12.5 15.5h-5c-2 0-3-1-3-3s1-3 3-3h10"/></g>`),
		g.Group(children),
	)
}