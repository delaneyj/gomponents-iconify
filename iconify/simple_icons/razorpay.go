package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Razorpay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m22.436 0l-11.91 7.773l-1.174 4.276l6.625-4.297L11.65 24h4.391l6.395-24zM14.26 10.098L3.389 17.166L1.564 24h9.008l3.688-13.902Z"/>`),
		g.Group(children),
	)
}