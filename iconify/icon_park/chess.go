package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chess(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="2" stroke-width="4"><path d="M34 4H14V14H34V4Z"/><path d="M27 14L33 37H15L21 14"/><path d="M40 44H8V40L14 37H34L40 40V44Z"/><path d="M12 23.02H36"/><path d="M20.5 4V8"/><path d="M27.5 4V8"/></g>`),
		g.Group(children),
	)
}