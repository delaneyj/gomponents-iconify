package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Beauty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M42 44V22c0-9.941-8.059-18-18-18S6 12.059 6 22v22"/><path d="M24 43c5.523 0 10-8.954 10-20H14c0 11.046 4.477 20 10 20Z"/></g>`),
		g.Group(children),
	)
}