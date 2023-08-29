package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NegativeDynamics(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><path d="M17 33.9502V42.1102"/><path d="M9 40V42.0556"/><path d="M25 27V42.0714"/><path d="M33 18.9614V42.0878"/><path d="M41 10.9707V42.0833"/><path d="M7 33L34 6"/><path d="M7 22L7 33"/></g>`),
		g.Group(children),
	)
}