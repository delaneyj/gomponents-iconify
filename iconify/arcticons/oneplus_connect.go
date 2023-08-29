package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OneplusConnect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m6.247 34.25l35.507-20.5L24 3.5L6.247 13.75v20.5L24 44.5l17.754-10.25l-35.507-20.5"/>`),
		g.Group(children),
	)
}