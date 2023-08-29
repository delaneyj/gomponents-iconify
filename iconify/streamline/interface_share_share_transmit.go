package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceShareShareTransmit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="2.75" cy="7" r="2.25"/><circle cx="11.25" cy="11.25" r="2.25"/><circle cx="11.25" cy="2.75" r="2.25"/><path d="m4.76 6l4.48-2.25M4.76 8l4.48 2.25"/></g>`),
		g.Group(children),
	)
}