package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WheelLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="12" cy="12" r="10"/><circle cx="12" cy="12" r="6"/><circle cx="12" cy="12" r="2"/><path stroke-linecap="round" d="M6 12h4m4 0h4m-9 5.196l2-3.464m2-3.464l2-3.464m0 10.392l-2-3.464m-2-3.464L9 6.804" opacity=".5"/></g>`),
		g.Group(children),
	)
}