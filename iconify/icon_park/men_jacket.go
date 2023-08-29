package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MenJacket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M6 10L18 4H30L42 10L40 35H34V44H24H14V35H8L6 10Z"/><path d="M14 35L14 20"/><path d="M34 35V20"/><path d="M24 10C27.3137 10 30 7.31371 30 4H18C18 7.31371 20.6863 10 24 10Z"/></g>`),
		g.Group(children),
	)
}