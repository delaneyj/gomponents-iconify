package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Refresh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M42 8v16M6 24v16m36-16c0-9.941-8.059-18-18-18a17.947 17.947 0 0 0-12.952 5.5M6 24c0 9.941 8.059 18 18 18a17.94 17.94 0 0 0 12.5-5.048"/>`),
		g.Group(children),
	)
}