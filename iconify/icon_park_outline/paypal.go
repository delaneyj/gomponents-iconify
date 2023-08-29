package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Paypal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M33.031 28C39 28 43 25.5 43 20s-4-8-9.969-8H22l-5 31h9l2-15h5.031Z" clip-rule="evenodd"/><path d="M18 36h-8l5-31h11.031C32 5 36 7.5 36 13s-4 8-9.969 8H21"/></g>`),
		g.Group(children),
	)
}